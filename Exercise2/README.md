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

To compile your file you can use the given makefile with the command `make questionOne` or use the following commands :

``` bash
mkdir -p ./objs | clang -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/helloworld.c -o objs/helloworld.o

clang -Wall -Wextra -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls  ./objs/helloworld.o -o ./build/questionOne.exe
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions q1
```

#### Exercise 2 : arguments/stdout

In this exercise you will create a function that returns the max of 3 integers. You will retrieve the 3 integers from the arguments of the programs and you will send the result on the stdout (terminal prompt).

**Question 1.**

- Create a `max3.c` file inside the `srcs` folder
- Create a `max3.h` file inside the `includes` folder and link this file to `max3.c`
- Inside this file create a function with the signature `int max(int a, int b, int c);` thats return the maximum between 3 integers.

You can test your function with the given snow tests. You can run it with the makefile with `make questionTwo` and then `./build/questionTwo.exe`.
You can also compile the files with the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3.c -o objs/max3.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/arguments_stdout_questionOne.c -o objs/arguments_stdout_questionOne.o

clang objs/max3.o objs/arguments_stdout_questionOne.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o questionTwo -I includes/
```

**Question 2.**

- Create a `max3_main.c` file inside the `srcs` folder
- Create a main function to retrieve the 3 integers from the call arguments and display the result on the stdout.

You can compile your code with the makefile : `make questionTwo` or you can also use the following code :

``` bash
mkdir -p ./objs | clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3.c -o objs/max3.o
clang -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -c srcs/max3_main.c -o objs/max3_main.o

clang objs/max3.o objs/max3_main.o -DSNOW_ENABLED -Wall -Wextra -Werror -pedantic -O0 -g3 -fsanitize=address -fno-omit-frame-pointer -fno-optimize-sibling-calls -o questionThree -I includes/
```

You can now test your function with the golang package :

``` bash
cd tests
go run main.go -questions q2
```
