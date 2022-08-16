#include<stdlib.h>
#include<stdio.h>
#include "arc.h"
#include "objects.h"

// ReCT ARC system
// ---------------
// this guy here is responsible for keeping track of how many references to an object exist
// if the number hits 0, the object will be destroyed.

// record a new reference being created
void arc_RegisterReference(class_Any* obj)
{
	if (obj == NULL) return;
    obj->referenceCounter++;
}

// record a reference being destroyed
void arc_UnregisterReference(class_Any* obj)
{
    if (obj == NULL) return;

    obj->referenceCounter--;

    // if the reference pointer is 0 (or negative for some random reason)
    // clear it!
    if (obj->referenceCounter <= 0) {
        obj->vtable->dieFunction((void*)obj); // destroy the objects data
		free(obj);                            // destroy the struct
    }
}

// destroy an object, ignoring ARC
void arc_DestroyObject(class_Any* obj)
{
    if (obj == NULL) return;
	
	obj->vtable->dieFunction((void*)obj); // destroy the objects data
	free(obj);                            // destroy the struct
}

// record a new reference being created
void arc_RegisterReferenceVerbose(class_Any* obj, char* comment)
{
    obj->referenceCounter++;

	// debug message
    printf("\33[36mARC \33[0m- \33[32mRegistered %s reference [%d] - %s\33[0m\n", obj->vtable->className, obj->referenceCounter, comment);
}

// record a reference being destroyed
void arc_UnregisterReferenceVerbose(class_Any* obj, char* comment)
{
    if (obj == NULL) return;

    obj->referenceCounter--;

	// debug message
    printf("\33[36mARC \33[0m- \33[33mUnregistered %s reference [%d] - %s\33[0m\n", obj->vtable->className, obj->referenceCounter, comment);

    // if the reference pointer is 0 (or negative for some random reason)
    // clear it!
    if (obj->referenceCounter == 0) {
		// debug message
        printf("\33[36mARC \33[0m- \33[31mDestroying %s instance - %s\33[0m\n", obj->vtable->className, comment);

        obj->vtable->dieFunction((void*)obj);
		free(obj); 
    }
	else if (obj->referenceCounter < 0) {
		// what??
		printf("\33[36mARC \33[0m- \33[0;35mWhat?? [%d] - %s\33[0m\n", obj->referenceCounter, comment);
	}
}