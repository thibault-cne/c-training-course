#include <stdio.h>
#include <stdlib.h>

#include "../includes/max3.h"

int main()
{
    char *numberOne = malloc(sizeof(char) * 11);
    char *numberTwo = malloc(sizeof(char) * 11);
    char *numberThree = malloc(sizeof(char) * 11);

    fgets(numberOne, 10, stdin);
    fgets(numberTwo, 10, stdin);
    fgets(numberThree, 10, stdin);

    printf("%d\n", max(atoi(numberOne), atoi(numberTwo), atoi(numberThree)));

    // Free allocated memory
    free(numberOne);
    free(numberTwo);
    free(numberThree);

    return 0;
}
