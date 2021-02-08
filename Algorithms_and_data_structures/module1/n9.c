#include <stdio.h>

int array[] = {
        100000000,
        200000000,
        300000000,
        400000000,
        500000000
};

void revarray(void*, unsigned long, unsigned long);

int main(int argc, char **argv)
{
        revarray(array, 5, sizeof(int));

        int i;
        for (i = 0; i < 5; i++) {
                printf("%d\n", array[i]);
        }

        return 0;
}

void revarray(void *base, unsigned long nel, unsigned long width)
{
	for (int i = 0; i < (nel/2); i++)
	{
		char c;
		for (int j = 0; j < width; j++)
		{
			c = *(char*)base;
			*(char*)base = *(char*)(base + (nel * width - width - 2*i*width));
			*(char*)(base + (nel * width - width - 2*i*width)) = c;
			(char*)base++;
		}	
	}	
}

