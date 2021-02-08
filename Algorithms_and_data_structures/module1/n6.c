#include <stdio.h>

int main(int args, char **argv)
{
	long m, n, el[4000], saddY = -1, saddX = -1, current;
	char flag = 0;
	scanf("%ld%ld", &m, &n);
	long maxstr = -4294967296;
	for (int i = 0; i < n; i++)
	{
		scanf("%ld", &el[i]);
		if (maxstr == el[i])
		{
			flag = 0;
		}
		if (el[i] > maxstr)
		{
			saddY = i;
			saddX = 0;
			maxstr = el[i];
			flag = 1;
		}
	}
	for (int i = 1; i < m; i++)
	{
		maxstr = -4294967296;
		for (int j = 0; j < n; j++)
		{
			scanf("%ld", &current);
			if ((current >= maxstr) && (saddX == i))
			{
				flag = 0;
			}
			if ((current <= el[j]) && (saddY == j))
			{
				flag = 0;
			}
			if ((current > maxstr) && (current < el[j]))
			{
				saddX = i;
				saddY = j;
				flag = 1;
			}
			if (current > maxstr)
			{
				maxstr = current;
			}
			if (current < el[j])
			{
				el[j] = current;
			}
		}
	}
	if (flag != 0)
	{
		printf("%ld %ld\n", saddX, saddY);
	}
	else
	{
		printf("none\n");
	}
	return 0;
}
















