#include <stdio.h>
#include <stdlib.h>

#include "../includes/dayofbirth.h"

int main(int argc, char **argv)
{
    if (argc != 4)
    {
        return 0;
    }

    int day;
    int month;
    int year;

    char *firstName;
    char *lastName;

    char *days[7] = {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"};

    firstName = argv[1];
    lastName = argv[2];

    sscanf(argv[3], "%d/%d/%d", &day, &month, &year);

    int dayNumber = dayfromdate(day, month, year);

    if (dayNumber == -1)
    {
        printf("The date %s don't exist.\n", argv[3]);
        return 0;
    }

    printf("%s %s was born on a %s.\n", firstName, lastName, days[dayNumber]);
    return 0;
}
