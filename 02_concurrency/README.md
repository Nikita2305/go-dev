# Concurrency

Мы рассмотрим здесь многое. Начнём с многопоточности (чтобы потом понять, в чём смысл многозадачности).

## Многопоточность

```
gcc -pthread thread_create.c -o threads
./threads
```

Мы попробовали простую программу на c. А следующий код на macos не запустится, потому что зависит от внутренностей Linux Kernel:

```
strace -f -e trace=clone,mmap,mprotect,arch_prctl -s 100 ./threads
```

Эта программа предназначена для перехвата syscalls из user space в kernel space. `-f` - следить за вызовами в дочерних потоках. `-e trace=` - фильтр по системным вызовам. `-s 100` - даём больше лимит для каждой строки вывода. Отметим что не все системные вызовы могут логироваться, увидим это далее:

Видим вывод:

```
arch_prctl(0x3001 /* ARCH_??? */, 0x7ffd8c025900) = -1 EINVAL (Invalid argument)
mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f5586f6f000
mmap(NULL, 26983, PROT_READ, MAP_PRIVATE, 3, 0) = 0x7f5586f68000
mmap(NULL, 2264656, PROT_READ, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f5586d3f000
mprotect(0x7f5586d67000, 2023424, PROT_NONE) = 0
mmap(0x7f5586d67000, 1658880, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x28000) = 0x7f5586d67000
mmap(0x7f5586efc000, 360448, PROT_READ, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x1bd000) = 0x7f5586efc000
mmap(0x7f5586f55000, 24576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x215000) = 0x7f5586f55000
mmap(0x7f5586f5b000, 52816, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7f5586f5b000
mmap(NULL, 12288, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f5586d3c000
arch_prctl(ARCH_SET_FS, 0x7f5586d3c740) = 0
mprotect(0x7f5586f55000, 16384, PROT_READ) = 0
mprotect(0x55a2975b8000, 4096, PROT_READ) = 0
mprotect(0x7f5586fa9000, 8192, PROT_READ) = 0
Main thread started: PID = 48196, TID=48196
```

Произошла подготовка к запуску главного потока. Для него верно tid=pid. 

```
mmap(NULL, 8392704, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f558653b000
mprotect(0x7f558653c000, 8388608, PROT_READ|PROT_WRITE) = 0
strace: Process 48197 attached
Thread started: PID = 48196, TID=48197
Thread finished
[pid 48197] +++ exited with 0 +++
+++ exited with 0 +++
```

Далее мы выделяем 8mb для стека второго потока, выдаём право на чтение на них (кроме первой страницы, в неё мы запрещаем писать, чтобы отслеживать переполнение стека).

Как-то мы не заметили clone - системный вызов для клонирования. Поэтому используем другую функцию создания потока:

```
gcc -pthread thread_clone.c -o threads
strace -f -e trace=clone,mmap,mprotect,arch_prctl -s 100 ./threads
```

И далее видим:

```
Main thread started: PID = 48834, TID = 48834
mmap(NULL, 1052672, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f65e2d41000
clone(child_stack=0x7f65e2e41000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|SIGCHLD) = 48835
strace: Process 48835 attached
Thread started: PID = 48834, TID = 48835
[pid 48835] +++ exited with 0 +++
Main thread finished
+++ exited with 0 +++
```

Получается если кратко - выделить mmap память на виртуальном пространстве процесса, сказать новому потоку что это его стек.
CLONE_VM | CLONE_FS | CLONE_SIGHAND | CLONE_THREAD - дают понимание clone() - что мы хотим создать новый поток, хотим чтобы эти потоки делили виртуальную память, файловые дескрипторы, обработчики сигналов. Разный будет только стек.

В pthread_create происходит ровно это, только дополнительно идёт работа с защитой первой страницы. В fork тоже используется clone! Только там уже не передаются вышеописанные флаги, а остаётся только SIGCHLD. Единственное для копирования VM - используется Copy On Write - чтобы копировать только то, что меняется и разделять память для чтения.

Отдельно отметим, что есть TLS - thread local storage. Там хранятся глобальные переменные, которые мы хотим не держать синхронизированными. Именно это делает arch_prctl - устанавливает для потока указатель на хранилище.

Мы рассмотрели `strace`, однако стоит не забывать про `ltrace` - вызовы библиотечных функций, `gdb` - отладчик с пошаговым исполнением и инвестигированием памяти вплоть до регистров, `perf, pprof, flamegraph, go tool trace`.

### Непрактическое примечание

Отдельно стоит рассмотреть механизм context switch для CPU и вообще CFS (completely fair scheduler) в linux. Идея в том, что есть сущность ядра - планировщик, которая решает какой тред исполнить на каком ядре. Всем тредам должно доставаться плюс минус одинаково времени исполнения. Поток снимается с исполнения при блокировке (на инпут, например) или благодаря планировщику (истекло время / preemption - вытеснение более приоритетным потоком). Далее ядро совершает сохранение контекста - сохраняет регистры, указатель на стек и инструкции - в память ядра и возвращает задачу планировщику.  

