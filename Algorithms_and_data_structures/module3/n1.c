#include <stdio.h>
#include <stdlib.h>
#include <string.h>
 
typedef struct Stack stack;
typedef struct Task task;

task pop();

int empty();

void quicksort_stack(int*, int, int);

int partition(int, int, int*);

void push(int, int);

struct Task {
    int low;
    int high;
} t;
 
struct Stack {
    task* data;
	int cap;
	int top;
} s;
 
int main(int argc, char **argv) {
	int n;
	scanf("%d", &n);
    int* mas = (int*)malloc(n * sizeof(int) + 100);
	
    s.top = 0;
	s.cap = n;
	s.data = (task*)malloc(n * sizeof(task));
	
    for(int i = 0; i < n; i++) {
		scanf("%d", &mas[i]);
	}
    
    quicksort_stack(mas, 0, n - 1);
	
    for(int i = 0; i < n; i++) {
		printf("%d ", mas[i]);
	}
	printf("\n");
    free(mas);
	free(s.data);
	return 0;
}

// нерекурсивная быстрая сортировка 
void quicksort_stack(int* mas, int l, int r) { /* нерекурсивная быстрая сортировка */
	task tmp;
	int  cur = 0;
    push(0, r);
	while(empty() > 0) {
		tmp = pop();
		l = tmp.low;
		r = tmp.high;
		cur = partition(l, r, mas);
		if(l < cur) push(l, cur - 1);
		if(cur < r) push(cur+ 1, r);
	}
}
 
int partition(int low, int high, int * mas) { /* граница разделения */
	int cur = low, t = 0;
	for (int j = low; j < high; j++) {
		if(mas[j] < mas[high]) {
			t = mas[cur];
			mas[cur] = mas[j];
			mas[j] = t;
            cur++;
		}
	}
	t = mas[cur];
	mas[cur] = mas[high];
	mas[high] = t;
	return cur;
}

// три операции над стеком  push, stackEmpty, pop 
void push(int l, int r) {  /* добавление элемента в стек */
	if(s.top == s.cap) {
        printf("error \"переполнение\"\n");
		return; /* выдает варнинг, хотя не знаю, как без ретерна тут быть */
	}
	s.data[s.top].low = l;
	s.data[s.top].high = r;
	s.top++;
}

int empty() { /* пустой ли стек */
	if(s.top == 0) return 0;
	return 1;
}

task pop() { /* удаление элемента из стека */
	if(empty() == 0) {
        printf("error опустошение");
        return;
    }
	s.top--;
	return s.data[s.top];
}
 