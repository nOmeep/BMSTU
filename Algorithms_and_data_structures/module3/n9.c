#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define SYMBNUM 6
#define FIRSTNUM 48
#define LASTNUM 57

//long global_res[2];

typedef struct pair pair;

typedef struct avlTree avlTree;

//typedef struct pair pair;

struct pair {
    int value;
    int index_i;
};

struct avlTree {
    avlTree* parent;
    avlTree* l;
    avlTree* r;
    int k;
    int v;
    int b;
};

int veryLongIf(char);

pair findConst(char*, int);

avlTree* Insert(avlTree**, int, int);

avlTree *insertAVL (avlTree*, int, int);

avlTree *RotateLeft(avlTree*, avlTree*);

avlTree *RotateRight(avlTree*, avlTree*);

avlTree* ReplaceNode(avlTree*, avlTree*,avlTree*);

pair findVar(char*, int);

avlTree* LookUp(avlTree*, int);

int SpecIndex(char*, char);

void finalIndex(avlTree*);

void finalFree(avlTree* t) {
    if (NULL != t) {
        finalFree((*t).l);
        finalFree((*t).r);
        free(t);
    }
}

int main(int argc, char** argv) {
    int nel;
    scanf("%d\n", &nel);
    avlTree* tree = NULL;
    int valIndex = 0;
    char* string = (char*)malloc((nel + nel) * sizeof(char));
    char* specs = (char*)malloc(SYMBNUM * sizeof(char));
    specs[0] = '+';
    specs[1] = '-';
    specs[2] = '*';
    specs[3] = '/';
    specs[4] = '(';
    specs[5] = ')';
    gets(string);
    for (int i = 0; string[i] != '\0'; i++) {
        //для пробелов
        if (SpecIndex(specs, string[i]) == -1) {
            continue; 
        }
        //для "+ - * / ( )"
        else if (SpecIndex(specs, string[i]) == 0) {
            printf("SPEC 0\n");
            continue;
        }
        else if (SpecIndex(specs, string[i]) == 1) {
            printf("SPEC 1\n");
            continue;
        }
        else if (SpecIndex(specs, string[i]) == 2) {
            printf("SPEC 2\n");
            continue;
        }
        else if (SpecIndex(specs, string[i]) == 3) {
            printf("SPEC 3\n");
            continue;
        }
        else if (SpecIndex(specs, string[i]) == 4) {
            printf("SPEC 4\n");
            continue;
        }
        else if (SpecIndex(specs, string[i]) == 5) {
            printf("SPEC 5\n");
            continue;
        }
        // константа
        else if ( (string[i] >= FIRSTNUM) && (string[i] <= LASTNUM) ) {
            pair constRES;
            constRES = findConst(string, i);
            printf("CONST %d\n", constRES.value);
            i = constRES.index_i - 1;
            continue;
        }
        //
        else {
            pair identRES;
            identRES = findVar(string, i);
            i = identRES.index_i;
            avlTree* tmp1 = LookUp(tree, identRES.value);
            if ((NULL == tmp1) || (*tmp1).k != identRES.value) {
                printf("IDENT %d\n", valIndex);
                tree = insertAVL(tree, identRES.value, valIndex);
                valIndex++;
                i--;
            }
            else {
                printf("IDENT %d\n", (*tmp1).v);
                i--;
            }
            //global_res[1] = NULL;
            //global_res[0] = NULL;
        }
    }
    free(string);
    //if (NULL != tree) {
      //  free((*tree).l);
        //free((*tree).r);
        //free(tree);
    //}
    //free(tree);
    finalFree(tree);
    free(specs);
    return 0;
}

int SpecIndex(char* speclist, char symbol) {
    for (int i = 0; i < SYMBNUM; i++) {
        if (symbol == speclist[i]) {
            return i;
        }
        else if (symbol == ' ') {
            return -1;
        }
    }
}

pair findConst(char* str, int i) {
    pair resultConst;
    resultConst.value = 0;
    char* t = (char*)malloc(50 * sizeof(char));
    int i1 = 0;
    for (;;) {
        if (veryLongIf(str[i])) {
            t[i1] = str[i];
            i1++;
            i++;
        }
        else break;
    }
    resultConst.index_i = i;
    i1--;
    for (int j = i1, pr = 1; j > -1; j--) {
        resultConst.value += (t[j] - 48) * pr;
        pr = pr * 10;
    }
    free(t);
    return resultConst;
}

