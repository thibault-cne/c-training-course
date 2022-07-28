# Exercise 2

## Table of contents

1. [Requirements](#requirements)
2. [Exercises](#exercises)
    - [Exercise 1](#exercise-1--hello-world)
    - [Exercise 2](#exercise-2--argumentsstdoutstdin)
    - [Exercise 3](#exercise-3--argumentsstdinstdout)

### Requirements

For this exercise, `golang` will be needed to run the tests. You can download it directly from the official golang website : [official golang page](https://go.dev/doc/install).

You can run all the golang tests inside the `tests` folder with the following command : 

``` bash
go run main.go -questions ${question list}
```

Where the question list is a list of all questions you want to test. Questions names should be written in this form : `q1` or `question1` for question one.

### Exercises

#### Exercise 1 : Hello World

- Create a file called `helloworld.c` inside the `srcs` folder.
- Inside this file, create a program that display on the terminal the message : Hello World !!

Help : you will need to create a `main()` function and use the `printf` function. Use the man of the function with the command `man 3 printf`.

To compile your file you can use the given makefile with the command `make questionOneOne` or use the following commands :

``` bash
mkdir -p ./objs | clang -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/helloworld.c -o objs/helloworld.o

clang -Wall -Wextra -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls  ./objs/helloworld.o -o ./build/questionOneOne.exe
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions q11
```

#### Exercise 2 : arguments/stdout/stdin

In this exercise you will create a function that returns the max of 3 integers. You will retrieve the 3 integers from the arguments of the programs and you will send the result on the stdout (terminal prompt).

**Question 1.**

- Create a `max3.c` file inside the `srcs` folder
- Create a `max3.h` file inside the `includes` folder and link this file to `max3.c`
- Inside this file create a function with the signature `int max(int a, int b, int c);` thats return the maximum between 3 integers.

You can test your function with the given snow tests. You can run it with the makefile with `make questionTwoOne` and then `./build/questionTwoOne.exe`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3.c -o objs/max3.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/exerciseTwo_questionOne.c -o objs/exerciseTwo_questionOne.o

clang objs/max3.o objs/exerciseTwo_questionOne.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o questionTwoOne -I includes/
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

clang objs/max3.o objs/max3_argv.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o questionTwoTwo -I includes/
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions q22
```

**Question 3.**

- Create a `max3_stdin.c` file inside the `srcs` folder
- Create a main function to calculate the maximum between 3 integers. The user will enter all numbers in the CLI (stdin). Each lines correspond to a number.

Help : you can use the `fgets` function to retrieve the use input. You can get all the information on the function with the command `man 3 fgets`.

You can compile your code with the makefile : `make questionTwoThree` or you can also use the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3.c -o objs/max3.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3_stdin.c -o objs/max3_stdin.o

clang objs/max3.o objs/max3_stdin.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o questionTwoThree -I includes/
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions q23
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

You can test your function with the given snow tests. You can run it with the makefile with `make questionThreeOne` and then `./build/questionThreeOne.exe`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/dayofbirth.c -o objs/dayofbirth.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/exerciseThree_questionOne.c -o objs/exerciseThree_questionOne.o

clang objs/dayofbirth.o objs/exerciseThree_questionOne.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o questionThreeOne -I includes/
```

**Question 2.**

- Create a file called `dayofbirth_argv.c` inside the `srcs` folder.
- Create a main function to retrieve the name and the birthdate of a user and display is day of birth. For exemple, `./questionThreeTwo Thibault C. 23/03/2001` prints `Thibault C. was born on a friday.`. If the date don't exist, for exemple 20/14/2022, the program prints `The date 20/14/2022 don't exist.`.

You can compile your code with the makefile : `make questionThreeTwo` or you can also use the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/dayofbirth.c -o objs/dayofbirth.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/dayofbirth_argv.c -o objs/dayofbirth_argv.o

clang objs/dayofbirth.o objs/dayofbirth_argv.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o questionThreeTwo -I includes/
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions q32
```
