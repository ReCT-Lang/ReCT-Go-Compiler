#include<stdlib.h>
#include<stdio.h>
#include<execinfo.h>
#include<string.h>
#include "exceptions.h"

// Very advanced ReCT exceptions
// -----------------------------

// some ANSI colors for the printout
#define BLK "\e[0;30m"
#define RED "\e[0;31m"
#define GRN "\e[0;32m"
#define YEL "\e[0;33m"
#define BLU "\e[0;34m"
#define MAG "\e[0;35m"
#define CYN "\e[0;36m"
#define WHT "\e[0;37m"

#define BBLK "\e[1;30m"
#define BRED "\e[1;31m"
#define BGRN "\e[1;32m"
#define BYEL "\e[1;33m"
#define BBLU "\e[1;34m"
#define BMAG "\e[1;35m"
#define BCYN "\e[1;36m"
#define BWHT "\e[1;37m"

#define RESET "\e[0m"

// the actual throw message
void exc_Throw(char *message) {
	// exception format:
	// [RUNTIME] Encountered Exception! '<exception>'
	// [Stacktrace]
	// ...

	// error head
	printf("%s[RUNTIME] %sEncountered Exception! %s'%s'\n", BRED, RED, BRED, message);

	// stacktrace
	printf("%s[STACKTRACE] %s\n", BYEL, YEL);

	// get the call stack
	void* callstack[128];
	int frames = backtrace(callstack, 128);
	char** strs = backtrace_symbols(callstack, frames);

	// print out the call stack
	// stop printing as soon as we get to non-program things
	for (int i = 1; i < frames; ++i) {
		// check if this string is from an external lib
		char *foundSO  = strstr(strs[i], ".so");
		char *foundDLL = strstr(strs[i], ".dll");

		// if so, destroy the loop
		if (foundSO)  break;
		if (foundDLL) break;

		printf("%s\n", strs[i]);
	}

	// destroy the strings
	free(strs);

	// die();
	exit(-1);
}