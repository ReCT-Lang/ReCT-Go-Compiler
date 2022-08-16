#include "objects.h"

#ifndef ARC_H
#define ARC_H

void arc_RegisterReference(class_Any*);
void arc_UnregisterReference(class_Any*);
void arc_RegisterReferenceVerbose(class_Any*, char*);
void arc_UnregisterReferenceVerbose(class_Any*, char*);

void arc_DestroyObject(class_Any*);

#endif
