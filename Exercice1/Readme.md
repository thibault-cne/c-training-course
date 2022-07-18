# Exercice 1

Here is the first exercice. The objective is to work with the operation in C. You will build a small calculator then build some more complex functions.

**Question 1.** Create a `calculator.c` file inside the `srcs` folder and a `calculator.h` file inside the `includes` folder.

The `calculator.h` is the header file. It should contain all functions. The `calculator.c` file is the file where you will write all the code for your functions. Make sure those files are linked.

**Question 2.** Write inside the `calculator.c` file 4 functions :

- A function to add two integers with a signature of : `int addition(int firstOperand, int secondOperand);`
- A function to substract two integers with a signature of : `int substraction(int firstOperand, int secondOperand);`
- A function to multiply two integers with a signature of  `int multiplication(int firstOperand, int secondOperand);`
- A function to divide two integers with a signature of `double division(int firstOperand, int secondOperand);`. Be carefull, in case `secondOperand = 0` make sure the function returns 0.

You can now run the `make questionTwo`command and run `build/questionTwo.exe` to check your answers.

**Question 3.** Write inside the `calculator.c` file 2 functions :

- A function to calculate the sum of an array of int (or a pointer to int : `*int`). The function must have a signature of `int sum(int *array);` where `arraySize` is the size of the array. The function must return `0` if the array pointers is `NULL`.
- A function to calculate the product of an array of int (or a pointer to in : `*int`). The function must have the following signature : `int product(int *array, int arraySize);` where `arraySize` is the size of the array. The function must return `0` if the array pointers is `NULL`.

You can now run the `make questionThree`command and run `build/questionThree.exe` to check your answers.

**Question 4.** Write inside the `calculator.c` file a function :

- A function to calculate the factorial of a number. The function must have the following signature : `int factorial(int n);`.

You can now run the `make questionFour`command and run `build/questionFour.exe` to check your answers.

**Question 5.** Write inside the `calculator.c` file 2 functions :

- A function to calculate the rest of the division of 2 integers. The function must have the following signature : `int modulo(int firstOperand, int secondOperand);`.
- A function to calculate the GCD (Geatest Common divisor). The function must have the following signature : `int gcd(int firstOperand, int secondOperand);`.

You can now run the `make questionFive`command and run `build/questionFive.exe` to check your answers.

**Question 6.** Write inside the `calculator.c` file 2 functions :

- A function to calculate the maximum of two integers. The function must have the following signature : `int max(int firstOperand, int secondOperand);`.
- A function to calculate the minimum of two integers. The function must have the following signature : `int min(int firstOperand, int secondOperand);`.

You can now run the `make questionSix`command and run `build/questionSix.exe` to check your answers.

