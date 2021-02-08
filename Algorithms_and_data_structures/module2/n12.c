#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* prefix(char*);

void kmp(char*, char*);

int main(int argc, char **argv) {
    kmp(argv[1], argv[2]);
    return 0;
}

char* prefix(char *s) {
  char *pi = (char*)calloc(strlen(s), sizeof(char));
  int t = 0;
  pi[0] = 0;
  for (int i = 1; i < strlen(s); i++) {
      while (t > 0 && s[i] != s[t]) {
          t = pi[t - 1];
      }
      if (s[t] == s[i]) {
          t++;
      }
      pi[i] = t;
  }
  return pi; 
}

void kmp(char *s, char *t) {
    char* pi = prefix(s);
    int q = 0;
    for (int k = 0; k < strlen(t); k++) {
        while (q > 0 && s[q] != t[k]) {
            q = pi[q - 1];
        }
        if (s[q] == t[k]) {
            q++;
        }
        if (q == strlen(s)) {
            k = k - strlen(s) + 1;
            printf("%d\n", k);
            k = k + strlen(s) - 1;
        }
    }
    free(pi);
}

