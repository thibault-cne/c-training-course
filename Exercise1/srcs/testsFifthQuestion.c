#include "../solution/includes/snow.h"
#include "../includes/calculator.h"

snow_main();

describe(FifthQuestion)
{
    it("Testing modulo function")
    {
        asserteq(modulo(5, 3), 2);
        asserteq(modulo(10, 10), 0);
        asserteq(modulo(10, 2), 0);
        asserteq(modulo(11, 7), 4);
    }

    it("Testing gcd function")
    {
        asserteq(gcd(5, 3), 1);
        asserteq(gcd(10, 2), 2);
        asserteq(gcd(10, 5), 5);
        asserteq(gcd(14, 21), 7);
    }
}
