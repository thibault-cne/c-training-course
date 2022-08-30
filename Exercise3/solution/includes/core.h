/* ************************************************************************************************************ */
/*                                                                                                              */
/*                                                                                                              */
// core.h
/*                                                                                                              */
// by Thibault Cheneviere : thibault.cheneviere@telecomnancy.eu
/*                                                                                                              */
// Created : 2022/08/30 16/09/47
/*                                                                                                              */
/*                                                                                                              */
/* ************************************************************************************************************ */

#ifndef CORE
#define CORE

#include <stdio.h>

#include "string.h"
#include "phoneBook.h"

int addNewContact(contact_linked_list *contactList);
int getContactField(char *fieldName, char **field, int size);
int searchContact(contact_linked_list *contactList, int index);
int displayField(const char *field);

#endif // CORE
