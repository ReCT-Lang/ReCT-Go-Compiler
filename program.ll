%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_Array = type { %struct.String_vTable*, i32, %struct.class_Any**, i32, i32, i32 }
%struct.class_Byte = type { %struct.String_vTable*, i32, i8 }
%struct.class_Float = type { %struct.String_vTable*, i32, float }
%struct.class_Int = type { %struct.String_vTable*, i32, i32 }
%struct.class_Long = type { %struct.String_vTable*, i32, i64 }
%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.class_Thread = type { %struct.Any_vTable*, i32, i8* (i8*)*, i8*, i64 }
%struct.class_pArray = type { %struct.String_vTable*, i32, i8*, i32, i32, i32, i32 }
%struct.ABC_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_ABC = type { %struct.ABC_vTable*, i32, i32 }

@Any_vTable_Const = external global %struct.Any_vTable
@String_vTable_Const = external global %struct.String_vTable
@.str.c.0 = constant [4 x i8] c"ABC\00"
@ABC_vTable_Const = global %struct.ABC_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr ([4 x i8], [4 x i8]* @.str.c.0, i32 0, i32 0), void (i8*)* @ABC_public_Die }

declare i32 @printf(i8* %format, ...)

declare i32 @scanf(i8* %format, i8* %dest, ...)

declare void @strcpy(i8* %dest, i8* %src)

declare void @strcat(i8* %dest, i8* %src)

declare i32 @strlen(i8* %str)

declare i32 @strcmp(i8* %left, i8* %right)

declare i8* @malloc(i32 %len)

declare void @free(i8* %dest)

declare i32 @snprintf(i8* %dest, i32 %len, i8* %format, ...)

declare i32 @atoi(i8* %str)

declare double @atof(i8* %str)

declare void @Any_public_Constructor(%struct.class_Any* noundef %0)

declare void @Any_public_Die(i8* noundef %0)

declare void @Float_public_Constructor(%struct.class_Float* noundef %0, float noundef %1)

declare void @Float_public_Die(i8* noundef %0)

declare float @Float_public_GetValue(%struct.class_Float* noundef %0)

declare void @Int_public_Constructor(%struct.class_Int* noundef %0, i32 noundef %1)

declare void @Int_public_Die(i8* noundef %0)

declare i32 @Int_public_GetValue(%struct.class_Int* noundef %0)

declare void @String_public_Constructor(%struct.class_String* noundef %0)

declare void @String_public_Die(i8* noundef %0)

declare void @String_public_Load(%struct.class_String* noundef %0, i8* noundef %1)

declare void @String_public_Resize(%struct.class_String* noundef %0, i32 noundef %1)

declare void @String_public_AddChar(%struct.class_String* noundef %0, i8 noundef signext %1)

declare %struct.class_String* @String_public_Concat(%struct.class_String* noundef %0, %struct.class_String* noundef %1)

declare i1 @String_public_Equal(%struct.class_String* noundef %0, %struct.class_String* noundef %1)

declare i8* @String_public_GetBuffer(%struct.class_String* noundef %0)

declare i32 @String_public_GetLength(%struct.class_String* noundef %0)

declare %struct.class_String* @String_public_Substring(%struct.class_String* noundef %0, i32 noundef %1, i32 noundef %2)

declare void @Thread_public_Constructor(%struct.class_Thread* noundef %0, i8* (i8*)* noundef %1, i8* noundef %2)

declare void @Thread_public_Die(i8* noundef %0)

declare void @Thread_public_Start(%struct.class_Thread* noundef %0)

declare void @Thread_public_Join(%struct.class_Thread* noundef %0)

declare void @Thread_public_Kill(%struct.class_Thread* noundef %0)

declare void @Array_public_Constructor(%struct.class_Array* noundef %0, i32 noundef %1)

declare void @Array_public_Die(i8* noundef %0)

declare %struct.class_Any* @Array_public_GetElement(%struct.class_Array* noundef %0, i32 noundef %1)

declare void @Array_public_SetElement(%struct.class_Array* noundef %0, i32 noundef %1, %struct.class_Any* noundef %2)

declare i32 @Array_public_GetLength(%struct.class_Array* noundef %0)

declare void @Array_public_Push(%struct.class_Array* noundef %0, %struct.class_Any* noundef %1)

declare void @Byte_public_Constructor(%struct.class_Byte* noundef %0, i8 noundef signext %1)

declare void @Byte_public_Die(i8* noundef %0)

declare i8 @Byte_public_GetValue(%struct.class_Byte* noundef %0)

declare void @Long_public_Constructor(%struct.class_Long* noundef %0, i64 noundef %1)

declare void @Long_public_Die(i8* noundef %0)

declare i64 @Long_public_GetValue(%struct.class_Long* noundef %0)

declare void @pArray_public_Constructor(%struct.class_pArray* noundef %0, i32 noundef %1, i32 noundef %2)

declare void @pArray_public_Die(i8* noundef %0)

declare i32 @pArray_public_GetLength(%struct.class_pArray* noundef %0)

declare i8* @pArray_public_Grow(%struct.class_pArray* noundef %0)

declare i8* @pArray_public_GetElementPtr(%struct.class_pArray* noundef %0, i32 noundef %1)

declare void @arc_RegisterReference(%struct.class_Any* noundef %0)

declare void @arc_UnregisterReference(%struct.class_Any* noundef %0)

declare void @arc_DestroyObject(%struct.class_Any* noundef %0)

declare void @arc_RegisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1)

declare void @arc_UnregisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1)

declare void @exc_Throw(i8* noundef %0)

declare void @exc_ThrowIfNull(i8* noundef %0)

declare void @exc_ThrowIfInvalidCast(%struct.class_Any* noundef %0, %struct.Any_vTable* noundef %1)

define void @ABC_public_Die(i8* %obj) {
0:
	%1 = bitcast i8* %obj to %struct.class_ABC*
	; <DieARC>
	; </DieARC>
	br label %$decl

$decl:
	ret void
}

define void @ABC_public_Constructor(%struct.class_ABC* %me) {
0:
	%1 = alloca %struct.class_ABC*
	store %struct.class_ABC* %me, %struct.class_ABC** %1
	%2 = load %struct.class_ABC*, %struct.class_ABC** %1
	%3 = getelementptr %struct.class_ABC, %struct.class_ABC* %2, i32 0, i32 0
	store %struct.ABC_vTable* @ABC_vTable_Const, %struct.ABC_vTable** %3
	%4 = getelementptr %struct.class_ABC, %struct.class_ABC* %2, i32 0, i32 1
	store i32 0, i32* %4
	br label %semiroot

semiroot:
	; <ReturnARC>
	; </ReturnARC>
	ret void
}

define void @main() {
0:
	%obj_ABC = alloca %struct.class_ABC*
	%var_x = alloca %struct.class_ABC*
	%load_var_x = alloca %struct.class_ABC*
	br label %semiroot

semiroot:
	; <... object init (rect stuff)>
	store %struct.class_ABC* %5, %struct.class_ABC** %obj_ABC
	store %struct.class_ABC* null, %struct.class_ABC** %var_x ; explicit init (not actually necessary)
	%7 = load %struct.class_ABC*, %struct.class_ABC** %var_x
	store %struct.class_ABC* %7, %struct.class_ABC** %load_var_x
	; <ReturnARC>
	ret void
}
