#include "objects.h"

#ifndef ARC_H
#define ARC_H

#ifdef __cplusplus
extern "C" {
#endif

void arc_RegisterReference(class_Any*);
void arc_UnregisterReference(class_Any*);
void arc_RegisterReferenceVerbose(class_Any*, char*);
void arc_UnregisterReferenceVerbose(class_Any*, char*);

void arc_DestroyObject(class_Any*);

#ifdef __cplusplus
}
#endif

#endif
