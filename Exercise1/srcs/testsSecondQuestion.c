#include "../solution/includes/snow.h"
#include "../includes/calculator.h"

snow_main();

describe(SecondQuestion)
{
    it("Testing addition function.")
    {
        asserteq(addition(1, 1), 2);
        asserteq(addition(3, 2), 5);
        assertneq(addition(1, 7), 5);
    }

    it("Testing substraction function")
    {
        asserteq(substraction(1, 1), 0);
        asserteq(substraction(5, 2), 3);
        assertneq(substraction(2, 5), 3);
        asserteq(substraction(2, 5), -3);
    }

    it("Testing multiplication function")
    {
        asserteq(multiplication(1, 1), 1);
        asserteq(multiplication(5, 2), 10);
        assertneq(multiplication(2, 5), 6);
        asserteq(multiplication(2, 5), 10);
    }

    it("Testing division function")
    {
        asserteq(division(1, 1), 1);
        asserteq(division(5, 2), 2.5);
        assertneq(division(2, 5), 1);
        asserteq(division(2, 5), 0.4);
        asserteq(division(10, 0), 0);
    }
}
