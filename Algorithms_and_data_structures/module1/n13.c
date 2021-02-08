#include  <stdio.h>

int strdiff(char *a, char *b);

int main(int argc, char **argv)
{
        char s1[1000], s2[1000];
        gets(s1);
        gets(s2);
        printf("%d\n", strdiff(s1, s2));

        return 0;
}
int strdiff(char *a, char *b)
{
	int j = 0, counter = 0;
	while ((a[j] != 0) || (b[j] != 0))
	{
		for (int i = 0; i < 8; i++)
		{
			if ((a[j] & (1 << i)) == (b[j] & (1 << i)))
			{
				counter++;
			}
			else 
			{
				return counter;
			}
		}
		j++;
	}
	/*if (a[j] != 0 && b[j] == 0 || a[j] == 0 && b[j] != 0)
	{
		return counter+1;
	}*/
	return (a[j] != 0 && b[j] == 0 || a[j] == 0 && b[j] != 0)?counter:-1;
}
