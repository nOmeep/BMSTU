#include <stdio.h>
#include <malloc.h>

int main(int args, char **argv)
{
	int n1, n2;
	scanf("%d", &n1);
	int *p1 = (int*)malloc(n1*sizeof(int));
	for (int i = 0; i < n1; i++)
	{
		scanf("%d", &p1[i]);
	}
	scanf("%d", &n2);
	int *p2 = (int*)malloc(n2*sizeof(int));
	for (int i = 0; i < n2; i++)
	{
		scanf("%d", &p2[i]);
	}
	int i = 0, j = 0;
	while ((i != n1) && (j != n2))
	{
		if (p1[i] < p2[j])
		{
			printf("%d ", p1[i]);
			i++;
		}
		else
		{
			printf("%d ", p2[j]);
			j++;
		}
	}	
	if (i == n1)
	{
		while (j != n2)
		{
			printf("%d ", p2[j]);
			j++;
		}
	}
	if (j == n2)
	{
		while (i != n1)
		{
			printf("%d ", p1[i]);
			i++;
		}
	}
	free(p1);
	free(p2);
	return 0;
}
