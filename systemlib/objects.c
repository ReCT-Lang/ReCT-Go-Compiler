#include<stdlib.h>
#include<string.h>
#include<stdbool.h>

// NOTE: I made all class names capitalised, this is to distinguish primitives
//       like int, float, etc from they "boxed" (objectified) versions

// -----------------------------------------------------------------------------
// base "any" object type
// Note: all object types will inherit from this
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
typedef struct any_vtable {
	const struct any_vtable* parentVTable; // will be NULL for "any" as its the root
	const char* className;                 // will be "Any"
	void (*Function)(class_Any*)           // destructor function pointer
} Any_vTable;

// the objects struct
typedef struct {
	const Any_vTable* vtable; // "any" is pretty boring, it doesnt store any data
} class_Any;

// definition for the Any vTable
const Any_vTable Any_vTable_Const = {NULL, "Any", &Any_public_Die};

// defintion for the objects constructor
void Any_public_Contructor(class_Any* this) {
	this->vtable = &Any_vTable_Const;
}

// defintion for the objects destructor
void Any_public_Die(class_Any* this) {}

// and thats it!


// -----------------------------------------------------------------------------
// "string" object type
// Note: this is our wrapper for strings!
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
typedef struct {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "String"
	void (*Function)(class_String*) // destructor function pointer
} String_vTable;

// the objects struct
typedef struct {
	const String_vTable* vtable;  // our vTable
	char* buffer;           // for info on this string implementation check out:
	int length;             // https://mapping-high-level-constructs-to-llvm-ir.readthedocs.io/en/latest/appendix-a-how-to-implement-a-string-type-in-llvm/index.html
	int maxLen;
	int factor;
} class_String;

// definition for the String vTable
const String_vTable String_vTable_Const = {&Any_vTable_Const, "String", &String_public_Die};

// defintion for the objects constructor
void String_public_Contructor(class_String* this) {
	this->vtable = &String_vTable_Const;
	this->buffer = NULL;
	this->length = 0;
	this->maxLen = 0;
	this->factor = 16;
}

// defintion for the objects destructor
void String_public_Die(class_String* this) {
	if (this->buffer != NULL) {
		free(this->buffer);
	}
}

// defintion for a string.Resize() method
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

// defintion for a string.Resize() method
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

// the object's vtable (for method lookup and method overriding)
typedef struct {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Int"
	void (*Function)(class_Int*)    // destructor function pointer
} Int_vTable;

// the objects struct
typedef struct {
	const Int_vTable* vtable;  // our vTable
	int value;
} class_Int;

// definition for the Int vTable
const Int_vTable Int_vTable_Const = {&Any_vTable_Const, "Int", &Int_public_Die};

// defintion for the objects constructor
void Int_public_Contructor(class_Int* this, int value) {
	this->vtable = &Int_vTable_Const;
	this->value = value;
}

// defintion for the objects destructor
void Int_public_Die(class_Int* this) {}

// defintion for a string.Resize() method
int Int_public_GetValue(class_Int* this) {
	return this->value;
}

// -----------------------------------------------------------------------------
// "float" object type
// Note: this is an object version of a float, this is to box and crunch it
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
typedef struct {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Float"
	void (*Function)(class_Float*)  // destructor function pointer
} Float_vTable;

// the objects struct
typedef struct {
	const Float_vTable* vtable;  // our vTable
	float value;
} class_Float;

// definition for the Float vTable
const Float_vTable Float_vTable_Const = {&Any_vTable_Const, "Float", &Float_public_Die};

// defintion for the objects constructor
void Float_public_Contructor(class_Float* this, float value) {
	this->vtable = &Float_vTable_Const;
	this->value = value;
}

// defintion for the objects destructor
void Float_public_Die(class_Float* this) {}

// defintion for a string.Resize() method
float Float_public_GetValue(class_Float* this) {
	return this->value;
}

// -----------------------------------------------------------------------------
// "bool" object type
// Note: this is an object version of a bool, this is to box and crunch it
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
typedef struct {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Bool"
	void (*Function)(class_Bool*)   // destructor function pointer
} Bool_vTable;

// the objects struct
typedef struct {
	const Bool_vTable* vtable;  // our vTable
	bool value;
} class_Bool;

// definition for the Bool vTable
const Bool_vTable Bool_vTable_Const = {&Any_vTable_Const, "Bool", &Bool_public_Contructor};

// defintion for the objects constructor
void Bool_public_Contructor(class_Bool* this, bool value) {
	this->vtable = &Bool_vTable_Const;
	this->value = value;
}

// defintion for the objects destructor
void Bool_public_Die(class_Bool* this) {}

// defintion for a string.Resize() method
bool Bool_public_GetValue(class_Bool* this) {
	return this->value;
}