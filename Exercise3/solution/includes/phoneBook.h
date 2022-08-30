/* ************************************************************************************************************ */
/*                                                                                                              */
/*                                                                                                              */
// phoneBook.h
/*                                                                                                              */
// by Thibault Cheneviere : thibault.cheneviere@telecomnancy.eu
/*                                                                                                              */
// Created : 2022/08/30 15/29/11
/*                                                                                                              */
/*                                                                                                              */
/* ************************************************************************************************************ */

#ifndef PHONEBOOK
#define PHONEBOOK

#include <stdlib.h>

#include "contact.h"

typedef struct _contact_linked_list contact_linked_list;
typedef struct _contact_linked_list_element contact_linked_list_element;

struct _contact_linked_list
{
    int index;
    int maxSize;
    int isFull;
    struct _contact_linked_list_element *first;
};

struct _contact_linked_list_element
{
    contact *con;
    struct _contact_linked_list_element *next;
};

contact_linked_list *create_contact_linked_list(int size);
int free_contact_linked_list(contact_linked_list *list);

int modify_contact(contact_linked_list *list, int index, contact *con);
int add_contact(contact_linked_list *list, contact *con);
contact *get_contact(contact_linked_list *list, int index);

contact_linked_list_element *create_contact_linked_list_element(contact *con);
int free_contact_linked_list_element(contact_linked_list_element *elem);

#endif // PHONEBOOK
