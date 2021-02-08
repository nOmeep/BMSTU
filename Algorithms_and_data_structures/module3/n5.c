
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Heap Heap;
typedef struct Queue queue;

void swap(Heap, Heap);

struct Heap {
    int index;
    int k;
    int v;
};

struct Queue {
    Heap* heap;
    int cap;
    int count;
} q;

int queueEmpty();

void insert1(Heap);

void insert2(int**, int*, int);

Heap extractM();

void heapify(int, int);

int main(int argc, char **argv) {
    int n;
    scanf("%d", &n);
    int* masLength = (int*)malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        scanf("%d", &masLength[i]);
    }
    int **mas = (int**)malloc(n * sizeof(int*));
    for (int i = 0; i < n; i++) {
        if (0 != masLength[i]) {
            mas[i] = (int*)malloc(masLength[i] * sizeof(int));
            for (int i1 = 0; i1 < masLength[i]; i1++) {
                scanf("%d", &mas[i][i1]);
            }
        }
    }
    q.heap = (Heap*)malloc(n * sizeof(Heap));
    q.cap = n;
    q.count = 0;
    Heap par;
    for (int i = 0; i < n; i++) {
        if (0 != masLength[i]) {
            par.k = i;
            par.index = 0;
            par.v = mas[i][0];
            insert1(par);
        }
    }
    insert2(mas, masLength, n);
    for (int i = 0; i < n; i++) {
        if (0 != masLength[i]) free(mas[i]);
    }
    printf("\n");
    free(mas);
    free(masLength);
    free(q.heap);
    return 0;
}

int queueEmpty() {
    return (0 == q.count) ? 1 : 0;
}

void insert1(Heap ptr) {
    Heap tmp;
    int i = q.count;
    if (i == q.cap) {
        printf("error: Переполнение\n");
        return;
    }
    q.count++;
    q.heap[i] = ptr;
    while ((i > 0) && (q.heap[(i - 1) / 2].v >= q.heap[i].v)) {
        tmp = q.heap[i];
        q.heap[i] = q.heap[(i - 1) / 2];
        q.heap[(i - 1) / 2] = tmp;
        //swap(q.heap[i], q.heap[(i - 1) / 2]);
        i = (i - 1) / 2;
    }
}

//Эта гадость почему-то не свапает, из-за нее я долго искал ошибку
void swap(Heap a, Heap b) {
    Heap tmp = a;
    a = b;
    b = tmp;
}
// 

Heap extractM() {
    if (0 == q.count) {
        printf("error: Очередь пуста\n");
        return;
    }
    Heap ptr = q.heap[0];
    q.count--;
    if (q.count > 0) {
        q.heap[0] = q.heap[q.count];
        heapify(0, q.count);
    }
    return ptr;
}

void heapify(int i, int n) {
    Heap tmp;
    for(;;) {
        int l = (2 * i + 1), r = (2 * i + 2);
        int j = i;
        if (l < n && q.heap[i].v > q.heap[l].v) i = l;
        if (r < n && q.heap[i].v > q.heap[r].v) i = r;
        if (j == i) break;
        tmp = q.heap[i];
        q.heap[i] = q.heap[j];
        q.heap[j] = tmp;
        //swap(q.heap[i], q.heap[j]);
    }
}

void insert2(int** mas, int* masLength, int n) {
    if (queueEmpty()) {
        //printf("error: Опустошение\n");
        return;
    }
    Heap tmp;
    tmp = extractM();
    printf("%d ", tmp.v);
    if (tmp.index != (masLength[tmp.k] - 1)) {
        tmp.index++;
        tmp.v = mas[tmp.k][tmp.index];
        insert1(tmp);
    }
    else insert2(mas, masLength, n);
    insert2(mas, masLength, n);
}

