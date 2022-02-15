#include "objects.h"

// ReCT ARC system
// ---------------
// this guy here is responsible for keeping track of how many references to an object exist
// if the number hits 0, the object will be destroyed.

// record a new reference being created
void arc_RegisterReference(class_Any* obj)
{
    obj->referenceCounter++;
}

// record a reference being destroyed
void arc_UnregisterReference(class_Any* obj)
{
    obj->referenceCounter--;

    // if the reference pointer is 0 (or negative for some random reason)
    // clear it!
    if (obj->referenceCounter <= 0) {
        obj->vtable->dieFunction(obj);
    }
}