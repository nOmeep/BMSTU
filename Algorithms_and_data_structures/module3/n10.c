#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct hash hash;

struct hash {
    hash* next;
    int k;
    int v;
};

hash* search(hash*, long);

int lookUp(hash**, long, long);

hash** insert(hash**, int, int, int);

hash* insertBeforeHead(hash*, int, int);

int main(int argc, char **argv) {
    long nel;
    scanf("%ld", &nel);
    long size;
    scanf("%ld", &size);
    hash** hashTable = (hash**)malloc(size * sizeof(hash*));
    for (long i = 0; i < size;) {
        hashTable[i++] = NULL;
    }
    char comand[6];
    for (long i = 0; i < nel; i++) {
        scanf("%s", comand);
        long index;
        long v;
        if (0 == strcmp(comand, "AT")) {
            scanf("%ld", &index);
            //printf("%d\n", index);
            printf("%d\n", lookUp(hashTable, index, size));
        }
        else if (0 == strcmp(comand, "ASSIGN")) {
            scanf("%ld", &index);
            scanf("%ld", &v);
            hashTable = insert(hashTable, index, v, size);
        }
        else {
            printf("как же я устал решать эти задачи, пожалейте :(");
        }
    }
    for(long i = 0; i < size; i++) {
        while (NULL != hashTable[i]) {
            hash* tmp = hashTable[i];
            hashTable[i] = (*hashTable[i]).next;
            free(tmp);
        }
    }
    free(hashTable);
    return 0;
}

hash* search(hash* t, long x) {
    hash* tmp = t;
    //printf("HERE\n");
    while ((NULL != tmp) && ((*tmp).k != x)) tmp = (*tmp).next;
    //printf("IM HERE\n");
    return tmp;
}

int lookUp(hash** t, long k, long m) {
    hash* p = search(t[k % m], k);
    //printf("IVE DONE search\n");
    if (NULL == p) {
        //printf("паникую, ну где вот ошибкааааааа...\n");
        return 0;
    }
    return (*p).v;
}

hash** insert(hash** t, int k, int v, int m) {
    t[k % m] = insertBeforeHead(t[k % m], k, v); 
    return t;
}

hash* insertBeforeHead(hash* t, int k, int v) {
    hash* new = (hash*)malloc(1 * sizeof(hash));
    (*new).next = t;
    (*new).k = k;
    (*new).v = v;
    return new;
}