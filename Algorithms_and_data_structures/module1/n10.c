#include  <stdio.h>

int array[] = {
        100000000,
        400000000,
        300000000,
        500000000,
        200000000
};

int compare(void  *a, void  *b)
{
        int va =  *(int*)a;
        int vb =  *(int*)b;
        if (va == vb) return 0;
        return va  < vb ? -1 : 1;
}

int maxarray(void*, unsigned long, unsigned long,
        int (*)(void  *a, void  *b));

int main(int argc, char  **argv)
{
        printf("%d\n", maxarray(array, 5, sizeof(int), compare));
        return 0;
}
int maxarray(void *base, unsigned long nel, unsigned long width,
	int (*compar)(void *a, void *b))
{
	
	char *el, *max;
	int index = 0;
	for (int i = 1; i < nel; i++)
	{
		max = (char*)base + index * width;
		el = (char*)base + i * width;
		if (compare(max, el) < 0)
		{
			index = i;
		}
	}
	return index;
	
}
