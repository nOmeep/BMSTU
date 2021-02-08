
#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void query(int**, int, int, int*);

int** build(int*, int, int*);

int my_pow(int, int);

int gcd(int, int);

int* CompLogs(int);

int main(int argc, char **argv) {
    int nel;
    scanf("%d", &nel);
    int *mas = (int*)malloc(nel * sizeof(int));
    for (int i = 0; i < nel; i++) { // скан элементов пос-ти
        scanf("%d", &mas[i]);
    }
    int m;
    scanf("%d", &m); // скан кол-ва операций
    int *lg = CompLogs(20);
    int **st = build(mas, nel, lg);
    free(mas); //!!!
    int l, r;
    for (int i = 0; i < m; i++) {
        scanf("%d%d", &l, &r);
        query(st, l, r, lg);
    }
    for (int i = 0; i < nel; i++) {
        free(st[i]);
    }
    free(st); //!!!
    free(lg); //!!!
    return 0;
}

int* CompLogs(int m) {
    int i = 1, j = 0;
    int *lg = (int*)malloc(100000 * m * sizeof(int));
    while (i < m) {
        while (j < my_pow(2, i)) {
            lg[j] = i - 1;
            j++;
        }
        i++;
    }
    return lg;
}

void query(int** st, int l, int r, int* lg) {
    int j = lg[r - l + 1];
    int v = gcd(st[l][j], st[r - my_pow(2, j) + 1][j]);
    printf("%d\n", v);
}

int gcd(int a, int b) {
    if (0 == b) return abs(a);
    gcd(b, (a % b));
}

/*int gcd(int a, int b) {
    int tmp;
    while (0 != b) {
        tmp = a % b;
        a = b;
        b = tmp;
    }
    return abs(a);
}*/

int** build(int* mas, int nel, int* lg) {
    int m = lg[nel] + 1;
    //int i = 0;
    int **st = (int**)malloc(nel * m * sizeof(int*));
    for (int i = 0; i < nel; i++) {
        st[i] = (int*)malloc(m * sizeof(int));
    }
    for(int i = 0; i < nel; i++) {
        st[i][0] = mas[i];
    }
    //int j = 1;
    for (int j = 1; j < m; j++) {
        //i = 0;
        for (int i = 0; i <= (nel - my_pow(2, j)); i++) {
            st[i][j] = gcd(st[i][j - 1], st[i + my_pow(2, j - 1)][j - 1]);
            }
    }
    return st;
}

int my_pow(int a, int b) {
    if (0 == b) {
        return 1;
    }
    if (0 == (b % 2)) {
        return my_pow(a * a, b / 2);
    }
    return a * my_pow(a, b - 1);
}