#include <stdio.h>
#include <stdlib.h>

int Partition(int, int, int*); 

void swap(int, int, int*);

void select_sort(int, int, int*);

void quicksort(int, int, int, int*);

int main(int args, char *argv)
{
	int size, m; 
	scanf("%d%d", &size, &m);
	int *mas = (int*)calloc(size, sizeof(int));
	for (int i = 0; i < size; i++)
	{
		scanf("%d", &mas[i]);
	}
	quicksort(0, (size - 1), m, mas);
	for (int i = 0; i < size; i++)
	{
		printf("%d ", mas[i]);
	}
	printf("\n");
	free(mas);
	return 0;
}

void quicksort(int low, int high, int m, int *mas)
{
	while (high - low > 0)
	{
		int q;
		q = Partition(low, high, mas);
		if (high - low > m)
		{
			if ((high - q) > (q - low)) //меньшая подпоследовательность
			{
				quicksort(low, (q - 1), m, mas);
				low = q + 1;
			}
			else
			{
				quicksort(q + 1, high, m, mas);
				high = q - 1;
			}
		}
		else
		{
			select_sort(low, high, mas);
			break;
		}
	}
}

void swap(int i, int j, int *mas)
{
	int help = mas[i];
	mas[i] = mas[j];
	mas[j] = help;
}

void select_sort(int low, int high, int *mas)
{
	int j = high;
	while (j > low)
	{
		int k = j;
		int i = j - 1;
		while (i >= 0)
		{
			if (mas[k] < mas[i])
			{
				k = i;
			}
			i--;
		}
		swap(j, k, mas);
		j--;
	}
}

int Partition(int low, int high, int *mas)
{
	int i = low;
	int j = low;
	while (j < high)
	{
		if (mas[j] < mas[high])
		{
			swap(i, j, mas);
			i++;
		}
		j++;
	}
	swap(i, high, mas);
	return i;
}












































