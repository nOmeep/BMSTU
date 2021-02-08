#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* Реализуйте операции InitQueue, QueueEmpty, Enqueue и Dequeue для очереди целых чисел,
представленной в виде кольцевого буфера. Начальный размер буфера – 4. 
В случае переполнения размер буфера должен увеличиваться в два раза. */

typedef struct Queue queue;

struct Queue {
    int tail;
    int head;
    int cap;
    int count;
    int* data;
} q;

int empty();

void enqueue(long);

int dequeue();

int main(int argc, char **argv) {
    int length = 4;
    char comand[5];
    int n;
    scanf("%d", &n);
    q.data = (int*)malloc(n * sizeof(int));
    q.cap = length;
    q.head = 0;
    q.tail = 0;
    q.count = 0;
    for (int i = 0; i < n; i++) {
        long tmpNum;
        //int tmp;
        scanf("%s", comand);
        if (0 == strcmp(comand, "ENQ")) { //ENQ добавить число x в хвост очереди, -2000000000 < x < 2000000000
            scanf("%ld", &tmpNum);
            if (q.cap == q.count) {
                length = length * 2;
                q.cap = length;
                q.data = (int*)realloc(q.data, (length * sizeof(int)) );
                memcpy(q.data + q.cap / 2, q.data, (q.cap / 2) * sizeof(int));
                q.tail += q.cap / 2;
            }
            enqueue(tmpNum);
        }
        else if (0 == strcmp(comand, "DEQ")) { //DEQ удалить головной элемент из очереди
            //tmp = dequeue();
            printf("%d\n", dequeue());
        }
        else if (0 == strcmp(comand, "EMPTY")) { //EMPTY проверить пустоту очереди
            if (0 != empty()) printf("false\n");
            else printf("true\n");
        }
    }
    free(q.data);
    return 0;
}

int empty() {
    return q.count;
}

void enqueue(long x) {
    if (q.count == q.cap) {
        printf("Переполнение ерор\n");
        return;
    }
    q.data[q.tail] = x;
    q.tail++;
    if (q.tail == q.cap) q.tail = 0;
    q.count++;
}

int dequeue() {
    int x;
    if (0 == empty()) {
        printf("Опустошение\n");
        return 0;
    }
    x = q.data[q.head];
    q.head++;
    if (q.head == q.cap) q.head = 0;
    q.count--;
    return x;
}

