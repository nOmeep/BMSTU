
#include <stdio.h>

int main(int args, char **argv)
{
	long n, k, x0, kf, itog = 0;
	scanf("%ld%ld%ld", &n, &k, &x0);
	for(int i = n; i >= k; i--)
	{
		scanf("%ld", &kf);
		for(int j = 0; j < k; j++)
		{
			kf *= (i-j);	
		}
		itog += kf;
		if (i != k) 
		{
			itog *= x0;
		}
	}
	printf("%ld\n", itog);
	return 0;	
}

