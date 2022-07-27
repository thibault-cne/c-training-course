#include "../solution/includes/snow.h"
#include "../includes/calculator.h"

snow_main();

describe(FourthQuestion)
{
    it("Testing factorial function")
    {
        asserteq(factorial(3), 6);
        asserteq(factorial(0), 1);
        asserteq(factorial(5), 120);
    }
}
