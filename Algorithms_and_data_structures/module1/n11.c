#include <stdio.h>

int compare(unsigned long i)
{
	if (i == 91052112271683405UL) return 0;
	if (i < 91052112271683405UL) return -1;
	return 1;
}

unsigned long binsearch(unsigned long, int (*)(unsigned long));

int main(int argc, char **argv)
{
	printf("%lu\n", binsearch(11742432862790734555UL, compare));
	return 0;
}

unsigned long binsearch(unsigned long nel, int (*compare)(unsigned long i))
{
	unsigned long index = -1, left = 0, right = nel, center;
	while ((right - left) > 1)
	{
		center = (right + left)/2;
		if (compare(center) == 0)
		{
			index = center;
			break;
		}
		else if (compare(center) == 1)
		{
			right = center;
		}
		else
		{
			left = center;
		}
	}
	if (index == -1)
	{
		return nel;
	}
	else 
	{
		return index;
	}
}
