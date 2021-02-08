#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Queue queue;

struct Queue {
    int cap;
    int count;
    int* heap;
} q;

void insert(int);

void Swap1(int, int);

void Swap2(int, int);

int extractMax();

void heapify(int, int);

int main(int argc, char **argv) {
    int N;
    scanf("%d", &N);
    int M;
    scanf("%d", &M);
    q.heap = (int*)malloc(M * sizeof(int));
    q.cap = M;
    q.count = 0;
    int result = 0;
    int k;
    int k1;
    for (int i = 0; i < N; i++) {
        scanf("%d%d", &k, &k1);
        result = k + k1;
        insert(result);
    }
    for (int i = 0; i < (M - N); i++) {
        scanf("%d%d", &k, &k1);
        result = extractMax();
        if (k > result) {
            result = k + k1;
        } else result += k1;
        insert(result);
    }
    while ( 0 != q.count) {
        result = extractMax();
    }
    printf("%d\n", result);
    free(q.heap);
    return 0;
}

void Swap1(int i, int j) {
    int c = q.heap[i];
    q.heap[i] = q.heap[j];
    q.heap[j] = c;
}

void Swap2(int i, int j) {
    int tmp = q.heap[i];
    q.heap[i] = q.heap[j];
    q.heap[j] = tmp;
}

void insert(int res) {
    int i = q.count;
    q.count++;
    if (i == q.cap) {
        printf("error: Переполнение\n");
        return;
    }
    q.heap[i] = res;
    while ((i > 0) && (q.heap[i] < q.heap[(i - 1) / 2])) {
        Swap2(i, (i - 1) / 2);
        i = (i - 1) / 2;
    }
}

int extractMax() {
    if (0 == q.count) {
        printf("error: Очередь пуста\n");
        return 0;
    }
    int ptr = q.heap[0];
    q.count--;
    if (q.count > 0) {
        q.heap[0] = q.heap[q.count];
        heapify(0, q.count);
    }
    return ptr;
}

void heapify(int i, int n) {
    for(;;) {
        int l = (2 * i + 1), r = (2 * i + 2);
        int j = i;
        if (l < n && q.heap[i] > q.heap[l]) i = l;
        if (r < n && q.heap[i] > q.heap[r]) i = r;
        if (j == i) break;
        Swap1(i, j);
    }
}
