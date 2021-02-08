
#include <stdio.h>
#include <malloc.h>

int counter = 0;

void search(int n, int sum, int position, int *mas)
{
	if (n > position)
	{
		sum += mas[position];
		if ((sum & (sum - 1)) == 0 && sum != 0)
		{
			counter = counter + 1;
		}
		search(n, sum, (position + 1), mas);
		sum -= mas[position];
		search(n, sum, (position + 1), mas);
	}
}

int main(int args, char **argv)
{
	int n, sum = 0, position = 0;
	scanf("%d", &n);
	int *mas = (int*)malloc(n * sizeof(int));
	for (int i = 0; i < n; i++)
	{
		scanf("%d", &mas[i]);
	}
	search(n, sum, position, mas);
	printf("%d\n", counter);
	free(mas);
	return 0;
}
