/* ************************************************************************************************************ */
/*                                                                                                              */
/*                                                                                                              */
// phoneBook.c
/*                                                                                                              */
// by Thibault Cheneviere : thibault.cheneviere@telecomnancy.eu
/*                                                                                                              */
// Created : 2022/08/30 15/26/00
/*                                                                                                              */
/*                                                                                                              */
/* ************************************************************************************************************ */

#include "../includes/phoneBook.h"

contact_linked_list *create_contact_linked_list(int size)
{
    contact_linked_list *list;

    list = malloc(sizeof(contact_linked_list));

    list->maxSize = size;
    list->index = 0;
    list->isFull = 0;
    list->first = NULL;

    return list;
}

int free_contact_linked_list(contact_linked_list *list)
{
    if (list->first != NULL)
    {
        free_contact_linked_list_element(list->first);
    }

    free(list);

    return 1;
}

int modify_contact(contact_linked_list *list, int index, contact *con)
{
    int i;
    contact_linked_list_element *elem;

    i = -1;
    elem = list->first;

    while (++i < index)
    {
        elem = elem->next;
    }

    free_contact(elem->con);
    elem->con = con;

    return 1;
}

int add_contact(contact_linked_list *list, contact *con)
{
    if (list->index >= list->maxSize)
    {
        list->index = 0;
        list->isFull = 1;
    }

    if (list->isFull == 1)
    {
        modify_contact(list, list->index, con);
    }
    else
    {
        contact_linked_list_element *elem;

        elem = list->first;

        if (elem == NULL)
        {
            list->first = create_contact_linked_list_element(con);
        }
        else
        {
            while (elem->next != NULL)
            {
                elem = elem->next;
            }

            elem->next = create_contact_linked_list_element(con);
        }
    }

    list->index++;

    return 1;
}

contact *get_contact(contact_linked_list *list, int index)
{
    int i;
    contact_linked_list_element *elem;

    i = -1;
    elem = list->first;

    while (++i < index)
    {
        elem = elem->next;
    }

    return elem->con;
}

contact_linked_list_element *create_contact_linked_list_element(contact *con)
{
    contact_linked_list_element *elem;

    elem = malloc(sizeof(contact_linked_list_element));

    elem->con = con;
    elem->next = NULL;

    return elem;
}

int free_contact_linked_list_element(contact_linked_list_element *elem)
{
    contact_linked_list_element *nextElem;
    contact_linked_list_element *temp;

    nextElem = elem;

    while (nextElem != NULL)
    {
        temp = nextElem;
        nextElem = nextElem->next;

        free(temp->con);
        free(temp);
    }

    return 1;
}
