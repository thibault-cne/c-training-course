/* ************************************************************************************************************ */
/*                                                                                                              */
/*                                                                                                              */
// main.c
/*                                                                                                              */
// by Thibault Cheneviere : thibault.cheneviere@telecomnancy.eu
/*                                                                                                              */
// Created : 2022/08/22 09/50/46
/*                                                                                                              */
/*                                                                                                              */
/* ************************************************************************************************************ */

#include "../includes/core.h"

int main()
{
    contact_linked_list *contactList;
    char *input;
    int exitStatus;

    exitStatus = 1;
    contactList = create_contact_linked_list(8);
    input = malloc(sizeof(char) * 7);

    printf("Welcome to the phoneBook application.\n");

    while (exitStatus == 1)
    {
        fgets(input, 7, stdin);

        if (_strncmp(input, "EXIT", 4) == 0)
        {
            printf("Closing the application. All contact lost.\n");
            break;
        }

        if (_strncmp(input, "ADD", 3) == 0)
        {
            printf("Adding new Contact\n");
            addNewContact(contactList);
        }

        if (_strncmp(input, "SEARCH", 6) == 0)
        {
            int index;

            printf("Enter an index : ");
            scanf("%d", &index);

            searchContact(contactList, index);
        }
    }

    free_contact_linked_list(contactList);

    return 0;
}

int addNewContact(contact_linked_list *contactList)
{
    char *firstName;
    char *lastName;
    char *phoneNumber;
    char *drakestSecret;
    char *nickName;

    contact *newCct;

    // Init all strings
    firstName = malloc(sizeof(char) * 101);
    lastName = malloc(sizeof(char) * 101);
    phoneNumber = malloc(sizeof(char) * 101);
    drakestSecret = malloc(sizeof(char) * 101);
    nickName = malloc(sizeof(char) * 101);

    getContactField("firstname", &firstName, 101);
    getContactField("lastname", &lastName, 101);
    getContactField("phone number", &phoneNumber, 101);
    getContactField("darkest secret", &drakestSecret, 101);
    getContactField("nickname", &nickName, 101);

    printf("FirstName : %s %s %s %s %s\n", firstName, lastName, phoneNumber, drakestSecret, nickName);

    newCct = create_contact(firstName, lastName, phoneNumber, drakestSecret, nickName);

    add_contact(contactList, newCct);

    return 1;
}

int getContactField(char *fieldName, char **field, int size)
{
    do
    {
        printf("Please enter %s : ", fieldName);
        fgets(*field, size, stdin);
    } while (_strncmp(*field, "\n", 1) == 0);

    _strformat(*field);

    return 1;
}

int searchContact(contact_linked_list *contactList, int index)
{
    if (index < contactList->maxSize && (index < contactList->index || contactList->isFull == 1))
    {
        contact *sContact;

        sContact = get_contact(contactList, index);

        displayField(sContact->firstName);
        printf(" | ");
        displayField(sContact->lastName);
        printf(" | ");
        displayField(sContact->nickName);
        printf(" | ");
        displayField(sContact->darkestSecret);

        printf("\n");

        return 1;
    }

    return 0;
}

int displayField(const char *field)
{
    int i;
    int strLen;

    i = -1;
    strLen = _strlen(field);

    while (++i < 10)
    {
        if (field[i] && i < strLen)
        {
            printf("%c", field[i]);
        }
        else
        {
            printf(" ");
        }
    }

    if (strLen > 10)
    {
        printf(".");
    }
    else
    {
        printf("%c", field[i]);
    }

    return 1;
}
