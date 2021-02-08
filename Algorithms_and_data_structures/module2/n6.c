#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void Heapify(void*, size_t, size_t, size_t, int (*compare)(const void *a, const void *b));

void BuildHeap(void*, size_t, size_t, int (*compare)(const void *a, const void *b));

int compare(const void*, const void*);

void swap(const void*, const void*, int);

void hsort(void *base, size_t nel, size_t width,
	int (*compare)(const void *a, const void *b));

int main(int args, char **argv)
{	
	int n;
	scanf("%d", &n);
	char strs[n][1000];
	for (int i = 0; i < n; i++)
	{
		scanf("%s", strs[i]);
	}	
	hsort(strs, n, 1000, compare);
	for (int i = 0; i < n; i++)
	{
		printf("%s\n", strs[i]);
	}
	return 0;
}

void hsort(void *base, size_t nel, size_t width,
	int (*compare)(const void *a, const void *b)) /* 64 page from presentation */
{
	/* P <- Idn chto eto takoe??? */
	BuildHeap(base, width, nel, compare);
	for (int i = nel - 1; i > 0; i--)
	{
		swap((char*)base, (char*)base + i * width, width);
		Heapify(base, width, i, 0, compare);
	}
}

void Heapify(void *base, size_t width, size_t nel, size_t i, int (*compare)(const void *a, const void *b)) /* heapify from 61 page presentation */
{
	int l, r, j;
	for(;;)
	{
		l = 2*i + 1;
		r = l + 1;
		j = i;
		if (l < nel && compare((char*)base + width * i, (char*)base + width * l) == -1)
		{
			i = l;
		}
		if (r < nel && compare((char*)base + width * i, (char*)base + width * r) == -1)
		{
			i = r;
		}
		if (i == j)
		{
			break;
		}
		swap((char*)base + width * i, (char*)base + width * j, width);
	}
}

void BuildHeap(void *base, size_t width, size_t nel, int (*compare)(const void *a , const void *b)) /* from 62 presentation */
{
	int i;
	i = (nel / 2) - 1;
	while (i >= 0)
	{
		Heapify(base, width, nel, i, compare);
		i--;
	}
}

int compare(const void *a, const void *b)
{
	int counta1 = 0, counta2 = 0;
	char *p, *p1;
	p = a;
	p1 = b;
	for (int i = 0; i < strlen(a); i++)
	{
		if (*(p + i) == 'a')
		{
			counta1++;
		}
	}
	for (int i = 0; i < strlen(b); i++)
	{
		if (*(p1 + i) == 'a')
		{
			counta2++;
		}
	}
	return (counta1 < counta2)?-1:1;
}

void swap(const void *a, const void *b, int width)
{
	char c;
	for (int i = 0; i < width; i++)
	{
		c = *(char*)(a + i);
		*(char*)(a + i) = *(char*)(b + i);
		*(char*)(b + i) = c;
	}
}































