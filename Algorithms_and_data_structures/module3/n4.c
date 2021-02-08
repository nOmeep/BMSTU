#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/*Реализуйте через двойной стек набор операций InitQueue, Enqueue, Dequeue, QueueEmpty и Maximum для работы с очередью целых чисел. 
Операция Maximum возвращает максимальное целое число, в данный момент времени находящееся в очереди. 
Операции Enqueue, QueueEmpty и Maximum должны работать за константное время, 
а операция Dequeue – за амортизированное константное время.*/

#define MINIMUM -2000000001

typedef struct Queue queue;

struct Queue {
    int cap;
    int top1;
    int top2;
    int* data;
    int maximum;
} q;

int empty1();

int empty2();

int queueEmpty();

void push1(int);

void push2(int);

void enqueue(int);

int dequeue();

int pop1();

int pop2();

void findMax();

int main(int argc, char **argv) {
    int n;
    scanf("%d", &n);
    char comand[5];
    q.data = (int*)malloc(n * sizeof(int));
    q.cap = n;
    q.top1 = 0;
    q.top2 = n - 1;
    q.maximum = MINIMUM;
    for (int i = 0; i < n; i++) {
        int tmpNum;
        int delHead;
        scanf("%s", comand);
        if (0 == strcmp(comand, "ENQ")) { //ENQ добавить число x в хвост очереди, -2000000000 < x < 2000000000
            scanf("%d", &tmpNum);
            enqueue(tmpNum);
        }
        else if (0 == strcmp(comand, "DEQ")) { // DEQ удалить головной элемент из очереди | вывести в стандартный поток вывода значение удаляемого головного элемента
            delHead = dequeue();
            printf("%d\n", delHead);
        }
        else if (0 == strcmp(comand, "MAX")) { // MAX показать текущее максимальное число | вывести в стандартный поток вывода текущее максимальное число
            printf("%d\n", q.maximum);
        }
        else if (0 == strcmp(comand, "EMPTY")) { // EMPTY проверить пустоту очереди | вывести в стандартный поток вывода «true» или «false»
            if (queueEmpty()) printf("true\n");
            else printf("false\n");
        }
    }
    free(q.data);
    return 0;
}

int empty1() {
    if (0 == q.top1) return 1;
    return 0;
}

int empty2() {
    if (q.top2 == (q.cap - 1)) return 1;
    return 0;
}

void push1(int x) {
    if (q.top2 < q.top1) {
        printf("error: Переполнение push1\n");
        return;
    }
    q.data[q.top1] = x;
    q.top1++;
    if (x >= q.maximum) q.maximum = x;
}

void push2(int x) {
    if (q.top2 < q.top1) {
        printf("error: Переполнение push2\n");
        return;
    }
    q.data[q.top2] = x;
    q.top2--;
    if (x >= q.maximum) q.maximum =x;
}

int pop1() {
    if (empty1()) {
        printf("error: Опустошение pop1\n");
        return 0;
    }
    q.top1--;
    return q.data[q.top1];
}

int pop2() {
    if (empty2()) {
        printf("error: Опустошение pop2\n");
        return 0;
    }
    q.top2++;
    return q.data[q.top2];
}

int queueEmpty() {
    if (empty1() && empty2()) return 1;
    return 0;
}

void enqueue(int x) {
    push1(x);
}

int dequeue() {
    int x;
    if (empty2()) {
        while(!empty1()) {
            push2(pop1());
        }
    }
    x = pop2();
    if (q.maximum == x) {
        q.maximum = MINIMUM;
        findMax();
    }
    return x;
}

void findMax() {
    for (int i = 0; i < q.top1; i++) {
        if (q.data[i] > q.maximum) q.maximum = q.data[i];
    }    
    for (int i = (q.top2 + 1); i < q.cap; i++) {
        if (q.data[i] > q.maximum) q.maximum = q.data[i];
    }
}