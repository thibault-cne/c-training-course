#include "../includes/pointers.h"
#include "../includes/snow.h"

snow_main();

describe(QuestionOne)
{
    it("Testing swap function")
    {
        int a;
        int b;

        a = 5;
        b = 7;
        swap(&a, &b);
        asserteq(b, 5);
        asserteq(a, 7);

        a = 0;
        b = 0;
        swap(&a, &b);
        asserteq(b, 0);
        asserteq(a, 0);

        a = -1;
        b = 10;
        swap(&a, &b);
        asserteq(a, 10);
        asserteq(b, -1);
    }
}
