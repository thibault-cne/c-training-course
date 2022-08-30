/* ************************************************************************************************************ */
/*                                                                                                              */
/*                                                                                                              */
// contact.h
/*                                                                                                              */
// by Thibault Cheneviere : thibault.cheneviere@telecomnancy.eu
/*                                                                                                              */
// Created : 2022/08/30 15/49/08
/*                                                                                                              */
/*                                                                                                              */
/* ************************************************************************************************************ */

#ifndef CONTACT
#define CONTACT

#include <stdlib.h>
#include <string.h>

typedef struct _contact contact;

struct _contact
{
    const char *firstName;
    const char *lastName;
    const char *phoneNumber;
    const char *darkestSecret;
    const char *nickName;
};

contact *create_contact(const char *firstName, const char *lastName, const char *phoneNumber, const char *darkestSecret, const char *nickName);
int free_contact(contact *con);

#endif // CONTACT
