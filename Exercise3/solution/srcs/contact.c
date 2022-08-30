/* ************************************************************************************************************ */
/*                                                                                                              */
/*                                                                                                              */
// contact.c
/*                                                                                                              */
// by Thibault Cheneviere : thibault.cheneviere@telecomnancy.eu
/*                                                                                                              */
// Created : 2022/08/30 15/46/34
/*                                                                                                              */
/*                                                                                                              */
/* ************************************************************************************************************ */

#include "../includes/contact.h"

contact *create_contact(const char *firstName, const char *lastName, const char *phoneNumber, const char *darkestSecret, const char *nickName)
{
    contact *con;

    con = malloc(sizeof(contact));

    con->firstName = firstName;
    con->lastName = lastName;
    con->phoneNumber = phoneNumber;
    con->darkestSecret = darkestSecret;
    con->nickName = nickName;

    return con;
}

int free_contact(contact *con)
{
    free(con);

    return 1;
}
