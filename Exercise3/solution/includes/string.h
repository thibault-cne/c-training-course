/* ************************************************************************************************************ */
/*                                                                                                              */
/*                                                                                                              */
// string.h
/*                                                                                                              */
// by Thibault Cheneviere : thibault.cheneviere@telecomnancy.eu
/*                                                                                                              */
// Created : 2022/08/28 01/05/37
/*                                                                                                              */
/*                                                                                                              */
/* ************************************************************************************************************ */

#ifndef STRING
#define STRING

#include <stdlib.h>

int _strformat(char *str);
int _strlen(const char *str);

// Compare functions
int _strncmp(const char *str1, const char *str2, int size);

#endif // STRING
