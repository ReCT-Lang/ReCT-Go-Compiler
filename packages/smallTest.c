#include <stdlib.h>
#include "./../systemlib/objects.h"
#include "./../systemlib/arc.h"

class_String *test_GetString() {
	// create a string object
	class_String *strInstance = (class_String*)malloc(sizeof(class_String));
	String_public_Constructor(strInstance);
	arc_RegisterReference((class_Any*)strInstance);

	// load "cool string" into it
	String_public_Load(strInstance, "cool string");

	// return it
	return strInstance;
}