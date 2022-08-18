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
	arc_UnregisterReference((class_Any*)me->someString);
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
DEFINE_ARRAY(T_array_$b$string$s$$e$);

void test_PrintArr(class_Array_String *arr) {
	for (int i = 0; i < arr->length; i++) {
		printf("%s\n", ((class_String*)Array_public_GetElement((class_Array*)arr, i))->buffer);
	}
}

class_String *test_GetString() {

	class_String *strInstance = (class_String*)malloc(sizeof(class_String));
	String_public_Constructor(strInstance);
	arc_RegisterReference((class_Any*)strInstance);

	String_public_Load(strInstance, "cool string business");

	return (class_String*)strInstance;
}

class_Array_String *test_GetStringArray(class_String *val, int count) {

	class_Array *arrInstance = (class_Array*)malloc(sizeof(class_Array));
	Array_public_Constructor(arrInstance, count);
	arc_RegisterReference((class_Any*)arrInstance);

	for (int i = 0; i < count; i++)
	{
		Array_public_SetElement(arrInstance, i, (class_Any*)val);
	}

	return (class_Array_String*)arrInstance;
}

class_Array_T_array_$b$string$s$$e$ *test_Get2DStringArray(class_String *val, int width, int height) {

	// the upper array
	class_Array *arrInstance = (class_Array*)malloc(sizeof(class_Array));
	Array_public_Constructor(arrInstance, width);
	arc_RegisterReference((class_Any*)arrInstance);

	// generate the lower array only once as they are all the same anyways
	class_Array_String *lower = test_GetStringArray(val, height);

	for (int i = 0; i < width; i++)
	{
		Array_public_SetElement(arrInstance, i, (class_Any*)lower);
	}

	arc_UnregisterReference((class_Any*)lower);

	return (class_Array_T_array_$b$string$s$$e$*)arrInstance;
}