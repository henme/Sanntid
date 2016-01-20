// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread

#include <pthread.h>
#include <stdio.h>

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
int i = 0;
// Note the return type: void*
void* someThreadFunction(){
    pthread_mutex_lock(&mutex);
	for(int j = 0; j < 1000000; j++) i += 1;
    pthread_mutex_unlock(&mutex);
    return NULL;
}

void* someThreadFunction2(){
    pthread_mutex_lock(&mutex);
	for(int j = 0; j < 1000000; j++) i -= 1;
    pthread_mutex_unlock(&mutex);
    return NULL;
}

int main(){
    pthread_t someThread;
    pthread_t someThread2;
    pthread_create(&someThread, NULL, someThreadFunction, NULL);
    pthread_create(&someThread2, NULL, someThreadFunction2, NULL);
    // Arguments to a thread would be passed here ---------^
    
    pthread_join(someThread, NULL);
    pthread_join(someThread2, NULL);
    printf("Result: %d \n", i);
    return 0;
    
}
