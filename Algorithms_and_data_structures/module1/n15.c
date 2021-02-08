#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main(int args, char **argv)
{
	long height = 0, width = 0, avrwidth = 0, avrheight = 0, textlenth = 0;
	if (args != 4) /* wrong input */
	{
		printf("Usage: frame <height> <width> <text>");
		return 0;
	}
	else /* correct input */
	{
		height = atoi(argv[1]);
		width = atoi(argv[2]);
		textlenth = strlen(argv[3]);
		if (((width - textlenth)/2 <= 0) || (height < 3)) /* wrong znachenia dlya str and td */
		{
			printf("Error");
			return 0;
		}
		else /* rabota s pechatyu */
		{
			avrheight = (height - 1) / 2; /* srednie positions */
			avrwidth = (width - textlenth) / 2; /* sredniye positions */
			for (int i = 0; i < height; i++) 
			{
				if ((i == 0) || (i == (height - 1))) /* print stars na first and last str */
				{
					for (int i1 = 0; i1 < width; i1++)
					{
						printf("*");
					}
				}
				else
				{
					if (i != avrheight) /* ne vstretili str s nadpisyu */
					{
						printf("*");
						for (int i2 = 0; i2 < width - 2; i2++)
						{
							printf(" ");
						}
						printf("*");
					}
					else /* vstretilas str s nadpisyu */ 
					{
						printf("*");
						for (int i3 = 0; i3 < avrwidth - 1; i3++)
						{
							printf(" ");
						}
						printf("%s", argv[3]);
						for (int i4 = avrwidth + textlenth; i4 < width - 1; i4++)
						{
							printf(" ");
						}
						printf("*");
					}
				}
				printf("\n");
			}
			
		}
	}
	return 0;
}
