#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

void kadane(float*, int);

int main(int argc, char **argv) {
    int nel;
    scanf("%d", &nel);
    float *mas = (float*)malloc(nel * sizeof(float));
    for (int i = 0; i < nel; i++) {
        float a, b;
        scanf("%f/%f", &a, &b);
        if (a != 0) {
            mas[i] = log(a / b);
        }
        else mas[i] = -100000;
    }
    kadane(mas, nel);
    free(mas);
    return 0;
}

void kadane(float *mas, int nel) {
    int l = 0, r = 0;
    float maxsum = mas[0];
    int start = 0;
    float sum = 0;
    int i = 0;
    if (nel > 1) {
        while (i < nel) {
            sum += mas[i];
            if (sum > maxsum) {
                maxsum = sum;
                l = start;
                r = i;
            }
            i++;
            if (sum < 0) {
                sum = 0;
                start = i;
            }
        }
    }
    else {
        printf("%d %d\n", (nel - 1), (nel - 1));
        return;
    }
    printf("%d %d\n", l, r);
}