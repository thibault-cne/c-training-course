#include "../includes/string.h"

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

int _strcmp(const char *str1, const char *str2)
{
    if (str1 == NULL || str2 == NULL)
    {
        return -1;
    }

    if (_strlen(str1) != _strlen(str2))
    {
        return 0;
    }

    int i;

    i = -1;

    while (str1[++i])
    {
        if (str1[i] != str2[i])
        {
            return 0;
        }
    }

    return 1;
}
