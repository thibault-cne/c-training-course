#include "../solution/includes/snow.h"
#include "../includes/string.h"

snow_main();

describe(_strlenTesting)
{
    it("Testing _strlen with normal values")
    {
        asserteq(_strlen("Hello"), 5);
        asserteq(_strlen("Test"), 4);
    }

    it("Testing _strlen with NULL pointer")
    {
        asserteq(_strlen(NULL), -1);
    }

    it("Testing _strlen with empty value : ''")
    {
        asserteq(_strlen(""), 0);
    }
}
