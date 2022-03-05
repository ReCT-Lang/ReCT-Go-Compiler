#include<stdlib.h>
#include<string.h>
#include<stdbool.h>
#include<stdio.h>
#include<pthread.h> // For thread object UwU
#include "objects.h"
#include "arc.h"

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

	// free the old buffer if theres anything in there
	if (this->buffer != NULL)
		free(this->buffer);

	// change our pointer
	this->buffer = output;

	// update our length and max length
	this->length = size;
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

// string utils
class_String* String_public_Concat(class_String* a, class_String* b) {
	// new buffer for concatinated string
	char *newBuffer = (char*)malloc(a->length + b->length + 1);
	strcpy(newBuffer, a->buffer);
	strcat(newBuffer, b->buffer);

	// create a new string object
	class_String *newStr = (class_String*)malloc(sizeof(class_String));
	String_public_Constructor(newStr);
	String_public_Load(newStr, newBuffer);
	arc_RegisterReference((class_Any*)newStr);

	free(newBuffer);
	return newStr;
}

bool String_public_Equal(class_String* a, class_String* b) {
	// use strcmp to check if they are equal
	int result = strcmp(a->buffer, b->buffer);

	return result == 0;
}

// i have no idea how to access the struct
char *String_public_GetBuffer(class_String* this) {
	return this->buffer;
}

int String_public_GetLength(class_String* this) {
	return this->length;
}

class_String *String_public_Substring(class_String* this, int start, int length) {
	
	// new string buffer
	char *subBuffer;

	// make sure the substring is valid
	if (start < 0 || length < 0 || start + length > this->length)
		subBuffer = "https://www.youtube.com/watch?v=dQw4w9WgXcQ";
	else
	{
		subBuffer = (char*)malloc(length + 1);
		memcpy(subBuffer, &this->buffer[start], length);
		subBuffer[length] = '\0';
	}
	
	// create a string object
	class_String *newString = (class_String*)malloc(sizeof(class_String));
	String_public_Constructor(newString);
	String_public_Load(newString, subBuffer);
	arc_RegisterReference((class_Any*)newString);

	// clear the work buffer
	free(subBuffer);

	// return the string object
	return newString;
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
	// if the object is null -> return the default value
	if (this == NULL) return 0;

	// if not -> return the stored value
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
	// if the object is null -> return the default value
	if (this == NULL) return 0.0;

	// if not -> return the stored value
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
	// if the object is null -> return the default value
	if (this == NULL) return false;

	// if not -> return the stored value
	return this->value;
}

// -----------------------------------------------------------------------------
// "array" object type
// Note: this is an object represents an array, it only holds object types atm
// The array wont make data copies, it will just hold references
// -----------------------------------------------------------------------------

// definition for the Bool vTable
const Array_vTable Array_vTable_Const = {&Any_vTable_Const, "Array", &Array_public_Die};

// definition for the objects constructor
void Array_public_Constructor(class_Array* this, int length) {
	this->vtable = &Array_vTable_Const;
	this->referenceCounter = 0;
	this->length = length;
	this->maxLen = length;
	this->factor = 5;

	// allocate space needed for our pointers
	this->elements = (class_Any**)calloc(length, sizeof(class_Any*));
}

// definition for the objects destructor
void Array_public_Die(void* this) {
	// bitcast the void* into an Array pointer
	class_Array *me = (class_Array*)this;

	// go through each object and dereference it
	for (int i = 0; i < me->length; i++) {
		arc_UnregisterReference(me->elements[i]);
	}

	// free the allocated space
	free(me->elements);
}

// definition for a element access
class_Any* Array_public_GetElement(class_Array* this, int index) {
	if (index < 0 || index >= this->length)
		return NULL;

	return this->elements[index];
}

// definition for a element assignment
void Array_public_SetElement(class_Array* this, int index, class_Any *element) {
	if (index < 0 || index >= this->length)
		return;

	// increase arc reference count
	arc_RegisterReference(element);

	// decrease arc reference count for the previous element
	arc_UnregisterReference((class_Any*)this->elements[index]);

	*(this->elements + index) = element;
}

// definition for array length
int Array_public_GetLength(class_Array* this) {
	return this->length;
}

