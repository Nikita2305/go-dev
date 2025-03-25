// file: threads.c
#include <pthread.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/syscall.h>

void* thread_func(void* arg) {
    printf("Thread started: PID = %d, TID=%ld\n", getpid(), syscall(SYS_gettid));
    return NULL;
}

int main() {
    printf("Main thread started: PID = %d, TID=%ld\n", getpid(), syscall(SYS_gettid));
    pthread_t tid;
    pthread_create(&tid, NULL, thread_func, NULL);
    pthread_join(tid, NULL);
    printf("Thread finished\n");
    return 0;
}
