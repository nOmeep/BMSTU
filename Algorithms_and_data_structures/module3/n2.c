#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Stack stack;

struct Stack {
    int* data;
    int cap;
    int top;
} s;

int compare(int, int);

void push(int);

int empty();

int main(int argc, char **argv) {
    int n; 
    int result;
    scanf("%d", &n);
    s.data = (int*)malloc((n + n) * sizeof(int));
    s.top = 0;
    s.cap = n;
    char comand[5];
    for (int i = 0; i < n; i++) {
        int tmpNum;
        int first;
        int second;
        scanf("%s", comand);
        if (0 == strcmp(comand, "CONST")) { // CONST кладёт в стек число x (-1000000000 < x < 1000000000);
            scanf("%d", &tmpNum);
            push(tmpNum);
        }
        else if (0 == strcmp(comand, "ADD")) { //ADD сложение (снимает со стека два операнда a и b и кладёт в стек их сумму);
            first = pop();
            second = pop();
            int sum = (first + second);
            push(sum);
        }
        else if (0 == strcmp(comand, "SUB")) { //SUB вычитание (снимает со стека операнд a, затем снимает со стека операнд b, кладёт в стек a - b);
            first = pop();
            second = pop();
            int diff = (first - second);
            push(diff);
        }
        else if (0 == strcmp(comand, "MUL")) { // MUL умножение (снимает со стека два операнда a и b и кладёт в стек их произведение);
            first = pop();
            second = pop();
            int mult = (first * second);
            push(mult);
        }
        else if (0 == strcmp(comand, "DIV")) { //DIV деление (снимает со стека операнд a, затем снимает со стека операнд b, кладёт в стек результат
            first = pop();
            second = pop();
            int div = (first / second);
            push(div);
        }
        else if (0 == strcmp(comand, "MAX")) { //MAX максимум двух чисел (снимает со стека два операнда a и b и кладёт в стек max(a,b));
            first = pop();
            second = pop();
            if (compare(first, second)) push(first);
            else push(second);
        }
        else if (0 == strcmp(comand, "MIN")) { //MIN минимум двух чисел (снимает со стека два операнда a и b и кладёт в стек min(a,b));
            first = pop();
            second = pop();
            if (compare(first, second)) push(second);
            else push(first);
        }
        else if (0 == strcmp(comand, "NEG")) { //NEG меняет знак числа, находящегося на вершине стека;
            first = pop();
            int negFirst = first * (-1);
            push(negFirst);
        }
        else if (0 == strcmp(comand, "DUP")) { //DUP кладёт в стек копию числа, находящегося на вершине стека;
            first = pop();
            //s.data = (int*)realloc(s.data, (n + 1) * sizeof(int));
            //second = pop();
            push(first);
            push(first);
        }
        else if (0 == strcmp(comand, "SWAP")) { //SWAP меняет местами два числа, находящиеся на вершине стека.
            first = pop();
            second = pop();
            push(first);
            push(second);
        }
    }
    result = pop();
    printf("%d\n", result);
    free(s.data);
    return 0;
}

void push(int t) {
    if (s.top == s.cap) {
        printf("error \"переполнение\"\n");
        return;
    }
    s.data[s.top] = t;
    s.top++;
}

int empty() {
    if (0 == s.top) return 0;
    return 1;
}

int pop() {
    if (0 == empty()) {
        printf("empty\n");
        return;
    }
    s.top--;
    return s.data[s.top];
}

int compare(int a, int b) {
    return (a > b) ? 1 : 0; 
}