#include "../includes/snow.h"
#include "../includes/max3.h"

snow_main();

describe(QuestionOne)
{
    it("Testing max function")
    {
        asserteq(max(5, 3, 4), 5);
        asserteq(max(10, 10, -1), 10);
        asserteq(max('a', 'b', 10), 'b');
        asserteq(max(11, 7, -22), 11);
        asserteq(max(-20, -30, 0), 0);
    }
}
