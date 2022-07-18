#include "../includes/maths.h"

int countAllIntegerOccurencies(int max, int integer)
{
    int i = 0;
    int cpt = 0;

    while (i <= max)
    {
        int div = i / 10;
        int rest = i % 10;

        if (rest == integer)
        {
            cpt++;
        }

        while (div != 0)
        {
            rest = div % 10;
            div = div / 10;

            if (rest == integer)
            {
                cpt++;
            }
        }

        i++;
    }

    return cpt;
}
