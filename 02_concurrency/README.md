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