## Многозадачность

Теперь перейдём к разбору многозадачности go.

`go build -o goroutines_test main.go`

`GODEBUG=schedtrace=1000 ./goroutines_test`

Мы дали планироващику go задачу - раз в 1000 миллисекунд печатать своё состояние.

Получим вывод:

```
Num CPU: 1
GOMAXPROCS: 1
SCHED 0ms: gomaxprocs=1 idleprocs=0 threads=4 spinningthreads=0 needspinning=1 idlethreads=1 runqueue=1935 [219]
Started 100000 goroutines
SCHED 1005ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 needspinning=0 idlethreads=2 runqueue=0 [0]
SCHED 2007ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 needspinning=0 idlethreads=2 runqueue=0 [0]
SCHED 3010ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 needspinning=0 idlethreads=2 runqueue=0 [0]
SCHED 4012ms: gomaxprocs=1 idleprocs=1 threads=4 spinningthreads=0 needspinning=0 idlethreads=2 runqueue=0 [0]
```

По умолчанию был создан один процессор задач. Это такая сущность под планировщиком со своей локальной очередью задач.

Попробуем описать первую строку:
- `gomaxprocs=1` - константа, описанная выше, сколько процессоров планировщик может создать максимально
- `idleprocs=0` - сколько процессоров простаивает,
 без возможности исполнить какую-нибудь горутину ни из глобальной очереди, ни из локальной.
- `threads=4` - сколько потоков выделено в этом планировщике
- `spinningthreads=0` - сколько потоков активно ищут себе задачу
- `needspinning=1` - столько потоков хочется назначить на работы 
- `idlethreads=1` - потоки, которые не привязаны к процессорам и ждут своих задач
- `runqueue=1935` - размер глобальной очереди планировщика
- `[219]` - размер P-локальной очереди

Чтобы проверить своё понимание, давай запустим код с несколькими процессорами и повышенной частотой опроса: `GOMAXPROCS=2 GODEBUG=schedtrace=100 ./goroutines_test`

```
Num CPU: 1
GOMAXPROCS: 2
SCHED 0ms: gomaxprocs=2 idleprocs=0 threads=5 spinningthreads=1 needspinning=0 idlethreads=2 runqueue=903 [0 169]
SCHED 100ms: gomaxprocs=2 idleprocs=0 threads=5 spinningthreads=1 needspinning=0 idlethreads=1 runqueue=0 [0 0]
SCHED 201ms: gomaxprocs=2 idleprocs=0 threads=5 spinningthreads=1 needspinning=0 idlethreads=1 runqueue=0 [0 0]
SCHED 302ms: gomaxprocs=2 idleprocs=0 threads=5 spinningthreads=1 needspinning=1 idlethreads=1 runqueue=904 [233 0]
Started 100000 goroutines
SCHED 404ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0]
SCHED 505ms: gomaxprocs=2 idleprocs=2 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0]
```

Получается пока мы планировали горутины, всё больше и больше потоков уходили из idle и назначались на процессоры для разгребания очередей горутин. Когда планирование закончилось, все потоки и процессоры ушли в idle, потому что больше исполнять особо нечего.

Стоит отметить, что на 400мс, мы видим пустые очереди, но ведь горутины где-то существуют? Ответ состоит в том, что есть отдельная структура внутри go runtime: timer heap. Когда горутина лежит в ожидании, она помещается в эту кучу и паркуется (снимается с потока). За пробуждение отвечает системная горутина sysmon - она будит горутины и помещает их в runqueue.

Несколько интересных нюансов:
- если одна горутина перепланирует что-то, то она перепланирует это на тот же процессор (если локальная очередь не забита, иначе на глобальную очередь)
- work stealing - когда крутящийся поток на процессоре понимает, что локальная очередь P пустая, то он отправляется в локальные очереди других P, и потом в глобальную очередь.

Это всё называется G - M - P модель планировщика.

Замечание: force preemption - снятие плохой горутины, например tight loop, когда мы не отдаём управление планировщику. Пример в `main_2.go`. Добавили начиная с go1.14

Классная опция `GODEBUG=scheddetail=1,schedtrace=100 ./preemption` - можно прямо смотреть за айтемами GMP модели в процессе исполнения.

**Стоит отметить, что уже на этом этапе мы можем заметить преимущество горутин перед потоками. Размер стека горутин значительно меньше, чем стек потоков. Создание 100'000 потоков бы спровоцировало выделение 800gb оперативной памяти, что практически невозможно. В свою очередь горутины съедают 200mb. Это свойство и называют легковесностью.**

