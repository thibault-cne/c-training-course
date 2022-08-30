# Exercise 2

## Table of contents

1. [Requirements](#requirements)
2. [Exercises](#exercises)
    - [Exercise 1](#exercise-1--hello-world)
    - [Exercise 2](#exercise-2--argumentsstdoutstdin)
    - [Exercise 3](#exercise-3--argumentsstdinstdout)
    - [Exercise 4](#exercise-4--pointers)
    - [Exercise 5](#exercise-5--stringh)

### Requirements

For this exercise, `golang` will be needed to run the tests. You can download it directly from the official golang website : [official golang page](https://go.dev/doc/install).For this exercise, `golang` will be needed to run the tests. You can download it directly from the official golang website : [official golang page](https://go.dev/doc/install). You will also need to install `stdbuf` command. It is inside the `coreutils GNU` package. You can get it here : [stdbuf installation guide](https://command-not-found.com/stdbuf).

You can run all the golang tests inside the `tests` folder with the following command :

``` bash
go run main.go -questions
```

It will automatically search for executable files inside the `build` folder and test those files.

### Exercises

#### Exercise 1 : Hello World

- Create a file called `helloworld.c` inside the `srcs` folder.
- Inside this file, create a program that display on the terminal the message : Hello World !!

Help : you will need to create a `main()` function and use the `printf` function. Use the man of the function with the command `man 3 printf`.

To compile your file you can use the given makefile with the command `make questionOneOne` or use the following commands :

``` bash
mkdir -p ./objs | clang -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/helloworld.c -o objs/helloworld.o

clang -Wall -Wextra -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls  ./objs/helloworld.o -o ./build/questionOneOne
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions
```

#### Exercise 2 : arguments/stdout/stdin

In this exercise you will create a function that returns the max of 3 integers. You will retrieve the 3 integers from the arguments of the programs and you will send the result on the stdout (terminal prompt).

**Question 1.**

- Create a `max3.c` file inside the `srcs` folder
- Create a `max3.h` file inside the `includes` folder and link this file to `max3.c`
- Inside this file create a function with the signature `int max(int a, int b, int c);` thats return the maximum between 3 integers.

You can test your function with the given snow tests. You can run it with the makefile with `make questionTwoOne` and then `./build/questionTwoOne`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3.c -o objs/max3.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/exerciseTwo_questionOne.c -o objs/exerciseTwo_questionOne.o

clang objs/max3.o objs/exerciseTwo_questionOne.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./build/questionTwoOne
```

**Question 2.**

- Create a `max3_argv.c` file inside the `srcs` folder
- Create a main function to retrieve the 3 integers from the call arguments and display the result on the stdout.

Make sure that the program end if their is more or less than 3 values entered.

Help : You can use the `atoi` functions to convert string to integer. You can check the manual of the function with `man 3 atoi`.

You can compile your code with the makefile : `make questionTwoTwo` or you can also use the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3.c -o objs/max3.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3_argv.c -o objs/max3_argv.o

clang objs/max3.o objs/max3_argv.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./build/questionTwoTwo
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions
```

**Question 3.**

- Create a `max3_stdin.c` file inside the `srcs` folder
- Create a main function to calculate the maximum between 3 integers. The user will enter all numbers in the CLI (stdin). Each lines correspond to a number.

Help : you can use the `fgets` function to retrieve the use input. You can get all the information on the function with the command `man 3 fgets`.

You can compile your code with the makefile : `make questionTwoThree` or you can also use the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3.c -o objs/max3.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3_stdin.c -o objs/max3_stdin.o

clang objs/max3.o objs/max3_stdin.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./build/questionTwoThree
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions
```

#### Exercise 3 : arguments/stdin/stdout

- Create a file called `dayofbirth.c` inside the `srcs` folder and a file `dayofbirth.h` inside the `includes` folder.

**Question 1.**

- Create a function that for a given date return the day of the date. The signature of the functions is : `int dayfromdate(int day, int month, int year);`. For exemple, `dayfromdate(28, 7, 2022)` returns `4` because the 28/07/2022 is a thursday and it is the fourth day of the week. Make sure that if it the input is not a date, the function returns `-1`

Help : you can use the Zeller formula (1885) :

> ω = d + ⌊2.6 × t - 0.2⌋ + e  + ⌊e / 4⌋ + ⌊c / 4⌋ + 5 × c
>
>Where :
>
> - d : number of the day in the month,
> - m : number of the month in the year,
> - y : number of the year,
> - ⌊x⌋ means the integer part of x,
> - t = m - 2 if m > 2 else m + 10
> - b = y if m > 2 else y - 1
> - c = ⌊b / 100⌋
> - e = b - 100 × c
>
>   ω ≡ 0 (mod 7) if it's sunday, 1 (mod 7) if it's monday, etc.

You can test your function with the given snow tests. You can run it with the makefile with `make questionThreeOne` and then `./build/questionThreeOne`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/dayofbirth.c -o objs/dayofbirth.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/exerciseThree_questionOne.c -o objs/exerciseThree_questionOne.o

clang objs/dayofbirth.o objs/exerciseThree_questionOne.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./build/questionThreeOne
```

**Question 2.**

- Create a file called `dayofbirth_argv.c` inside the `srcs` folder.
- Create a main function to retrieve the name and the birthdate of a user and display is day of birth. For exemple, `./questionThreeTwo Thibault C. 23/03/2001` prints `Thibault C. was born on a friday.`. If the date don't exist, for exemple 20/14/2022, the program prints `The date 20/14/2022 don't exist.`.

You can compile your code with the makefile : `make questionThreeTwo` or you can also use the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/dayofbirth.c -o objs/dayofbirth.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/dayofbirth_argv.c -o objs/dayofbirth_argv.o

clang objs/dayofbirth.o objs/dayofbirth_argv.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./build/questionThreeTwo
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions
```

**Question 3.**

- Create a file called `dayofbirth_stdin.c` inside the `srcs` folder.
- Create a main function to retrieve the name and the birthdate of a user and display is day of birth. For exemple :

```bash
./questionThreeThree
Thibault C.
23/03/2001
```

Displays `Thibault C. was born on a friday.`. If the date don't exist, for exemple 20/14/2022, the program prints `The date 20/14/2022 don't exist.`. The name of the user won't exceed 100 characters.

You can compile your code with the makefile : `make questionThreeThree` or you can also use the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/dayofbirth.c -o objs/dayofbirth.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/dayofbirth_stdin.c -o objs/dayofbirth_stdin.o

clang objs/dayofbirth.o objs/dayofbirth_stdin.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./build/questionThreeThree
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions
```

#### Exercise 4 : pointers

- Create a file called `pointers.c` inside the `srcs` folder and a file `pointers.h` inside the `includes` folder.

**Question 1.**

- Create a function that swaps two integers passed in parameters. The signature of the function is `void swap(int *a, int *b);`.

You can test your function with the given snow tests. You can run it with the makefile with `make questionFourOne` and then `./build/questionFourOne`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/pointers.c -o objs/pointers.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/exerciseFour_questionOne.c -o objs/exerciseFour_questionOne.o

clang objs/pointers.o objs/exerciseFour_questionOne.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./build/questionFourOne
```

**Question 2.**

- Create a function that adds two integers passed in parameters as pointers. The signature of the function is `int add(int *a, int *b);`.

You can test your function with the given snow tests. You can run it with the makefile with `make questionFourTwo` and then `./build/questionFourTwo`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/pointers.c -o objs/pointers.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/exerciseFour_questionTwo.c -o objs/exerciseFour_questionTwo.o

clang objs/pointers.o objs/exerciseFour_questionTwo.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./build/questionFourTwo
```

#### Exercise 5 : string.h

In this exercise, you will create a small `string.h` library. You can't use any other library and builtins functions. To use the `NULL` value you need these lines :

```c
#undef NULL
#define NULL ((void)*0)
```

- Create a file called `string.h` inside the `includes` folder and a file called `string.c` inside the `srcs` folder.

**Question 1.**

- Create a function to calculate the length of a string aka `char*`. The signature of the function is : `int _strlen(const char *str);`

Make sure that the function returns `-1` in case of a `NULL` value.

You can test your function with the given snow tests. You can run it with the makefile with `make questionFiveOne` and then `./build/questionFiveOne`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/pointers.c -o objs/pointers.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/exerciseFive_questionOne.c -o objs/exerciseFive_questionOne.o

clang objs/pointers.o objs/exerciseFive_questionOne.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./buildFive_questionOne
```

**Question 2.**

- Create a function to compare two strings. It should return `1` if strings are equal and `0` if not. Make sure that the function returns `-1` in case of a `NULL` value.

The function signature is : `int _strcmp(const char *str1, const char *str2);`.

You can test your function with the given snow tests. You can run it with the makefile with `make questionFiveTwo` and then `./build/questionFiveTwo`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/pointers.c -o objs/pointers.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/exerciseFive_questionTwo.c -o objs/exerciseFive_questionTwo.o

clang objs/pointers.o objs/exerciseFive_questionTwo.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o ./buildFive_questionTwo
```
