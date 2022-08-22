#ifndef EXCEPTIONS_H
#define EXCEPTIONS_H

// define the throwing function (very athletic)
// ============================================

// standard exception throwing
void exc_Throw(char *message);

// exception shortcuts
void exc_ThrowIfNull(void *pointer);
void exc_ThrowIfInvalidCast(class_Any* from, Any_vTable *to);

#endif