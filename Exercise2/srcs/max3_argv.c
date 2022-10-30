#include <stdio.h>
#include <stdlib.h>

#include "../includes/max3.h"

int main(int argc, char **argv)
{
    if (argc != 4)
    {
        return 0;
    }

    int i = 0;

    printf("%d\n", max(atoi(argv[i + 1]), atoi(argv[i + 2]), atoi(argv[i + 3])));

    return 0;
}
