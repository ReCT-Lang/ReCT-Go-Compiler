#include<stdbool.h>

#ifndef OBJECTS_H
#define OBJECTS_H

// declare all struct names
typedef struct Any_vTable    Any_vTable;
typedef struct class_Any     class_Any;
typedef struct String_vTable String_vTable;
typedef struct class_String  class_String;
typedef struct Int_vTable    Int_vTable;
typedef struct class_Int     class_Int;
typedef struct Float_vTable  Float_vTable;
typedef struct class_Float   class_Float;
typedef struct Bool_vTable   Bool_vTable;
typedef struct class_Bool    class_Bool;

// declare destructor function pointer
typedef void (*DiePointer)(void*);

// declare all destructors
void Any_public_Die   (void*);
void String_public_Die(void*);
void Int_public_Die   (void*);
void Float_public_Die (void*);
void Bool_public_Die  (void*);

// -----------------------------------------------------------------------------
// base "any" object type
// Note: all object types will inherit from this
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct Any_vTable {
	const void* parentVTable; // will be NULL for "any" as its the root
	const char* className;                 // will be "Any"
	DiePointer dieFunction;                 // destructor function pointer
};

// the objects struct
struct class_Any {
	const Any_vTable* vtable; // "any" is pretty boring, it doesnt store any data
	int referenceCounter;     // except for the objects reference counter (required for ARC)
};

// -----------------------------------------------------------------------------
// "string" object type
// Note: this is our wrapper for strings!
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct String_vTable {
	const Any_vTable* parentVTable;  // will be a pointer to the "any" vTable
	const char* className;           // will be "String"
	DiePointer dieFunction;           // destructor function pointer
};

// the objects struct
struct class_String {
	const String_vTable* vtable;  // our vTable
	int referenceCounter;   // implementation of the reference counter
	char* buffer;           // for info on this string implementation check out:
	int length;             // https://mapping-high-level-constructs-to-llvm-ir.readthedocs.io/en/latest/appendix-a-how-to-implement-a-string-type-in-llvm/index.html
	int maxLen;
	int factor;
};

// -----------------------------------------------------------------------------
// "int" object type
// Note: this is an object version of an int, this is to box and crunch it
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct Int_vTable {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Int"
	DiePointer dieFunction;          // destructor function pointer
};

// the objects struct
struct class_Int {
	const Int_vTable* vtable;  // our vTable
	int referenceCounter;      // implementation of the reference counter
	int value;
};

// -----------------------------------------------------------------------------
// "float" object type
// Note: this is an object version of a float, this is to box and crunch it
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct Float_vTable {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Float"
	DiePointer dieFunction;          // destructor function pointer
};

// the objects struct
struct class_Float {
	const Float_vTable* vtable;  // our vTable
	int referenceCounter;   // implementation of the reference counter
	float value;
};

// -----------------------------------------------------------------------------
// "bool" object type
// Note: this is an object version of a bool, this is to box and crunch it
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct Bool_vTable {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Bool"
	DiePointer dieFunction;          // destructor function pointer
};

// the objects struct
struct class_Bool {
	const Bool_vTable* vtable;  // our vTable
	int referenceCounter;   // implementation of the reference counter
	bool value;
};

#endif