#include "../includes/dayofbirth.h"

int checkDate(int day, int month, int year)
{
    if (year <= 0 || day <= 0 || month <= 0 || month > 12)
    {
        return 0;
    }

    if (month < 8)
    {
        if (month % 2 == 0)
        {
            if (month == 2 && day > 28)
            {
                return 0;
            }
            else
            {
                if (day > 30)
                {
                    return 0;
                }
            }
        }
        else
        {
            if (day > 31)
            {
                return 0;
            }
        }
    }
    else
    {
        if (month % 2 == 0 && day > 31)
        {
            return 0;
        }

        if (month % 2 == 1 && day > 30)
        {
            return 0;
        }
    }

    return 1;
}

int dayfromdate(int day, int month, int year)
{
    if (!checkDate(day, month, year))
    {
        return -1;
    }

    int t;
    int b;
    int c;
    int d;
    int w;

    if (month > 2)
    {
        t = month - 2;
        b = year;
    }
    else
    {
        t = month + 10;
        b = year - 1;
    }

    c = floor(b / 100);
    d = b - 100 * c;

    w = day + floor(2.6 * t - 0.2) + d + floor(d / 4) + floor(c / 4) + 5 * c;

    return w % 7;
}
