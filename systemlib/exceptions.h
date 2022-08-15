#ifndef EXCEPTIONS_H
#define EXCEPTIONS_H

// define the throwing function
void exc_Throw(char *message);
void exc_ThrowIfNull(void *pointer);

#endif