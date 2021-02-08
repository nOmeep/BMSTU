#include <stdio.h>

int main(int args, char **argv)
{
	long kolvo;
	long n1, n2, n1st, n2st;
	scanf("%ld", &kolvo);
	if (kolvo == 1) /* dlya odnogo el */
	{
		scanf("%ld", &n1st);
		if (n1st == 0)
		{
			n1st = 1;
			printf("%ld\n", n1st);
		}
		else
		{
			n1st = 0;
			printf("%ld %d\n", n1st, 1);
		}
	}
	else if (kolvo == 2) /* dlya 2x el */
	{
		scanf("%ld%ld", &n1st, &n2st);
		if (n1st == 0)
		{
			n1st = 1;
			if ((n2st == 1) && (n1st == 1))
			{
				n2st = 0;
				n1st = 0;
				printf("%ld %ld %d", n1st, n2st, 1);
			}
		}
		else 
		{
			n1st = 0;
			n2st = 1;
			printf("%ld %ld\n", n1st, n2st);
		}
	}
	else if (kolvo == 3) /* dlya 3x el */
	{
		scanf("%ld%ld", &n1st, &n2st);
		if (n1st == 0)
		{
			n1st = 1;
			scanf("%ld", &n1); /* scan 3go el */
			if ((n1 == 0) && (n2st == 1))
			{
				n1st = 0;
				n2st = 0;
				n1 = 1;
				printf("%ld %ld %ld\n", n1st, n2st, n1);
			}
			else
			{
				if ((n1 == 1) && (n2st == 1))
				{
					printf("%ld %d %d %d\n", n1st, 0, 0, 1);
				}
				else 
				{
					printf("%ld %ld %ld\n", n1st, n2st, n1);
				}
			}
		}
		else 
		{
			n1st = 0;
			n2st = 1;
			scanf("%ld", &n1);
			if ((n1 == 1) && (n2st == 1))
			{
				printf("%ld %d %d %d\n", n1st, 0, 0, 1);
			}
			else 
			{
				printf("%ld %ld %ld\n", n1st, n2st, n1);
			}
		} 
	}
	else /* dlya 4x i bolee elementov */
	{
		scanf("%ld%ld", &n1st, &n2st);
		if (n1st == 0)
		{
			n1st = 1;
		}
		else 
		{
			n1st = 0;
			n2st = 1;
		}
		for (long i = 2; i < kolvo; i += 2)
		{
			if (i != (kolvo - 1))
			{
				scanf("%ld%ld", &n1, &n2);
				if ((n2st == 1) && (n1st == 1))
				{
					n1 = 1;
					n2st = 0;
					n1st = 0;
				}
				if ((n1 == 1) && (n2st == 1))
				{
					n2 = 1;
					n1 = 0;
					n2st = 0;
				}
				printf("%ld %ld ", n1st, n2st);
				n1st = n1;
				n2st = n2;
			}
			else 
			{
				printf("%ld ", n1st);
				scanf("%ld", &n1);
				if ((n2st == 1) && (n1st == 1))
				{
					n1 = 1;
					n2st = 0;
					n1st = 0;
				}
				n1st = n2st;
				n2st = n1;
			}			
		}
		if ((n1st == 1) && (n2st == 1))
		{
			n2st = 0;
			n1st = 0;
			printf("%d %d %d\n", 0, 0, 1);
		}
		else
		{
			printf("%ld %ld", n1st, n2st);
		}
	}
	return 0;
}








































