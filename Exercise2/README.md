# Exercise 2

## Table of contents

### Requirements

For this exercise, `golang` will be needed to run the tests. You can download it directly from the official golang website : [official golang page](https://go.dev/doc/install).

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
go run main.go -question=q1
```
