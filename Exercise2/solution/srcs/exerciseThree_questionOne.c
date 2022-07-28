#include "../includes/snow.h"
#include "../includes/dayofbirth.h"

snow_main();

describe(QuestionOne)
{
    it("Testing dayfromdate function")
    {
        asserteq(dayfromdate(28, 7, 2022), 4);
        asserteq(dayfromdate(23, 3, 2001), 5);
        asserteq(dayfromdate(32, 7, 2022), -1);
        asserteq(dayfromdate(28, 13, 2022), -1);
    }
}
