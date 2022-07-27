#include "../includes/snow.h"
#include "../includes/calculator.h"
#include "../includes/maths.h"

snow_main();

describe(Calculator)
{
    subdesc(BasicArithmetic)
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

    subdesc(NAryOperations)
    {
        int arr1[] = {1, 6, 8, 1, 9};
        int *arr2 = NULL;
        int arr3[] = {0, -1, 10, 7};

        it("Testing sum function")
        {
            asserteq(sum(arr1, 5), 25);
            asserteq(sum(arr2, 1), 0);
            asserteq(sum(arr3, 4), 16);
        }

        it("Testing multiply function")
        {
            asserteq(product(arr1, 5), 432);
            asserteq(product(arr2, 2), 0);
            asserteq(product(arr3, 4), 0);
        }
    }

    subdesc(ComplexArithmeticOperation)
    {
        it("Testing factorial function")
        {
            asserteq(factorial(3), 6);
            asserteq(factorial(0), 1);
            asserteq(factorial(5), 120);
        }

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

    subdesc(ComparisonFunctions)
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
}

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
