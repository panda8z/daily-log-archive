/* praise2.c  */
#include <stdio.h>
#include <string.h> // strlen() is here
#define PRAISE "What a super marvelous name!"
int main(void)
{
    char name[40];

    printf("What`s your name?\n");
    scanf("%s", name);
    printf("Hello, %s. %s\n", name, PRAISE);
    printf("Your name of %d letters occupies %d memory cells.\n",strlen(name), sizeof name);
    printf("The phrase of prasie has %d letters",strlen(PRAISE));
    printf("and occupies %d memory cells.\n", sizeof(PRAISE));
}