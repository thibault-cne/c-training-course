#include "../includes/pointers.h"

void swap(int *a, int *b)
{
    int c = *a;

    *a = *b;
    *b = c;
}

int add(int *a, int *b)
{
    return *a + *b;
}
