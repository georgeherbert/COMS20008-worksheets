#include <pthread.h>
#include <stdio.h>

void *hello_world(void *args) {
    int *n = args;
    printf("Hello from thread %d\n", *n);
    pthread_exit(NULL);
}

int main(int argc, char const *argv[]) {
    pthread_t thread[5];
    int threadNum[5] = {1, 2, 3, 4, 5};

    for (int i = 0; i < 5; i++) {
        if (pthread_create(&thread[i], NULL, hello_world, &threadNum[i])) {
            printf("Error creating thread\n");
        }
    }

    for (int i = 0; i < 5; i++) {
        if (pthread_join(thread[i], NULL)) {
            printf("Error joining thread\n");
        }
    }
}
