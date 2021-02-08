#include <stdio.h>
#include <stdlib.h>

void update(int, int, int, int, int*, int*, int);

int query(int*, int, int, int, int, int);

void build(int*, int, int, int, int*, int);

int peak(int, int, int*, int);


// MAIN HERE
int main(int argc, char **argv) {
    int nel, m;
    char s[4];
    scanf("%d", &nel);
    int *mas = (int*)malloc(nel * sizeof(int));
    int *t = (int*)calloc((nel * 10), sizeof(int));
    for (int i = 0; i < nel; i++) {
        scanf("%d", &mas[i]);
    }
    scanf("%d", &m);
    build(mas, 0, 0, nel - 1, t, nel);
    int a, b;
    for (int i = 0; i < m; i++) {
        scanf("%s %d %d", s, &a, &b);
        if (s[0] == 'P') {
            printf("%d\n", query(t, a, b, 0, 0, nel - 1));
        }
        else {
            mas[a] = b;
            update(a, 0, 0, (nel - 1), t, mas, nel);
            if (a > 0) {
                update((a - 1), 0, 0, (nel - 1), t, mas, nel);
            }
            if (a < nel - 1) {
                update((a + 1), 0, 0, (nel - 1), t, mas, nel);
            }
        }
    }
    free(mas);
    free(t);
    return 0;
}

int query (int* t, int l, int r, int v, int a, int b) {
    if (l == a && r == b) {
        return t[v];
    }
    else {
        int m = (a + b) / 2;
        if (r <= m) {
            return query(t, l, r, (2 * v + 1), a, m);
        }
        else if (l > m) {
            return query(t, l, r, (2 * v + 2), (m + 1), b);
        }
        else {
            return (query(t, (m + 1), r, (2 * v + 2), (m + 1), b) + query(t, l, m, (2 * v + 1), a, m));
        }
    }
}

void build(int* mas, int v, int a, int b, int* t, int nel) {
    if (a == b) {
        t[v] = peak(a, b, mas, nel);
    }
    else {
        int m = (a + b) / 2;
        build(mas, (2 * v + 1), a, m, t, nel);
        build(mas, (2 * v + 2), (m + 1), b, t, nel);
        t[v] = peak(a, b, mas, nel);
    }
}

int peak(int left, int right, int * mas, int nel)
{
	if(nel == 1)
	{
		return 1;
	}
	int i = 0;
	int count = 0;
	for(i = left; i <= right; i++)
	{
		if((i == 0 && mas[i+1] <= mas[i]) ||  (i == nel - 1 && mas[i] >= mas[i-1])) 
		{
			count++;
		}
		else if (i != 0 && i != nel-1 && mas[i-1] <= mas[i] && mas[i+1] <= mas[i]) {
			count++;
		}
	}
	return count;
}


void update(int l, int v, int a, int b, int* t, int* mas, int nel) {
    if (a == b) {
        t[v] = peak(l, l, mas, nel);
    }
    else {
        int m = (a + b) / 2;
        if (l <= m) {
            update(l, (2 * v + 1), a, m, t, mas, nel);
        }
        else update(l, (2 * v + 2), (m + 1), b, t, mas, nel);
        t[v] = t[v * 2 + 1] + t[v * 2 + 2];
    }
}



