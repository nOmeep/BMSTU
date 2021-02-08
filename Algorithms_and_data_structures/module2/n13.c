#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* prefix(char*);

int find_word (char*, char*);

int main(int argc, char **argv)  {
    int flag = find_word(argv[1], argv[2]);
    if (1 == flag) {
        printf("no\n");
    }
    else printf("yes\n");
    return 0;
}

char* prefix(char* s) {
    char *pi = (char*)calloc(strlen(s), sizeof(char));
    int t = 0;
    pi[0] = 0;
    for (int i = 1; i < strlen(s); i++) {
        while (t > 0 && s[t] != s[i]) {
            t = pi[t - 1];
        }
        if (s[t] == s[i]) {
            t++;
        }
        pi[i] = t;
    }
    return pi;
}

int find_word(char* s1, char* s2) {
    char *pi = (char*)malloc(10000 * sizeof(char));
    strcat(strcpy(pi, s1), s2);
    //strcpy(pi, s1);
    //strcat(pi, s2);
    //printf("%ld", strlen(pi));
    pi = prefix(pi);
    int flag = 0;
    for (int i = strlen(s1); i < (strlen(s1) + strlen(s2)); i++) {
        if (0 == pi[i]) {
            flag = 1;
            break;
        }
    }
    free(pi);
    return flag;
}

