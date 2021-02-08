#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct table table;

struct table {
    //table* next;
    int k;
    long v;
};

int main(int argc, char** argv) {
    int nel;
    scanf("%d", &nel);
    long* mas = (long*)malloc(nel * sizeof(long));
    for (int i = 0; i < nel;) {
        scanf("%ld", &mas[i++]);
    }
    table* mas1 = (table*)calloc(nel, sizeof(table));
    for (int i = 0; i < nel; i++) {
        mas1[i].k = 1; 
    }
    int res = 0;
    int f = nel - 1;
    long chek = 0;
    for (int i = 0; i < nel; i++) {
        chek ^= mas[i];
        int tmp = chek % nel;
        int i1 = tmp;
        while (1) {
            if (0 == mas1[i1].v) {
                mas1[i1].v = chek;
                if (!chek) {
                    res = res + mas1[i1].k;
                    mas1[i1].k++;
                } break;
            }
            else if (chek == mas1[i1].v) {
                res = res + mas1[i1].k;
                mas1[i1].k++;
                break;
            }
            else if (i1 == f) {
                i1 = 0;
                continue;
            }
            i1++;
        }
    }
    free(mas1);
    printf("%d\n", res);
    free(mas);
    return 0;
}