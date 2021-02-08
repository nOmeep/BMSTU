
#include <stdio.h>

int main(int args, char **argv)
{
	long n, k, s = 0, smax = 0, mas[100000], dopustimA, el;
	scanf("%ld%ld", &n, &k);
	for (int i = 0; i < k; i++)
	{
		scanf("%ld", &dopustimA);
		s += dopustimA;
		mas[i] = dopustimA;
	}
	smax = s;
	for (int i = k; i < n; i++)
	{
		scanf("%ld", &dopustimA);
		s += dopustimA;
		s -= mas[i%k];
		mas[i%k] = dopustimA;
		if (s > smax) 
		{
			smax = s;
			el = i-k+1;
		}
	}
	printf("%ld\n", el);
	return 0;
}
