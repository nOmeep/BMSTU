#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void csort(char*, char*);

int wordcount(char*); /* counter of words in str */

char fillwithnull(char*, char*, char*, char*, int nel); /*fill mas with nulls*/  

int main(int args, char **argv) /* MAIN TUT */
{
	char src[1000];
	char dest[1000];
	gets(src);
	csort(src, dest);
	for (int i = 0; dest[i] != 0; i++)
	{
		printf("%c", dest[i]);
	}
	return 0;
}

void csort(char *src, char *dest) /* sama f sortirovky */ 
{
		int n = wordcount(src);
		char *mas = (char*)malloc(n * sizeof(char));
		char *mas1 = (char*)malloc(n * sizeof(char));
		char *dop = (char*)malloc(n * sizeof(char));
		char *count = (char*)malloc(n * sizeof(char));
		fillwithnull(mas, count, mas1, dop, n);
		/*fillwithnull(mas1, n);
		fillwithnull(dop, n);
		fillwithnull(count, n);*/
		int i = 0, foundword = 0, i1 = 0;
		while (src[i] != 0)
		{
			if (src[i] == 32 && foundword != 0)
			{
				dop[i1] = foundword;
				foundword = 0;
				i1++;
			}
			else if (src[i] != 32)
			{
				foundword++;
			}
			i++;
		}
		dop[i1] = foundword;
		i = i1 = 0;
		while (src[i] != 0)
		{
			if (src[i] != 32)
			{
				count[i1] = src[i];
				mas[i1] = i;
				i += dop[i1];
				i1++;
			}
			else 
			{
				i++;
			}
		}
		int j = 0;
		i = 0;
		while (i < n) /* countsort from presentation */
		{
			j = i + 1;
			while (j < n)
			{
				if (dop[j] < dop[i])
				{
					mas1[i] += 1;
				}
				else 
				{
					mas1[j] += 1;
				}
				j++;
			}
			i++; 
		}
		i = j = 0;
		int nom = 0;
		for (i = 0; i < n; i++)
		{
			if (mas1[i] == j)
			{
				for (int j1 = mas[i]; j1 < mas[i] + dop[i]; j1++)
				{
					dest[nom] = src[j1];
					nom++;
				}
				if ((j + 1) != n) /* prisvoenie probela */
				{
					dest[nom] = 32;
					nom++; 
				}
				j++;
				i = -1;
			}
		}
		dest[nom] = 0;
		free(mas);
		free(mas1);
		free(dop);
		free(count);
}

int wordcount(char *str) /* counter of words in str */
{
	int counter = 0, counter1 = 0;
	int n = strlen(str);
	for (int i = 0; i < n; i++)
	{
		if ( (isgraph(*(str + i)) != 0) && (isgraph(*(str + i + 1)) == 0) ) /* skip probelov i podschet slov */
		{
			counter++;
		}
		counter1++;
	}
	if (isgraph(*(str + counter1)) == 1) /* proverka na ende */
	{
		counter++;
	}
	return counter;
}

char fillwithnull(char *mas, char *mas1, char *mas2, char *mas3, int nel) /* to je ruchnoye zapolneniye null */
{
	for (int i = 0; i < nel; i++)
	{
		mas[i] = 0;
		mas1[i] = 0;
		mas2[i] = 0;
		mas3[i] = 0;
	}
}
























