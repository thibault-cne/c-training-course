#include "../includes/snow.h"
#include "../includes/dayofbirth.h"

snow_main();

describe(dayfromdateTesting)
{
    it("Testing dayfromdate with normal dates")
    {
        asserteq(dayfromdate(28, 7, 2022), 4);
        asserteq(dayfromdate(23, 3, 2001), 5);
    }

    it("Testing dayfromdate with unexisting dates")
    {
        asserteq(dayfromdate(32, 7, 2022), -1);
        asserteq(dayfromdate(28, 13, 2022), -1);
    }
}
