#include "../includes/calculator.h"

int addition(int a, int b)
{
    return a + b;
}

int substraction(int a, int b)
{
    return a - b;
}

int multiplication(int a, int b)
{
    return a * b;
}

double division(int a, int b)
{
    if (b == 0)
    {
        return 0;
    }

    return (double)((double)a / (double)b);
}

int sum(int *arr, int arrSize)
{
    // Verify the NULL case
    if (arr == NULL)
    {
        return 0;
    }

    int i = 0;
    int result = 0;

    while (i < arrSize)
    {
        result += arr[i++];
    }

    return result;
}

int product(int *arr, int arrSize)
{
    // Verify the NULL case
    if (arr == NULL)
    {
        return 0;
    }

    int i = 0;
    int result = 1;

    while (i < arrSize)
    {
        result *= arr[i++];
    }

    return result;
}

int factorial(int n)
{
    if (n == 0 || n == 1)
    {
        return 1;
    }

    return n * factorial(n - 1);
}

int modulo(int a, int b)
{
    int quotient = a / b;

    return a - quotient * b;
}

int gcd(int a, int b)
{
    int rest = a % b;

    while (rest != 0)
    {
        a = b;
        b = rest;

        rest = a % b;
    }

    return b;
}

int min(int firstNumber, int secondNumber)
{
    if (firstNumber > secondNumber)
    {
        return secondNumber;
    }

    return firstNumber;
}

int max(int firstNumber, int secondNumber)
{
    if (firstNumber < secondNumber)
    {
        return secondNumber;
    }

    return firstNumber;
}
