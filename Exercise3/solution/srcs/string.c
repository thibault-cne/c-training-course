/* ************************************************************************************************************ */
/*                                                                                                              */
/*                                                                                                              */
// string.c
/*                                                                                                              */
// by Thibault Cheneviere : thibault.cheneviere@telecomnancy.eu
/*                                                                                                              */
// Created : 2022/08/28 01/05/12
/*                                                                                                              */
/*                                                                                                              */
/* ************************************************************************************************************ */

#include "../includes/string.h"

// Format the string to remove the \n characters at the end
// It also returns the length of the string.
int _strformat(char *str)
{
    int i;

    i = -1;

    while (str[++i] && str[i] != '\n')
    {
        continue;
    }

    str[i] = '\0';

    return i;
}

int _strlen(const char *str)
{
    if (str == NULL)
    {
        return -1;
    }

    int i;

    i = -1;

    while (str[++i])
    {
        continue;
    }

    return i;
}

int _strncmp(const char *str1, const char *str2, int size)
{
    if (_strlen(str1) < size || _strlen(str2) < size)
    {
        return -1;
    }

    int i;

    i = -1;

    while (++i < size)
    {
        if (str1[i] != str2[i])
        {
            return i + 1;
        }
    }

    return 0;
}
