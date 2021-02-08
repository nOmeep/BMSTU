#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void build(int*, int*, int, int, int);

void update(int*, int, int, int, int, int);

int maximum(int*, int, int, int, int, int);

int main(int argc, char** argv) {
    int nel, opcount;
    scanf("%d", &nel); //quantity of elements
    int *arr = (int*)malloc(nel * sizeof(int)); // for elements
    for (int i = 0; i < nel; i++) { //input elements
        scanf("%d", &arr[i]);
    }
    scanf("%d", &opcount); //op quantity
    int *t = (int*)malloc(10 * nel * sizeof(int)); //new tree 
    build(arr, t, 1, 0, (nel - 1));
    char string[4]; //UPD or MAX
    int *result = (int*)malloc(opcount * sizeof(int)); //for output
    int output = 0; // how many outputs will be
    int l, r;
    for (int i = 0; i < opcount; i++) { //doing MAX or UPD opcount times
        scanf("%s %d %d", string, &l, &r);
        if (strcmp(string, "MAX") == 0) { //MAX
            //printf("I WAS HERE\n");
            result[output] = maximum(t, 1, 0, (nel - 1), l, r);
            output++;
        }
        else { // UPD
            update(t, 1, 0, (nel - 1), l, r);
        }
    }
    for (int i = 0; i < output; i++) { //Outup results
        printf("%d\n", result[i]);
    }
    free(t);
    free(result);
    free(arr);
    return 0;
}

void build(int *arr, int* t, int v, int a, int b) {
    if (a == b) {
        t[v] = arr[a];
    }
    else {
        int center = (a + b) / 2;
        build(arr, t, (v * 2), a, center);
        build(arr, t, (v * 2 + 1), (center + 1), b);
        t[v] = (t[v * 2] > t[v * 2 + 1])?t[v * 2]:t[v * 2 + 1];
        //t[v] = t[v * 2] + t[v * 2 + 1];
    }
}

void update(int* t, int v, int a, int b, int pos, int new) {
    if (a == b) {
        t[v] = new;
    }
    else {
        int center = (a + b) / 2;
        if (pos <= center) {
            update(t, (v * 2), a, center, pos, new);
        }
        else {
            update(t, (v * 2 + 1), (center + 1), b, pos, new);
        }
        t[v] = (t[v * 2] > t[v * 2 + 1])?t[v * 2]:t[v * 2 + 1];
        //t[v] = t[v * 2] + t[v * 2 + 1];
    }
}

int maximum(int *t, int v, int a, int b, int l, int r) {
    if (a == l && b == r) {
        return t[v];
    }
    else {
        int center = (a + b) / 2;
        if (r <= center) {
            maximum(t, (v * 2), a, center, l, r);
        }
        else if (l > center) {
            maximum(t, (v * 2 + 1), (center + 1), b, l, r);
        }
        else {
            int r1 = (r < center)?r:center;
            int l1 = (l > (center + 1))?l:(center + 1);
            return (maximum(t, (v * 2), a, center, l, r1) > maximum(t, (v * 2 + 1), (center + 1), b, l1, r))?maximum(t, (v * 2), a, center, l, r1) : maximum(t, (v * 2 + 1), (center + 1), b, l1, r);
        }
    }
}

