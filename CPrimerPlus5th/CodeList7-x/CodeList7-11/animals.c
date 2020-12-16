/* animals.c -- 使用switch语句 */
#include <stdio.h>
#include <ctype.h>
int main(void)
{
    char ch;
    printf("Give me a letter of the alphabet, and I will give ");
    printf("an animal name\nbeginning with that letter.\n");
    printf("Please type in a letter: type # to end my act.\n");
    while ((ch = getchar()) != '#')
    {
        if (ch == '\n')
        {
            continue;
        }
        if (islower(ch)) //只识别小写字母
        {
            switch (ch)
            {
            case 'a':
                printf("argali, a wild sheep of Asia\n");
                break;
            case 'b':
                printf("babirusa, a wild pig of Malay\n");
                break;
            case 'c':
                printf("coati, racoonlike animal\n");
                break;
            case 'd':
                printf("desman, aquatic, molelike critter\n");
                break;
            case 'e':
                printf("echidna, the spiny anteater\n");
                break;
            case 'f':
                printf("fisher, brownish marten\n");
                break;
            default:
                printf("That`s a stumper!\n");
                break;
            }
        }
        else
        {
            printf("I recognize only lowercase letters.\n");
        }

        while (getchar() != '\n')
        {
            continue;
        }
        printf("Please type another letter or a #.\n");
    }
    printf("Bye!\n");
    return 0;
}