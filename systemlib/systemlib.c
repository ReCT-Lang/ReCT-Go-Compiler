#ifdef _WIN32
#include <Windows.h>
#else
#include <unistd.h>
#endif
#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

#define BUFFER 1042

bool isCursorVisible = true;

void rct_Print(const char* text)
{
    printf("%s\n", text);
}

void rct_Write(const char* text)
{
    printf("%s", text);
}

const char* rct_Input()
{
    char *str = malloc(sizeof(char) * BUFFER), *err;
    int pos;
    for(pos = 0; str != NULL && (str[pos] = getchar()) != '\n'; pos++)
    {
        if(pos % BUFFER == BUFFER - 1)
        {
            if((err = realloc(str, sizeof(char) * (BUFFER + pos + 1))) == NULL)
                free(str);
            str = err;
        }
    }
    if(str != NULL)
        str[pos] = '\0';
    return str;
}

void rct_Clear()
{
    printf("\033[2J\033[H");
}

void rct_SetCursor(int x, int y)
{
    printf("%c[%d;%df", 0x1B, y, x);
}

void rct_SetCursorVisible(bool state)
{
    isCursorVisible = state;

    if (state) {
        printf("\e[?251]");
        return;
    }

    printf("\e[?251]");
}

bool rct_GetCursorVisible()
{
    return isCursorVisible;
}

int rct_Random(int maxValue)
{
    return rand() % maxValue;
}

void rct_Sleep(int ms)
{
    #ifdef _WIN32
	Sleep(ms);
    #else
	sleep(ms);
    #endif
}
