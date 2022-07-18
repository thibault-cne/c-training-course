#include "../solution/includes/snow.h"
#include "../includes/calculator.h"

snow_main();

describe(ThirdQuestion)
{
    int arr1[] = {1, 6, 8, 1, 9};
    int *arr2 = NULL;
    int arr3[] = {0, -1, 10, 7};

    it("Testing sum function")
    {
        asserteq(sum(arr1, 5), 25);
        asserteq(sum(arr2, 1), 0);
        asserteq(sum(arr3, 4), 16);
    }

    it("Testing multiply function")
    {
        asserteq(product(arr1, 5), 432);
        asserteq(product(arr2, 2), 0);
        asserteq(product(arr3, 4), 0);
    }
}
