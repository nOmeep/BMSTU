
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int gcd(int, int);

void prefind(char*);

int main(int argc, char **argv) {
    prefind(argv[1]);
    return 0;
}

int gcd(int a, int b) {
    if (0 == b) {
        return abs(a);
    }
    gcd(b, a % b);
}

void prefind(char* str) {
    int t;
    char *pi = (char*)calloc(strlen(str), sizeof(char));
    t = 0;
    for (int i = 1; i < strlen(str); i++) {
        while (t > 0 && str[i] != str[t]) {
            t = pi[t - 1];
        }
        if (str[t] == str[i]) {
            t++;
        }
        pi[i] = t;
    }
    for (int i = 1; i < strlen(str) + 1; i++) {
        int flag = 0;
        if (pi[i - 1] != 0) {
            flag = 0;
            t = gcd(pi[i -1], i);
            for (int j = 0; j + t < i; j++) {
                if (str[j] != str[t + j]) {
                    flag = 1;
                }
            }
            if (flag == 0) {
                printf("%d %d\n", i, (i / gcd(pi[i - 1], i)));
            }
        }
    }
    free(pi);
}