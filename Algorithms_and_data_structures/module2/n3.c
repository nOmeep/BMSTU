#include <stdlib.h> 
#include <stdio.h> 
 
int *array; 
 
int compare(unsigned long i, unsigned long j) 
{ 
        if (i <= j) { 
                printf("COMPARE %lu %lu\n", i, j); 
        } else { 
                printf("COMPARE %lu %lu\n", j, i); 
        } 
 
        if (array[i] == array[j]) return 0; 
        return array[i] < array[j] ? -1 : 1; 
} 
 
void swap(unsigned long i, unsigned long j) 
{ 
        if (i <= j) { 
                printf("SWAP %lu %lu\n", i, j); 
        } else { 
                printf("SWAP %lu %lu\n", j, i); 
        } 
 
        int t = array[i]; 
        array[i] = array[j]; 
        array[j] = t; 
} 
 
void bubblesort(unsigned long, 
        int (*)(unsigned long, unsigned long), 
        void (*)(unsigned long, unsigned long)); 
 
int main(int argc, char **argv) 
{ 
        int i, n; 
        scanf("%d", &n); 
 
        array = (int*)malloc(n * sizeof(int)); 
        for (i = 0; i < n; i++) scanf("%d", array+i); 
 
        bubblesort(n, compare, swap); 
        for (i = 0; i < n; i++) printf("%d␣", array[i]); 
        printf("\n"); 
 
        free(array); 
        return 0; 
}

void bubblesort(unsigned long nel,
	int(*compare)(unsigned long i, unsigned long j),
	void (*swap)(unsigned long i, unsigned long j))
{
	int t = 0, bound, bound1 = 0, i = 0, i1 = 0;
	bound = nel -1;
	while (t < bound)
	{
		while (i < bound)
		{
			if (compare(i, i + 1) == 1)
			{
				swap(i, i+1);
				t = i;
			}
			i++;
		}
		i1 = i = t;
		while (i > bound1)
		{
			if (compare(i - 1, i) == 1)
			{
				swap(i-1, i);
				i1 = i;
				
			}
			i--;
		}
		bound1 = i1;
		bound = t;
		t = i1;
		i = i1;
	}	
}




























