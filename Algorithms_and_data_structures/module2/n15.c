#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

void update(int*, int, char, int, int, int);

void build(char*, int*, int, int, int);

int query(int*, int, int, int, int, int);

int main(int argc, char** argv) {
    char *mas = (char*)malloc(1000001 * sizeof(char));
    scanf("%s", mas);
    int strnel = strlen(mas);
    int *t = (int*)calloc(strnel * 10, sizeof(int));
    int m;
    scanf("%d", &m);
    char action[3];
    char *tmp = (char*)malloc(1000001 * sizeof(char));
    build(mas, t, 0, 0, (strnel - 1));
    for (int i = 0; i < m; i++) {
        scanf("%s", action);
        if ('H' == action[0]) {
            int a, b;
            scanf("%d %d", &a, &b);
            int flag = query(t, a, b, 0, 0, (strnel - 1));
            if ((flag & (flag - 1)) == 0) {
                printf("YES\n");
            }
            else printf("NO\n");
        }
        else {
            int a;
            scanf("%d %s", &a, tmp);
            int tmpnel = strlen(tmp);
            int x = strnel / log2(strnel);
            memcpy(mas + a, tmp, tmpnel);
            if (x < tmpnel) {
                build(mas, t, 0, 0, (strnel - 1));
            }
            else {
                for (int i = 0; i < tmpnel; i++) {
                    update(t, a + i, tmp[i], 0, 0, (strnel - 1));
                }
            }
        }
    } 
    free(mas);
    free(t);
    free(tmp);
    return 0;
}

void build(char* mas, int* t, int v, int a, int b){
    if (a == b) {
        t[v] = 1 << (mas[a] - 97);
    }
    else {
        int m = (a + b) / 2;
        build(mas, t, (2 * v + 1), a, m);
        build(mas, t, (2 * v + 2), (m + 1), b);
        t[v] = t[v * 2 + 1] ^ t[v * 2 + 2];
    }
}

int query(int* t, int l, int r, int v, int a, int b) {
    int res = 0;
    if (l == a && r == b) {
        res = t[v];
    }
    else {
        int m = (a + b) / 2;
        if (r <= m) {
            res = query(t, l, r, (2 * v + 1), a, m); 
        }
        else if (l > m) {
            res = query(t, l, r, (2 * v + 2), (m + 1), b);
        } 
        else {
            res = (query(t, l, m, (2 * v + 1), a, m) ^ query(t, (m + 1), r, (2 * v + 2), (m + 1), b));
        }
    }
    return res;
}

void update(int* t, int j, char new, int v, int a, int b) {
    int m = (a + b) / 2;
    if (a == b) {
        t[v] = 1 << (new - 97);
    }
    else {
        if (j <= m) {
            update(t, j, new, (2 * v + 1), a, m);
        }
        else {
            update(t, j, new, (2 * v + 2), (m + 1), b); 
        }
        t[v] = t[v * 2 + 1] ^ t[v * 2 + 2];
    }
}