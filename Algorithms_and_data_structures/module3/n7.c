#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Elem elem;

struct Elem { 
    struct Elem *next; 
    char *word; 
}; //sp;

elem* bsort(elem*);

elem* bubblesort(int, elem*);

int ListLength(elem*);

int ListEmpty(elem*);

elem* Insert(elem*, char*);

int main(int argc, char **argv) {
    char* strings = (char*)malloc(10000 * sizeof(char));//free
    //scanf("%s", strings);
    gets(strings);
    elem* list;
    list = (elem*)malloc(sizeof(elem)); //free
    (*list).word = (char*)malloc(100 * sizeof(char));
    (*list).next = NULL;
    elem* beg = list;
    for (int i = 0; strings[i] != 0;) {
        int word_index = 0;
        char* one_word = (char*)malloc(100 * sizeof(char));//free
        while (32 == strings[i]) {
            i++;
        }
        while (32 != strings[i] && 0 != strings[i]) {
            one_word[word_index] = strings[i];
            i++;
            word_index++;
        }
        one_word[word_index] = 0;
        list = Insert(list, one_word);
        while (32 == strings[i]) {
            i++;
        }
        free(one_word);
    }
    list = beg;
    list = (*list).next;
    free((*beg).word);
    free(beg);
    list = bsort(list);
    for (;NULL != list;) {
        //printf("I am here\n");
        printf("%s ", (*list).word);
        free((*list).word);
        list = (*list).next;
    }
    printf("\n");
    free(strings);
    return 0;
}

elem* Insert(elem* list, char* one_word) {
    elem* new_list;
    new_list = (elem*)malloc(10000 * sizeof(elem));
    (*new_list).word = (char*)malloc(100 * sizeof(char));
    (*new_list).next = NULL;
    strcpy((*new_list).word, one_word);
    (*list).next = new_list;
    return new_list;
}

int ListLength(elem* list) {
    int length = 0;
    elem* x = list;
    while (1 != ListEmpty(x)) { //!!!!!!!!!!!!!!!!!
        length++;
        x = (*x).next;
    }
    return length;
}

int ListEmpty(elem* list) {
    if (NULL == list) return 1;
    return 0;
}

elem* bsort(elem* list) {
    return bubblesort(ListLength(list), list);
}

elem* bubblesort(int nel, elem* list) {
    int t = nel - 1;
    elem* tmp;
    int bound;
    char* c;
    int i;
    while (t > 0) {
        bound = t;
        t = 0;
        i = 0;
        tmp = list;
        while (i < bound) {
            if (strlen((*tmp).word) > strlen((*(*tmp).next).word)) {
                c = (*tmp).word;
                (*tmp).word = (*(*tmp).next).word;
                (*(*tmp).next).word = c;
                t = i;
            }
            i++;
            tmp = (*tmp).next;
        }
    }
    return list;
}