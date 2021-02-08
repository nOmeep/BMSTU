
#include <stdio.h>
#include <stdlib.h>

/*void daysort(int, int, int*);

void monthsort(int, int, int*);

void yearsort(int, int, int*); */

typedef struct DATE 
{
	int Day, Month, Year/*1-31, 1-12, 1970-2030*/;
} DATE;

DATE* datesort(int, int, DATE*);

int main(int argc, char **argv)
{
	int nel;
	scanf("%d", &nel);
	DATE* date = (DATE*)malloc((nel + 1) * sizeof(DATE));
	for (int i = 0; i < nel; i++)
	{
		scanf("%d", &date[i].Year);
		scanf("%d", &date[i].Month);
		scanf("%d", &date[i].Day);
	}
	date = datesort(nel, 0, date);//days
	date = datesort(nel, 1, date);//months
	date = datesort(nel, 2, date); //years
	for (int i = 0; i < nel; i++)
	{
		printf("%d %d %d\n", date[i].Year, date[i].Month, date[i].Day);
	}
	free(date);
	return 0;
}

DATE* datesort(int nel, int indicator,  DATE* date) //sort за линейное время, distribution sort
{
	int current;
	if (0 == indicator)
	{
		current = 32;
	}
	if (1 == indicator)
	{
		current = 13;
	}
	if (2 == indicator)
	{
		current = 61;
	}
	int *count = (int*)calloc(current, sizeof(int));
	int k;
	for (int i = 0; i < nel; i++)
	{
		if (0 == indicator)
		{
			k = date[i].Day;	
		}
		if (1 == indicator)
		{
			k = date[i].Month;
		}
		if (2 == indicator)
		{
			k = date[i].Year - 1970;//for correct result
		}
		count[k]++;
	}
	for (int i = 1; i < current; i++)
	{
		count[i] = count[i] + count[i - 1];
	}
	DATE* D = ( DATE*)calloc((nel + 1), sizeof( DATE));
	for (int i = 0; i < nel; i++)
	{
		D[i].Day = date[i].Day;
		D[i].Month = date[i].Month;
		D[i].Year = date[i].Year;
	}
	int j = nel - 1;
	int i;
	while (j >= 0)
	{
		if (0 == indicator)
		{
			k = date[j].Day;
		}
		else if (1 == indicator)
		{
			k = date[j].Month;
		}
		else
		{
			k = date[j].Year - 1970;
		}
		i = count[k] - 1;
		count[k] = i;
		D[i].Day = date[j].Day;
		D[i].Month = date[j].Month;
		D[i].Year = date[j].Year;
		j--;
	}
	free(date);
	free(count);
	return D;
}













