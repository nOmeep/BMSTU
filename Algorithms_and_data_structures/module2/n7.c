#include <stdio.h>
#include <stdlib.h>

void merge(int, int, int, int*); //сортировка слиянием

void Insert_sort(int, int, int*); //сортировка вставками

void rec_merge(int, int, int*); //рекусрсивное слияние

int main(int args, char *argv)
{
	int size;
	scanf("%d", &size);
	int *mas = (int*)malloc(size * sizeof(int));
	for (int i = 0; i < size; ++i)
	{
		scanf("%d", &mas[i]);
		//printf("elementi  - %d\n", mas[i]);
	}
	rec_merge(0, size - 1, mas);
	for (int i = 0; i < size; ++i)
	{
		printf("%d ", mas[i]);
	}
	printf("\n");
	free(mas);
	return 0;
}

void rec_merge(int low, int high, int *mas)
{
	if (low < high)
	{
		int med = ((low + high) / 2);
		if (high - low <= 4)
		{
			Insert_sort(low, high, mas);
		}
		else
		{
			rec_merge(low, med, mas);
			rec_merge(med + 1, high, mas);
			merge(low, med, high, mas);
		}
	}
}

void merge(int low, int med, int high, int *mas)
{
	int *helper = (int*)calloc((high - low + 1), sizeof(int));
	int i = low, j = med + 1, h = 0;
	while (h < (high - low + 1))
	{
		if ((j <= high) && ((i == j) || (abs(mas[j]) < abs(mas[i])) || ((i >= med + 1)) ))
		{
			helper[h] = mas[j];
			++j;
		}
		else
		{
			helper[h] = mas[i];
			++i;
		}
		h++;
	}
	for (int i1 = low, i2 = 0; i1 < high + 1; i1++, i2++)
	{
		mas[i1] = helper[i2];
	}
	free(helper);
}

void Insert_sort(int low, int high, int *mas)
{
	int i = low + 1, elem, loc;
	while (i < high + 1)
	{
		elem = mas[i];
		loc = i - 1;
		while ( (loc >= low) && (abs(mas[loc]) > abs(elem)) )
		{
			mas[loc + 1] = mas[loc];
			loc = loc - 1;
		}
		mas[loc + 1] = elem;
		++i;
	}
}














