#include <stdio.h>
#include <stdlib.h>
#include <string.h>
/*Представление бинарного дерева поиска нужно модифицировать 
следующим образом: в каждую вершину нужно добавить поле count, 
содержащее размер поддерева с корнем в данной вершине. 
Размер поддерева – это количество вершин в нём.*/

typedef struct tree tree;

struct tree {
    tree* parent;
    tree* left;
    tree* right;
    int k;
    int count;
    char v[9]; //!!!!!
};

tree* Min(tree*);

tree* Succ(tree*);

tree* Descend(tree*, int); //!!!! long int?

char* Lookup(tree*, int);

tree* Insert(tree*, int);

tree* ReplaceNode(tree*, tree*, tree*);

tree* Delete(tree*, int);

char* Search(tree*, int);

tree* descend(tree*, int);

void finalFree(tree*);

int main(int argc, char **argv) {
    int N;
    scanf("%d", &N);
    tree* t = NULL;
    char comand[6];
    for (int i = 0; i < N; i++) {
        scanf("%s", comand);
        if (0 == strcmp(comand, "INSERT")) {
            int tmp;
            scanf("%d", &tmp);
            t = Insert(t, tmp);
        }
        else if (0 == strcmp(comand, "LOOKUP")) {
            int tmp;
            scanf("%d", &tmp);
            char* find;
            find = Lookup(t, tmp);
            printf("%s\n", find);
        }
        else if (0 == strcmp(comand, "DELETE")) {
            int tmp;
            scanf("%d", &tmp);
            t = Delete(t, tmp);
        }
        else if (0 == strcmp(comand, "SEARCH")) {
            int tmp;
            scanf("%d", &tmp);
            char* res = Search(t, tmp);
            printf("%s \n", res);
        }
    }
    finalFree(t);
    return 0;
}

tree* Min(tree* t) {
    tree* x;
    if (NULL == t) x = NULL;
    else { 
        x = t;
        while (NULL != (*x).left) {
            (*x).count--;
            x = (*x).left;
        }
    }
    return x;
}

tree* Succ(tree* x) {
    tree* y;
    if (NULL != (*x).right) y = Min((*x).right);
    else {
        y = (*x).parent;
        while (NULL != y && x == (*y).right) {
            x = y;
            y = (*y).parent;
        }
    }
    return y;
}

tree* Descend(tree* t, int k) {
    tree* x;
    x = t;
    while (NULL != x && (*x).k != k) {
        if (k < (*x).k) x = (*x).left;
        else x = (*x).right;
    }
    return x;
}

char* Lookup(tree* t, int k) {
    tree* x = Descend(t, k);
    if (NULL == x) {
        printf("аааааа паника\n");
    }
    return (*x).v;
}

tree* Insert(tree* t, int k) {
    tree* y;
    y = (tree*)malloc(1 * sizeof(tree));
    scanf("%s", (*y).v);
    (*y).k = k;
    (*y).parent = NULL;
    (*y).left = NULL;
    (*y).right = NULL;
    (*y).count = 0;
    if (NULL == t) t = y;
    else {
        tree* x = t;
        for(;;) {
            (*x).count++;
            if ((*x).k == k) {
                printf("аааа паникую\n");
            }
            if (k < (*x).k) {
                if (NULL == (*x).left) {
                    (*x).left = y;
                    (*y).parent = x;
                    break;
                }
                x = (*x).left;
            }
            else {
                if (NULL == (*x).right) {
                    (*x).right = y;
                    (*y).parent = x;
                    break;
                }
                x = (*x).right;
            }
        }
    }
    return t;
}

tree* ReplaceNode(tree* t, tree* x, tree* y) {
    if (x == t) {
        t = y;
        if(NULL != y) (*y).parent = NULL;
    }
    else {
        tree* p = (*x).parent;
        if (NULL != y) (*y).parent = p;
        if ((*p).left == x) (*p).left = y;
        else (*p).right = y;
    }
    return t;
}

tree* Delete(tree* t, int k) {
    tree* x = descend(t, k);
    if (NULL == x) {
        printf("паника ааа ааа\n");
    }
    if (NULL == (*x).left && NULL == (*x).right) {
        t = ReplaceNode(t, x, NULL);
    }
    else if (NULL == (*x).left) { //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
        t = ReplaceNode(t, x, (*x).right);
    }
    else if (NULL == (*x).right) {
        t = ReplaceNode(t, x, (*x).left);
    }
    else {
        tree* y = Succ(x);
        t = ReplaceNode(t, y, (*y).right);
        (*(*x).left).parent = y;
        (*y).left = (*x).left;
        if (NULL != (*x).right) (*(*x).right).parent = y;
        (*y).right = (*x).right;
        (*y).count = (*x).count - 1;
        t = ReplaceNode(t, x, y);
    }
    free(x);
    return t;
}

tree* descend(tree* t, int k) {
    tree* x;
    x = t;
    while (NULL != x && (*x).k != k) {
        (*x).count--;
        if (k < (*x).k) x = (*x).left;
        else x = (*x).right;
    }
    return x;
}

char* Search(tree* t, int k) {
    tree* x = t;
    while (NULL != x) {
        if (NULL != (*x).left) {
            if (((*(*x).left).count + 1) > k) x = (*x).left;
            else if (((*(*x).left).count + 1) < k) {
                    k -= ((*(*x).left).count + 2);
                    x = (*x).right;
            }
            else return (*x).v;
        }
        else if (0 != k) {
            k--;
            x = (*x).right;
        }
        else return (*x).v;
    }
}

void finalFree(tree* t) {
    if (NULL != t) {
        finalFree((*t).left);
        finalFree((*t).right);
        free(t);
    }
}