Перед переходом к обсуждению этого свойства, давайте ещё немного обсудим природу планировщика. Мы помним, что для тредов это затратная операция - нужно где-то в памяти ядра хранить много структур, запоминать регистры.

Здесь go тоже победил, ведь нам не нужно обращаться к ядру для смены горутины. P-процессор (аналог ядра CPU) сохраняет указатель на стек, указатель на операцию (аналог rip) в некоторую секцию G (секция - sched). А затем берёт новую горутину, находит указатель на её стек и инструкцию, и продолжает исполнение. Для счедулера вся эта информация доступна без обращений к ядру.

### Обеспечение легковесности goroutines

Так, ну допустим горутине при её создании выдаётся маленький стек (2kb), а как же мы там заведём массив на 10^7 для решета эратосфена? Попробуем ответить на это:

```
go build -gcflags='-m -l' -o main main_3.go
# command-line-arguments
./main_3.go:8:13: make([]byte, 10000) does not escape
./main_3.go:10:13: ... argument does not escape
./main_3.go:10:14: "v:" escapes to heap
./main_3.go:10:20: v escapes to heap
```

Хм, кажется это не совсем то, что ожидалось. Мы наблюдали escape анализ! Это отдельная тема для разговора, получается когда мы используем что-то как аргумент функции и не можем гарантировать, что переменная не будет использована, после её снятия со стека, то происходит эскейп - выделяется память на куче и информация уезжает туда.

Более того make выделяет память на куче.

Однако нам интересно, как происходит выделение на стеке 10000 byte = 10kb, что больше чем стек горутины.

```
root@2377933-ms45795:~/go-dev/02_concurrency# go build -o main main_4.go
root@2377933-ms45795:~/go-dev/02_concurrency# ./main
sum: 200
depth:  200 | SP ~= 0xc00007bb38
sum: 199
depth:  199 | SP ~= 0xc00007b6e0
sum: 198
depth:  198 | SP ~= 0xc00008d288 -> growth happend
sum: 197
depth:  197 | SP ~= 0xc00008ce30
sum: 196
depth:  196 | SP ~= 0xc00008c9d8
sum: 195
depth:  195 | SP ~= 0xc000096580 -> growth happend
sum: 194
depth:  194 | SP ~= 0xc000096128
sum: 193
depth:  193 | SP ~= 0xc000095cd0
sum: 192
depth:  192 | SP ~= 0xc000095878
sum: 191
depth:  191 | SP ~= 0xc000095420
sum: 190
depth:  190 | SP ~= 0xc000094fc8
sum: 189
depth:  189 | SP ~= 0xc000094b70
sum: 188
depth:  188 | SP ~= 0xc000094718 
sum: 187
depth:  187 | SP ~= 0xc00009c2c0 -> growth happend
```

Идея этого кода состоит в том, что мы создаём рекурсию, которая неанализируется статически и только в рантайме можно понять сколько стека нам понадобится. Поэтому компилятор выделяет нам 2kb, затем видно что через 2 итерации у нас стек растёт (SP поднимается вверх), потом ещё через 4 итерации растёт, потом ещё через 8. Значит действительно фактор роста около 2. 

Примечание про рост стека. Как горутина, стоя на стеке, может его увеличить? В дело вступает g0 - служебная горутина привязанная к треду. Её стек побольше (8+kb) и на ней запускаются системные вызовы, например, так и расширение и переезд стека.

Без g0 невозможно:

- Stack growth - Нельзя копировать стек, стоя на нём → переключение на g0
- Garbage Collection - Нужен служебный стек для GC логики
- sysmon / background	- Планировщик, preempt, timers — всё работает на g0
- Start goroutine	- newproc запускается через g0
- Panic + recover	- Обработка stack unwinding делается с помощью g0

Забавно, что код планировщика исполняется прямо на треде (которым планировщик управляет). Stack unwinding - это когда стек исполнения goroutine разматывается в обратную сторону, вызывает defer и ищет recover. runtime.newproc - это то что происходит когда мы запускаем go f() - буквально создать G и положить её во внутреннюю очередь P.  

### Непрактическое замечание

В go реализованы stackful корутины - ведь у каждой горутины есть свой стек. Для stackless - это работает по другому (например python). Пример:

```
async def foo():
    x = 42
    await asyncio.sleep(1)
    print(x)
```

Под капотом компилятор превращает это в машину состояний, похожую на (то есть весь стек горутины хранится прямо в объекте горутины):

```
class FooCoroutine:
    def __init__(self):
        self._state = 0
        self.x = None

    def __await__(self):
        if self._state == 0:
            self.x = 42
            self._state = 1
            yield from asyncio.sleep(1).__await__()
        if self._state == 1:
            print(self.x)
```

При таком подходе у нас для создания каждой корутины происходит по несколько аллокаций: на словарь переменных, на сам объект корутины и на другие штуки. 