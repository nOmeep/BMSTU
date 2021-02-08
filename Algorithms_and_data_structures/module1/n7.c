#include <stdio.h>
#include <malloc.h>

int main(int args, char **argv)
{
	long k, n;
	scanf("%ld%ld", &k, &n);
	char *mas = (char*)malloc((n+1)*sizeof(char));
	for (int i = 2; i < (n + 1); i++)
	{
		mas[i] = 1;
	}
	for (int i = 2; i < (n + 1); i++)
	{
		if (mas[i] == 1)
		{
			for (int j = (2 * i); j < (n + 1); j += i)
			{
				mas[j] = mas[j/i] + 1;
			}
		}
	}
	for (int i = 2; i < (n + 1); i++)
	{
		if (mas[i] == k)
		{
			printf("%d ", i);
		}
	}
	free(mas);
	return 0;
}
