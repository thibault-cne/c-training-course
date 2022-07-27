# Exercise 2

## Table of contents

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

clang -Wall -Wextra -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls  ./objs/helloworld.o -o ./build/questionOne.exe
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions q11
```

#### Exercise 2 : arguments/stdout

In this exercise you will create a function that returns the max of 3 integers. You will retrieve the 3 integers from the arguments of the programs and you will send the result on the stdout (terminal prompt).

**Question 1.**

- Create a `max3.c` file inside the `srcs` folder
- Create a `max3.h` file inside the `includes` folder and link this file to `max3.c`
- Inside this file create a function with the signature `int max(int a, int b, int c);` thats return the maximum between 3 integers.

You can test your function with the given snow tests. You can run it with the makefile with `make questionTwo` and then `./build/questionTwoOne.exe`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3.c -o objs/max3.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/arguments_stdout_questionOne.c -o objs/arguments_stdout_questionOne.o

clang objs/max3.o objs/arguments_stdout_questionOne.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o questionTwo -I includes/
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
