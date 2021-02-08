
#include <stdio.h>
#include <stdlib.h>

//не забыть про байты в обратном порядке и старший
typedef union Int32 {
	int x;
	unsigned char bytes[4];
} Int32;

Int32* DistributionSort(int, Int32*, int); //same as 9 task

Int32* Headbytesort(int, Int32*, int); //для старшего байта

int main(int argc, char **argv) {
	int nel;
	scanf("%d", &nel);
	Int32* mas = (Int32*)malloc((nel + 1) * sizeof(Int32));
	for (int i = 0; i < nel; i++) {
		scanf("%d", &mas[i].x);
	}
	for (int i = 0; i < 4; i++) {
		mas = DistributionSort(nel, mas, i);
	}
	mas = Headbytesort(nel, mas, 3);
	for (int i = 0; i < nel; i++) {
		printf("%d ", mas[i].x);
	}
	printf("\n");
	free(mas);
	return 0;
}

Int32* DistributionSort(int nel, Int32* mas, int nom) {
	char *count = (char*)calloc(256, sizeof(char));
	unsigned long k;
	for (int i = 0; i < nel; i++) {
		k = mas[i].bytes[nom];
		count[k]++;
	}
	for (int i = 1; i < 256; i++) {
		count[i] = count[i] + count[i - 1];
	}
	Int32* d = (Int32*)malloc((nel + 1) * sizeof(Int32*));
	for (int i = 0; i < nel; i++) {
		d[i] = mas[i]; 
	}
	for (int i, j = nel - 1; j >= 0; j--) {
		k = mas[j].bytes[nom];
		i = count[k] - 1;
		count[k] = i;
		d[i].x = mas[j].x;
	}
	free(count);
	free(mas);
	return d;
}

Int32* Headbytesort(int nel, Int32* mas, int nom) {
	char *count = (char*)calloc(256, sizeof(char));
	unsigned long k;
	for (int i = 0; i < nel; i++) {
		k = mas[i].bytes[nom];
		count[128 & k]++;
	}
	for (int i = 254; i >= 0; i--) {
		count[i] = count[i] + count[i + 1];
	}
	Int32* d = (Int32*)malloc((nel + 1) * sizeof(Int32));
	for (int i = 0; i < nel; i++) {
		d[i] = mas[i];
	}
	for (int i, j = nel - 1; j >= 0; j--) {
		k = mas[j].bytes[nom];
		i = count[128 & k] - 1;
		count[128 & k] = i;
		d[i].x = mas[j].x;
	}
	free(count);
	free(mas);
	return d;
}


























