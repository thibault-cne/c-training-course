#include "../includes/pointers.h"
#include "../includes/snow.h"

snow_main();

describe(QuestionTwo)
{
    it("Testing add function")
    {
        int a;
        int b;

        a = 5;
        b = 7;
        asserteq(add(&a, &b), 12);
        asserteq(add(&b, &a), 12);

        a = 0;
        b = 0;
        asserteq(add(&a, &b), 0);
        asserteq(add(&b, &a), 0);

        a = -1;
        b = 10;
        asserteq(add(&a, &b), 9);
        asserteq(add(&b, &a), 9);
    }
}
