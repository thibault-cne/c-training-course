#include "../solution/includes/snow.h"
#include "../includes/maths.h"

snow_main();

describe(Maths)
{
    it("Testing countAllIntegerOccurencies function")
    {
        asserteq(countAllIntegerOccurencies(12, 1), 5);
        asserteq(countAllIntegerOccurencies(30, 1), 13);
        asserteq(countAllIntegerOccurencies(15, 3), 2);
        asserteq(countAllIntegerOccurencies(30, 2), 13);
    }
}
