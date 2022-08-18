#include <stdio.h>
#include <stdlib.h>
#include "./../systemlib/objects.h"
#include "./../systemlib/arc.h"

// =============================================================================
// MY VERY COOL CLASS
// =============================================================================
// too lazy to use a header file
typedef struct class_Thing  class_Thing;
typedef struct Thing_vTable Thing_vTable;

void Thing_public_Constructor(class_Thing*, class_String*);
void Thing_public_Die(void*);
void Thing_public_Output(class_Thing*);
void Thing_public_ChangeString(class_Thing*, class_String*);
class_String *Thing_public_GetString(class_Thing*);

// =============================================================================
struct Thing_vTable {
	const Any_vTable *parentVTable;
	const char *className;
	DiePointer dieFunction;
};

struct class_Thing {
	const Thing_vTable *vtable;
	int referenceCounter;
	class_String *someString;
};

const char *Thing_Fields_Const[] = {"someString"};
const Thing_vTable Thing_vTable_Const = {NULL, "Thing", &Thing_public_Die};

void Thing_public_Constructor(class_Thing *this, class_String *val) {
	this->vtable = &Thing_vTable_Const;
	this->referenceCounter = 0;
	this->someString = val;
}

void Thing_public_Die(void *this) {
	class_Thing* me = (class_Thing*)this;
}

void Thing_public_Output(class_Thing *this) {
	printf("%s\n", this->someString->buffer);
}

void Thing_public_ChangeString(class_Thing *this, class_String *val) {
	this->someString = val;
}

class_String *Thing_public_GetString(class_Thing *this) {
	return this->someString;
}

// =============================================================================
// OTHER THINGS
// =============================================================================

void test_PrintArr(class_Thing *arr) {
	//printf("Got array of length %d", Array_public_GetLength(arr));
}

class_String *test_GetString() {
	class_String *strInstance = (class_String*)malloc(sizeof(class_String));
	String_public_Constructor(strInstance);
	arc_RegisterReference((class_Any*)strInstance);

	String_public_Load(strInstance, "cool string");

	return strInstance;
}