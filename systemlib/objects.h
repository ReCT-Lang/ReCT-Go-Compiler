#include<stdbool.h>
#include "pthread.h"

#ifndef OBJECTS_H
#define OBJECTS_H

// declare all struct names
typedef struct Any_vTable    Any_vTable;
typedef struct class_Any     class_Any;
typedef struct String_vTable String_vTable;
typedef struct class_String  class_String;
typedef struct Int_vTable    Int_vTable;
typedef struct class_Int     class_Int;
typedef struct Byte_vTable    Byte_vTable;
typedef struct class_Byte     class_Byte;
typedef struct Long_vTable    Long_vTable;
typedef struct class_Long     class_Long;
typedef struct Float_vTable  Float_vTable;
typedef struct class_Float   class_Float;
typedef struct Bool_vTable   Bool_vTable;
typedef struct class_Bool    class_Bool;
typedef struct Array_vTable  Array_vTable;
typedef struct class_Array   class_Array;
typedef struct pArray_vTable  pArray_vTable;
typedef struct class_pArray   class_pArray;
typedef struct class_Thread class_Thread;
typedef struct Thread_vTable Thread_vTable;

// declare destructor function pointer
typedef void (*DiePointer)(void*);

// declare all destructors
void Any_public_Die    (void*);
void String_public_Die (void*);
void Int_public_Die    (void*);
void Byte_public_Die   (void*);
void Long_public_Die   (void*);
void Float_public_Die  (void*);
void Bool_public_Die   (void*);
void Array_public_Die  (void*);
void pArray_public_Die (void*);
void Thread_public_Die (void*);

// declare all constructors
void Any_public_Constructor(class_Any*);
void String_public_Constructor(class_String*);
void Int_public_Constructor(class_Int*, int);
void Byte_public_Constructor(class_Byte*, char);
void Long_public_Constructor(class_Long*, long);
void Float_public_Constructor(class_Float*, float);
void Bool_public_Constructor(class_Bool*, bool);
void Array_public_Constructor(class_Array*, int);
void pArray_public_Constructor(class_pArray*, int, int);
void Thread_public_Constructor(class_Thread*, void *(*)(void *), void *);

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

// the objects methods
void String_public_Load(class_String*, char*);
void String_public_Resize(class_String*, int);
void String_public_AddChar(class_String*, char);
class_String* String_public_Concat(class_String*, class_String*);
bool String_public_Equal(class_String*, class_String*);
char *String_public_GetBuffer(class_String*);
int String_public_GetLength(class_String*);
class_String *String_public_Substring(class_String*, int, int);

// -----------------------------------------------------------------------------
// "int" object type
// Note: this is an object version of an int, this is to box and crunch it
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct Int_vTable {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Int"
	DiePointer dieFunction;         // destructor function pointer
};

// the objects struct
struct class_Int {
	const Int_vTable* vtable;  // our vTable
	int referenceCounter;      // implementation of the reference counter
	int value;
};

// the objects methods
int Int_public_GetValue(class_Int*);

// -----------------------------------------------------------------------------
// "byte" object type
// Note: this is an object version of a byte, this is to box and crunch it
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct Byte_vTable {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Byte"
	DiePointer dieFunction;         // destructor function pointer
};

// the objects struct
struct class_Byte {
	const Byte_vTable* vtable;  // our vTable
	int referenceCounter;       // implementation of the reference counter
	char value;
};

// the objects methods
char Byte_public_GetValue(class_Byte*);

// -----------------------------------------------------------------------------
// "long" object type
// Note: this is an object version of a long, this is to box and crunch it
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct Long_vTable {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Long"
	DiePointer dieFunction;         // destructor function pointer
};

// the objects struct
struct class_Long {
	const Long_vTable* vtable;  // our vTable
	int referenceCounter;       // implementation of the reference counter
	long value;
};

// the objects methods
long Long_public_GetValue(class_Long*);

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

// the objects methods
float Float_public_GetValue(class_Float*);

// -----------------------------------------------------------------------------
// "bool" object type
// Note: this is an object version of a bool, this is to box and crunch it
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct Bool_vTable {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Bool"
	DiePointer dieFunction;         // destructor function pointer
};