int veryLongIf(char s) {
    if ((s != ' ') && (s != '\0') && (s != '+' != 0) && (s != '-') && (s != '*') && (s != '/') && (s != '(') && (s != ')')) {
        return 1;
    }
    return 0;
}

avlTree* Insert(avlTree** t, int k, int value) {
    avlTree* y = (avlTree*)malloc(sizeof(avlTree));
    (*y).parent = NULL;
    (*y).l = NULL;
    (*y).r = NULL;
    (*y).k = k;
    (*y).v = value;
    if (NULL == *t) *t = y;
    else {
        avlTree* x = *t;
        for(;;) {
            if ((*x).k == k) {
                printf("аааа паникую\n");
            }
            if (k < (*x).k) {
                if (NULL == (*x).l) {
                    (*x).l = y;
                    (*y).parent = x;
                    break;
                }
                x = (*x).l;
            }
            else {
                if (NULL == (*x).r) {
                    (*x).r = y;
                    (*y).parent = x;
                    break;
                }
                x = (*x).r;
            }
        }
    }
    return y;
}

avlTree* insertAVL (avlTree* t, int k, int val) {
    avlTree* a;
    a = Insert(&t, k, val);
    (*a).b = 0;
    while(1) {
        avlTree* x = (*a).parent;
        if (x == NULL) {
            break;
        }
        if (a == (*x).l) {
            (*x).b--;
            if ((*x).b == 0) {
                break;
            }
            if ((*x).b == -2) {
                if ((*a).b == 1) {
                    t = RotateLeft(t, a);
                }
                t = RotateRight(t, x);
                break;
            }
        }
        else {
            (*x).b++;
            if ((*x).b == 0) {
                break;
            }
            if ((*x).b == 2) {
                if ((*a).b == -1) {
                    t = RotateRight(t,a);
                }
                t = RotateLeft(t, x);
                break;
            }
        }
        a = x;
    }
    //free(a);
    return t;
}

avlTree *RotateLeft(avlTree* t, avlTree* x) {
    avlTree* y = (*x).r;
    t = ReplaceNode(t, x, y);
    avlTree* b = (*y).l;
    if (b != NULL) {
        (*b).parent = x;
    }
    (*x).r = b;
    (*x).parent = y;
    (*y).l = x;
    (*x).b--;
    if ((*y).b > 0) {
        (*x).b = (*x).b - (*y).b;
    }
    (*y).b--;
    if ((*x).b < 0) {
        (*y).b = (*y).b + (*x).b;
    }
    return t;
}

avlTree *RotateRight(avlTree* t, avlTree* x) {
    avlTree* y;
    y = (*x).l;
    t = ReplaceNode(t, x, y);
    avlTree* b;
    b = (*y).r;
    if (b != NULL) {
        (*b).parent = x;
    }
    (*x).l = b;
    (*x).parent = y;
    (*y).r = x;
    (*x).b++;
    if ((*y).b < 0) {
        (*x).b -= (*y).b;
    }
    (*y).b++;
    if ((*x).b > 0) {
        (*y).b += (*x).b;
    }
    return t;
}

avlTree* ReplaceNode(avlTree* t, avlTree* x,avlTree* y) {
    if (x == t) {
        t = y;
        if (y != NULL) {
            (*y).parent = NULL;
        }
    }
    else {
        avlTree *p;
        p = (*x).parent;
        if (y != NULL) {
            (*y).parent = p;
        }
        if ((*p).l == x) {
            (*p).l = y;
        }
        else {
            (*p).r = y;
        }
    }
    return t;
}

pair findVar(char* str, int i) {
    pair resVar;
    resVar.value = 0;
    for (int pr = 1;;) {
        if (veryLongIf(str[i])) {
            resVar.value += (str[i] * pr);
            i++;
            pr = pr * 36;
        }
        else {
            break;
        }
    }
    resVar.index_i = i;
    return resVar;
}

avlTree* LookUp(avlTree* t, int k) {
    avlTree *x;
    x = t;
    while ((NULL != x) && (*x).k != k) {
        if (k < (*x).k) {
            x = (*x).l;
        }
        else {
            x = (*x).r;
        }
    }
    return x;
}







