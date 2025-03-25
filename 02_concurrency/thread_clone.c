#define _GNU_SOURCE
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/syscall.h>
#include <sys/types.h>
#include <signal.h>


#define STACK_SIZE (1024 * 1024)  // 1 MB

int thread_func(void* arg) {
    printf("Thread started: PID = %d, TID = %ld\n", getpid(), syscall(SYS_gettid));
    return 0;
}

int main() {
    printf("Main thread started: PID = %d, TID = %ld\n", getpid(), syscall(SYS_gettid));

    // Выделяем память под стек для потока
    char* stack = malloc(STACK_SIZE);
    if (stack == NULL) {
        perror("malloc");
        exit(1);
    }

    // clone флаги — делаем именно поток (не процесс)
    int flags = CLONE_VM | CLONE_FS | CLONE_FILES | CLONE_SIGHAND | CLONE_THREAD | SIGCHLD;

    // Запускаем поток: стек растёт вниз, поэтому передаём верхнюю границу
    pid_t tid = clone(thread_func, stack + STACK_SIZE, flags, NULL);
    if (tid == -1) {
        perror("clone");
        exit(1);
    }

    sleep(1);  // Дать время потоку завершиться
    printf("Main thread finished\n");
    return 0;
}