// the objects struct
struct class_Bool {
	const Bool_vTable* vtable;  // our vTable
	int referenceCounter;   // implementation of the reference counter
	bool value;
};

// the objects methods
bool Bool_public_GetValue(class_Bool*);

// -----------------------------------------------------------------------------
// "array" object type
// Note: this is an object represents an array, it only holds object types atm
// The array wont make data copies, it will just hold references
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct Array_vTable {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "Array"
	DiePointer dieFunction;         // destructor function pointer
};

// the objects struct
struct class_Array {
	const Array_vTable* vtable;  // our vTable
	int referenceCounter;        // implementation of the reference counter
	class_Any **elements;        // marks the start of our array
	int length;					 // the length of this array
	int maxLen;                  // buffer length
	int factor;                  // growth factor 
};

// the objects methods
class_Any* Array_public_GetElement(class_Array*, int);
void Array_public_SetElement(class_Array*, int, class_Any*);
int Array_public_GetLength(class_Array*);
void Array_public_Push(class_Array*, class_Any*);

// helper for lazy people (me)
#define DEFINE_ARRAY(class)                                 \
	typedef struct class_Array_##class class_Array_##class; \
	struct class_Array_##class {                            \
		const Array_vTable* vtable;                         \
		int referenceCounter;                               \
		class_Any **elements;                               \
		int length;                                         \
		int maxLen;                                         \
		int factor;                                         \
	};

// predefined Array Types
DEFINE_ARRAY(String);
DEFINE_ARRAY(Any);

// -----------------------------------------------------------------------------
// "parray" object type
// Note: this is a primitive version of "array"
// -----------------------------------------------------------------------------

// the object's vtable (for method lookup and method overriding)
struct pArray_vTable {
	const Any_vTable* parentVTable; // will be a pointer to the "any" vTable
	const char* className;          // will be "pArray"
	DiePointer dieFunction;         // destructor function pointer
};

// the objects struct
struct class_pArray {
	const pArray_vTable* vtable;  // our vTable
	int referenceCounter;         // implementation of the reference counter
	void *elements;               // marks the start of our array
	int length;                   // array length
	int maxLen;                   // buffer length
	int factor;                   // growth factor 
	int elemSize;                 // size of one element 
};

// the objects methods
int pArray_public_GetLength(class_pArray*);
void *pArray_public_Grow(class_pArray*);
void *pArray_public_GetElementPtr(class_pArray*, int);

// helper for lazy people (me)
#define DEFINE_PARRAY(type)                                   \
	typedef struct class_pArray_##type class_pArray_##type;   \
	struct class_pArray_##type {                              \
		const pArray_vTable* vtable;                          \
		int referenceCounter;                                 \
		void *elements;                                       \
		int length;                                           \
		int maxLen;                                           \
		int factor;                                           \
		int elemSize;                                         \
	};

// predefined pArray Types
DEFINE_PARRAY(Bool);
DEFINE_PARRAY(Byte);
DEFINE_PARRAY(Int);
DEFINE_PARRAY(Float);

// -----------------------------------------------------------------------------
// base "thread" object type
// Note: Multithreading!
// Developer Note: This requires -lpthread flag because we're using pthread.h
// -----------------------------------------------------------------------------

// Thread are a recreation using the old documentation: https://docs.rect.ml/threading

// the object's vtable (for method lookup and method overriding)
struct Thread_vTable {
	const void* parentVTable;              // will be a pointer to the "any" vTable
	const char* className;                 // will be "Thread" (I think)
	DiePointer dieFunction;                // destructor function pointer
};

// the objects struct
struct class_Thread {
	const Thread_vTable* vtable;   // the epic vTable
	int referenceCounter;       // you guessed it, reference counter for the ARc
	void *(*__routine)(void*);  // thread routine (this is the function the thread runs)
	void *args;                 // (the arguments to the function the thread runs)
	pthread_t id;               // thread id
};

// the objects methods
void Thread_public_Start(class_Thread*);
void Thread_public_Join(class_Thread*);
void Thread_public_Kill(class_Thread*);

#endif