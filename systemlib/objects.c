#include<stdlib.h>
#include<string.h>
#include<stdbool.h>
#include<stdio.h>
#include "objects.h"

// NOTE: I made all class names capitalised, this is to distinguish primitives
//       like int, float, etc from they "boxed" (objectified) versions

// -----------------------------------------------------------------------------
// base "any" object type
// Note: all object types will inherit from this
// -----------------------------------------------------------------------------

// definition for the Any vTable
const Any_vTable Any_vTable_Const = {NULL, "Any", &Any_public_Die};

// definition for the objects constructor
void Any_public_Constructor(class_Any* this) {
	this->vtable = &Any_vTable_Const;
	this->referenceCounter = 0;
}

// definition for the objects destructor
void Any_public_Die(void* this) {}

// -----------------------------------------------------------------------------
// "string" object type
// Note: this is our wrapper for strings!
// -----------------------------------------------------------------------------

// definition for the String vTable
const String_vTable String_vTable_Const = {&Any_vTable_Const, "String", &String_public_Die};

// definition for the objects constructor
void String_public_Constructor(class_String* this) {
	this->vtable = &String_vTable_Const;
	this->referenceCounter = 0;
	this->buffer = NULL;
	this->length = 0;
	this->maxLen = 0;
	this->factor = 16;
}

// definition for the objects destructor
void String_public_Die(void* this) {
    // convert generic pointer to string class pointer
    class_String* me = (class_String*)this;

	if (me->buffer != NULL) {
		free(me->buffer);
	}
}

// definition for a string.Load method
// -----------------------------------------------------------------------------
// [i] this is for loading char arrays into a string object
// -----------------------------------------------------------------------------
void String_public_Load(class_String* this, char* source) {

	// get the length of out source
	int size = strlen(source);

	// allocate a new buffer
	char* output = malloc(size + 1);

	// copy over the source
	memcpy(output, source, size + 1);

	// free the old buffer
	free(this->buffer);

	// change our pointer
	this->buffer = output;

	// update our max length
	this->maxLen = size;
}

// definition for a string.Resize() method
void String_public_Resize(class_String* this, int size) {
	// allocate a new buffer
	char* output = malloc(size);

	// copy over the old one
	memcpy(output, this->buffer, this->length);

	// free the old buffer
	free(this->buffer);

	// change our pointer
	this->buffer = output;

	// update our new max length!
	this->maxLen = size;
}

// definition for a string.Resize() method
void String_public_AddChar(class_String* this, char value) {
	// check if we need to grow the string
	if (this->length == this->maxLen) {
		// grow the string buffer by our growing factor
		String_public_Resize(this, this->maxLen + this->factor);
	}

	// put the chat at the end of the buffer
	this->buffer[this->length] = value;

	// increase our length
	this->length++;
}

// -----------------------------------------------------------------------------
// "int" object type
// Note: this is an object version of an int, this is to box and crunch it
// -----------------------------------------------------------------------------

// definition for the Int vTable
const Int_vTable Int_vTable_Const = {&Any_vTable_Const, "Int", &Int_public_Die};

// definition for the objects constructor
void Int_public_Constructor(class_Int* this, int value) {
	this->vtable = &Int_vTable_Const;
	this->referenceCounter = 0;
	this->value = value;
}

// definition for the objects destructor
void Int_public_Die(void* this) {}

// definition for a string.Resize() method
int Int_public_GetValue(class_Int* this) {
	return this->value;
}

// -----------------------------------------------------------------------------
// "float" object type
// Note: this is an object version of a float, this is to box and crunch it
// -----------------------------------------------------------------------------

// definition for the Float vTable
const Float_vTable Float_vTable_Const = {&Any_vTable_Const, "Float", &Float_public_Die};

// definition for the objects constructor
void Float_public_Constructor(class_Float* this, float value) {
	this->vtable = &Float_vTable_Const;
	this->referenceCounter = 0;
	this->value = value;
}

// definition for the objects destructor
void Float_public_Die(void* this) {}

// definition for a string.Resize() method
float Float_public_GetValue(class_Float* this) {
	return this->value;
}

// -----------------------------------------------------------------------------
// "bool" object type
// Note: this is an object version of a bool, this is to box and crunch it
// -----------------------------------------------------------------------------

// definition for the Bool vTable
const Bool_vTable Bool_vTable_Const = {&Any_vTable_Const, "Bool", &Bool_public_Die};

// definition for the objects constructor
void Bool_public_Constructor(class_Bool* this, bool value) {
	this->vtable = &Bool_vTable_Const;
	this->referenceCounter = 0;
	this->value = value;
}

// definition for the objects destructor
void Bool_public_Die(void* this) {}

// definition for a string.Resize() method
bool Bool_public_GetValue(class_Bool* this) {
	return this->value;
}