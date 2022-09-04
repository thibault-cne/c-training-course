#include "../includes/snow.h"
#include "../includes/string.h"

snow_main();

describe(_strcmpTesting)
{
    it("Testing _strcmp with normal values")
    {
        asserteq(_strcmp("Test", "Test"), 0);
        asserteq(_strcmp("Tests", "Hello"), 1);
    }

    it("Testing _strcmp with NULL pointer")
    {
        asserteq(_strcmp("Test", NULL), -1);
        asserteq(_strcmp(NULL, "Test"), -1);
    }

    it("Testing _strcmp with empty value : ''")
    {
        asserteq(_strcmp("Test", ""), 1);
        asserteq(_strcmp("", ""), 0);
    }
}
