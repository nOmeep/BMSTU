#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int Overlap(char *s1, char *s2) /* покрытие */
{
	//printf("OVERLAP-IN\n");
	int s1_lenth = strlen(s1), s2_lenth = strlen(s2), overlap = 0;
	int s1_last = s1_lenth - 1;
	int i = s1_last, j = 1;
	char *helper;
	while (i > 0 && j < s2_lenth)
	{
		//printf("COME WHILE\n");
		char *pref = (char*)malloc(70 * sizeof(char));
		char *suff = (char*)malloc(70 * sizeof(char));
		pref[j] = 0;
		strncpy(pref, s2, j);
		helper = &s1[i];
		strcpy(suff, helper);
		if (strcmp(pref, suff) == 0)
		{
			//printf("IF\n");
			overlap = j;
		}
		free(pref);
		free(suff);
		--i;
		++j;
	}
	//printf("OVERLAP - %d\n", overlap);
	return overlap;
}

void merge(char *s1, char *s2, char *super, int lenth) /* слияние строк */ 
{
	//printf("MERGE-IN\n");
	char *s22; //*s11;
	char *temp = (char*)malloc(70 * sizeof(char));
	s22 = &s2[lenth];
	strcpy(temp, s22);
	strcpy(super, s1);
	strcat(super, temp);
	//printf("MERGED STRING %s\n", super);
	free(temp);
	//printf("MERGE-OUT\n");
}

int CreateSuperString(char **mas, int nel) /* создание суперстроки */
{
	int len = nel, current_overlap, i1 = 0, j1 = 0, correct_string, returner/*counter_helper = 1*/; /* условие для вайла, текущий оверлап, индексы обнуления, индекс нужной строки для длины, возващаемое значение */
	char *s1, *s2; /* указатели на строчки */
	while (len > 1)
	{
		//printf("WHILE_IN ITERATION - %d\n", counter_helper);
		int max_overlap = 0; /* обнулять в каждой итерации */
		char *super = (char*)malloc(70 * sizeof(char)); /* вспомогательный массив, для слияния строчек */
		if (super == NULL) /* проверка, аналогичная 106 строке */
		{
			return -1;
		}
		for (int i = 0; i < nel; i++) 
		{
			if (strlen(mas[i]) == 0)
			{
				continue;
			}
			else
			{
				for (int j = 0; j < nel; j++)
				{
					if (strlen(mas[j]) == 0 || i == j)
					{
						continue;
					}
					else
					{
						current_overlap = Overlap(mas[i], mas[j]); /* рассматриваемые сейчас */
						if ((current_overlap > max_overlap) || (max_overlap == 0)) /* сранение с максимумом */
						{
							//printf("SWAP MAXOVERLAP\n");
							max_overlap = current_overlap; 
							s1 = mas[i];
							s2 = mas[j];
							i1 = i; 
							j1 = j;
						}
					}
				}
			}
		}
		//printf("HE WAS HERE, MISTAKE DOWN\n");
		merge(s1, s2, super, max_overlap); /* слияние с1 и с2 в супер */
		//printf("HE ALSO WAS HERE\n");
		//printf("FORMING - %s\n", super);
		strcpy(mas[i1], super); /* запоминание супера в массив строк */
		//printf("THATS NORM TOO\n");
		mas[j1][0] = 0; /* обнуление */
		free(super); 
		--len;
		//printf("WHILE-OUT ITERATION - %d\n", counter_helper);
		//counter_helper++;  
	}
	for (int i = 0; i < nel; i++) /* поиск слитой строки */
	{
		//printf("FOR-IN OK\n");
		if (mas[i][0] != 0) 
		{
			correct_string = i;
		}
		//printf("FOR-END OK\n");
	}
	returner = strlen(mas[correct_string]);
	return returner;
}

int main(int args, char *argv)
{
	int quantity, super_lenth; /* количество строк, длина суперстроки */ 
	scanf("%d", &quantity); /*сканиуем кол-во строк */
	char *mas1 = (char*)malloc(quantity * 10000 * sizeof(char)); /* память на количество строк */
	if (NULL == mas1) /* проверка памяти */
	{
		return -1;
	}
	char **mas = (char**)malloc(10000 * sizeof(char*)); /* память для самих строчек */
	if (NULL == mas) /* аналогично 106 строке */
	{
		return -1;
	}
	for (int i = 0; i < quantity; i++) /* заполнение строчек */
	{
		mas[i] = mas1 + (i * quantity * 100);
		scanf("%s", mas[i]);
		//printf("STROKA VVEDENNAYA %s\n", mas[i]);
	}
	super_lenth = CreateSuperString(mas, quantity); /* итоговая длина кратч. суперстроки */
	printf("%d\n", super_lenth); /* вывод длины суперстроки */
	free(mas1);
	free(mas);
	return 0;
}



























