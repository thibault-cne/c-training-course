#include <stdio.h>
#include <stdlib.h>

#include "../includes/dayofbirth.h"

void displayName(char **name)
{
    int i = -1;

    while (name[0][++i] != '\n')
    {
        continue;
    }

    name[0][i] = '\0';
}

int main()
{
    int day;
    int month;
    int year;

    char *name = malloc(sizeof(char) * 101);
    char *date = malloc(sizeof(char) * 11);

    char *days[7] = {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"};

    fgets(name, 100, stdin);
    fgets(date, 11, stdin);

    sscanf(date, "%d/%d/%d", &day, &month, &year);

    int dayNumber = dayfromdate(day, month, year);

    if (dayNumber == -1)
    {
        printf("The date %s don't exist.\n", date);
        return 0;
    }

    displayName(&name);

    printf("%s was born on a %s.\n", name, days[dayNumber]);
    return 1;
}
