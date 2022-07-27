#include "../solution/includes/snow.h"
#include "../includes/calculator.h"

snow_main();

describe(FifthQuestion)
{
    it("Testing min function")
    {
        asserteq(min(5, 2), 2);
        asserteq(min(7, 5), 5);
        asserteq(min(-1, 7), -1);
        asserteq(min(5, 5), 5);
    }

    it("Testing max function")
    {
        asserteq(max(10, 11), 11);
        asserteq(max(-30, -10), -10);
        asserteq(max(50, 0), 50);
        asserteq(max(4, 4), 4);
    }
}
