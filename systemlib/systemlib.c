#ifdef _WIN32
#include <Windows.h>
#else
#include <unistd.h>
#endif
#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

bool isCursorVisible = true;

void rct_Print(const char* text)
{
    printf("%s\n", text);
}

void rct_Write(const char* text)
{
    printf("%s", text);
}

char* rct_Input()
{
    char *input = (char*)malloc(1024);
    scanf("%1023[^\n]", input);
    return input;
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
    if (state)
	{
        printf("\e[?25l");
        isCursorVisible = true;
	}
	else
	{
		printf("\e[?25h");
        isCursorVisible = false;
	}
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
