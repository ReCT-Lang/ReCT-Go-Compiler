#ifdef _WIN32
#include <Windows.h>
#else
#include <unistd.h>
#endif
#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

bool isCursorVisible = true;

void _Print(const char* text)
{
    printf("%s", text);
    putch('\n');
}

void _Write(const char* text)
{
    printf("%s", text);
}

char* _Input()
{
    char input[1023];
    scanf("%1023[^\n]", input, 1023);
    return input;
}

void _Clear()
{
    printf("\033[2J\033[H");
}

void _SetCursor(int x, int y)
{
    printf("%c[%d;%df", 0x1B, y, x);
}

void _SetCursorVisible(bool state)
{
    switch (state)
    {
        case true:
            printf("\e[?25l");
            isCursorVisible = true;
            break;
        case false:
            printf("\e[?25h");
            isCursorVisible = false;
            break;
    }
}

bool _GetCursorVisible()
{
    return isCursorVisible;
}

int _Random(int maxValue)
{
    return rand() % maxValue;
}

void _Sleep(int ms)
{
    #ifdef _WIN32
	Sleep(ms);
    #else
	sleep(ms);
    #endif
}
