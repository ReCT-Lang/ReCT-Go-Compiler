#ifdef _WIN32
#include <Windows.h>
#else
#include <unistd.h>
#endif
#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <time.h>
#include <gc.h>

#include "../systemlib/objects.h"
#include "../systemlib/exceptions.h"

#define BUFFER 1042

bool isCursorVisible = true;

void sys_Print(class_String *text)
{
	// if theres no string, do a little trolling
	if (text == NULL)
		printf("\n");
	else
    	printf("%s\n", text->buffer);
}

void sys_Write(class_String *text)
{
	if (text != NULL)
    	printf("%s", text->buffer);
}

class_String* sys_Input()
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

	class_String *strInstance = (class_String*)GC_MALLOC(sizeof(class_String));
	strInstance->vtable = (Standard_vTable){NULL, "String"};
    strInstance->vtable.fingerprint = "TO_string[]_";

	String_public_Constructor(strInstance);
	String_public_Load(strInstance, str);

	if(str != NULL)
		free(str);

    return strInstance;
}

void sys_Clear()
{
    printf("\033[2J\033[H");
}

void sys_SetCursor(int x, int y)
{
    printf("%c[%d;%df", 0x1B, y, x);
}

void sys_SetCursorVisible(bool state)
{
    isCursorVisible = state;

    if (state) {
        printf("\e[?251]");
        return;
    }

    printf("\e[?251]");
}

bool sys_GetCursorVisible()
{
    return isCursorVisible;
}

int sys_Random(int maxValue)
{
    return rand() % maxValue;
}

void sys_Sleep(int ms)
{
    #ifdef _WIN32
	Sleep(ms);
    #else
	usleep(ms * 1000);
    #endif
}

int sys_Sqrt(int num)
{
    return (int)floor(sqrt((double)num));
}

int sys_Now()
{
    return (int)clock();
}

class_String *sys_Char(int index)
{
	char *singleChar = (char*)malloc(1);
	singleChar[0] = (char)index;

	class_String *strInstance = (class_String*)GC_MALLOC(sizeof(class_String));
    strInstance->vtable = (Standard_vTable){NULL, "String"};
    strInstance->vtable.fingerprint = "TO_string[]_";

	String_public_Constructor(strInstance);
	String_public_Load(strInstance, singleChar);

	free(singleChar);

    return strInstance;
}