#ifndef SOLUTION_CALCULATOR
#define SOLUTION_CALCULATOR

// Include stdio for NULL value
#include <stdio.h>

// Standards operations
int addition(int firstNumber, int secondNumber);
int substraction(int firstNumber, int secondNumber);
int multiplication(int firstNumber, int secondNumber);
double division(int firstNumber, int secondNumber);

// Operations on multiples values
int sum(int *array, int arraySize);
int product(int *array, int arraySize);

// Complex operations
int factorial(int n);
int modulo(int firstNumber, int secondNumber);
int gcd(int firstNumber, int secondNumber);

// Comparison functions
int min(int firstNumber, int secondNumber);
int max(int firstNumber, int secondNumber);

#endif // SOLUTION_CALCULATOR