// definition for a push function
void Array_public_Push(class_Array* this, class_Any *element) {

	// if the array buffer needs to grow
	if (this->length == this->maxLen) {
		int newLength = this->length + this->factor;

		// try to use realloc() first bc its faster
		class_Any **newBuffer = realloc(this->elements, sizeof(class_Any*) * newLength);

		// if that failed, do it the long way
		if (newBuffer == NULL) {
			// allocate a new buffer
			newBuffer = (class_Any**)malloc(sizeof(class_Any*) * newLength);

			// copy over the old one
			memcpy(newBuffer, this->elements, sizeof(class_Any*) * this->length);

			// free the old buffer
			free(this->elements);

		}
		
		// change our pointer
		this->elements = newBuffer;

		// NULL the new array slots
		for (int i = 0; i < this->factor; i++)
			*(this->elements + this->length + i) = NULL;

		// update our max length 
		this->maxLen = newLength;

	}

	// update our length
	this->length++;

	// assign the given element to the new slot
	Array_public_SetElement(this, this->length - 1, element);
}

// -----------------------------------------------------------------------------
// "parray" object type
// Note: this is a primitive version of "array"
// -----------------------------------------------------------------------------

// definition for the Bool vTable
const pArray_vTable pArray_vTable_Const = {&Any_vTable_Const, "pArray", &pArray_public_Die};

// definition for the objects constructor
void pArray_public_Constructor(class_pArray* this, int length, int elemSize) {
	this->vtable   = &pArray_vTable_Const;
	this->referenceCounter = 0;
	this->length   = length;
	this->maxLen   = length;
	this->factor   = 4;
	this->elemSize = elemSize;

	this->elements = calloc(length, elemSize);
}

// definition for the objects destructor
void pArray_public_Die(void* this) {
	// bitcast the void* into an Array pointer
	class_pArray *me = (class_pArray*)this;

	// free the allocated space
	free(me->elements);
}

// definition for array length
int pArray_public_GetLength(class_pArray* this) {
	return this->length;
}

// definition for an array.Grow() method
void *pArray_public_Grow(class_pArray* this) {

	// check if growing is actually needed
	if (this->length == this->maxLen)
	{
		int newLength = (this->length + this->factor) * this->elemSize;

		// try to use realloc() first bc its faster
		void *output = realloc(this->elements, newLength);

		// if that failed, do it the long way
		if (output == NULL) {
			// allocate a new buffer
			output = malloc(newLength);

			// copy over the old one
			memcpy(output, this->elements, this->length * this->elemSize);

			// free the old buffer
			free(this->elements);

		}

		// change our pointer
		this->elements = output;

		// update our new max length!
		this->maxLen = this->length + this->factor;
	}

	// increase the length variable as grow being called means something will be pushed
	this->length++;

	// return a pointer where the new element is supposed to go
	return (void*)(this->elements + (this->length-1) * this->elemSize);
}

void *pArray_public_GetElementPtr(class_pArray* this, int index) {
	if (index < 0 || index >= this->length)
		return NULL;

	return (void*)(this->elements + index * this->elemSize);
}

// -----------------------------------------------------------------------------
// "thread" object type
// Note: this uses the pthread library. Make sure executable is compiled with flag -lpthread
// Information: Want to know more about function pointers and pthreads?:
//      (ALL GUIDES ARE IN ENGLISH)
//      Function pointers:  https://www.cprogramming.com/tutorial/function-pointers.html
//      Multithreading:     https://www.geeksforgeeks.org/thread-functions-in-c-c/
//      Multithreading 2:   https://www.geeksforgeeks.org/multithreading-c-2/
// -----------------------------------------------------------------------------

// definition for the Thread vTable
const Thread_vTable Thread_vTable_Const = {&Any_vTable_Const, "Thread", &Thread_public_Die};

// definition for the objects constructor
void Thread_public_Constructor(class_Thread *this, void *(*__routine) (void*), void *args) {
	this->vtable = &Thread_vTable_Const;
	this->referenceCounter = 0;
	this->__routine = __routine;
	this->args = args;
}

// definition for the objects destructor
void Thread_public_Die(void* this) {}

// start thread
void Thread_public_Start(class_Thread *this) {

    // Args: thread id, attributes, function, arguments
    // if attributes are NULL, they are set to default.
    pthread_create(&this->id, NULL, this->__routine, this->args);
}

// join thread
void Thread_public_Join(class_Thread *this) {
	pthread_join(this->id, NULL);
}

// end thread
void Thread_public_Kill(class_Thread *this) {
    pthread_exit(NULL);
}