#include "../includes/max3.h"

int max(int a, int b, int c)
{
    int max;

    if (a > b)
    {
        max = a;
    }
    else
    {
        max = b;
    }

    if (max < c)
    {
        max = c;
    }

    return max;
}
