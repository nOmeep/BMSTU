#include <stdlib.h> 
#include <stdio.h> 
 
int *array; 

int opyatfib(int n); 
 
int compare(unsigned long i, unsigned long j) 
{ 
        if (i <= j) { 
                printf("COMPARE␣%lu␣%lu\n", i, j); 
        } else { 
                printf("COMPARE␣%lu␣%lu\n", j, i); 
        } 
 
        if (array[i] == array[j]) return 0; 
        return array[i] < array[j] ? -1 : 1; 
} 
 
void swap(unsigned long i, unsigned long j) 
{ 
        if (i <= j) { 
                printf("SWAP␣%lu␣%lu\n", i, j); 
        } else { 
                printf("SWAP␣%lu␣%lu\n", j, i); 
        } 
 
        int t = array[i]; 
        array[i] = array[j]; 
        array[j] = t; 
} 
 
void shellsort(unsigned long, 
        int (*)(unsigned long, unsigned long), 
        void (*)(unsigned long, unsigned long)); 
 
int main(int argc, char **argv) 
{ 
        int i, n; 
        scanf("%d", &n); 
 
        array = (int*)malloc(n * sizeof(int)); 
        for (i = 0; i < n; i++) scanf("%d", array+i); 
 
        shellsort(n, compare, swap); 
        for (i = 0; i < n; i++) printf("%d␣", array[i]); 
        printf("\n"); 
 
        free(array); 
        return 0; 
}

int fibstep(int n)
{
	int n1, n2, n3;
	n1 = 1;
	n2 = 1;
	n3 = 2;
	int doper;
	while (n3 < n)
	{
		n1 = n2;
		n2 = n3;
		n3 = n2 + n1;
	}
	return (n3 - n1);
}

void shellsort(unsigned long nel, 
	int (*compare)(unsigned long i, unsigned long j),
	void (*swap)(unsigned long i, unsigned long j))
{
	int step = nel;
	while (step >= 1)
	{
		step = fibstep(step);
		for (int i = step; i < nel; i++)
		{
			for (int j = i - step; j >= 0 && compare(j, j + step) == 1; j -= step)
			{
				swap(j, j+step);
			}
		}
		if (step == 1) break;
	}
}
























