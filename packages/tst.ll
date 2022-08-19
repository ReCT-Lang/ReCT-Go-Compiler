%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_Array = type { %struct.String_vTable*, i32, %struct.class_Any**, i32, i32, i32 }
%struct.class_Bool = type { %struct.String_vTable*, i32, i8 }
%struct.class_Float = type { %struct.String_vTable*, i32, float }
%struct.class_Int = type { %struct.String_vTable*, i32, i32 }
%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.class_Thread = type { %struct.Any_vTable*, i32, i8* (i8*)*, i8*, i64 }
%struct.class_pArray = type { %struct.String_vTable*, i32, i8*, i32, i32, i32, i32 }
%struct.TestClass_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_TestClass = type { %struct.TestClass_vTable*, i32, %struct.class_String*, i32 }

@Any_vTable_Const = external global %struct.Any_vTable
@String_vTable_Const = external global %struct.String_vTable
@.tst.str.c.0 = constant [9 x i8] c"FieldOne\00"
@.tst.str.c.1 = constant [9 x i8] c"FieldTwo\00"
@TestClass_Fields_Const = global [2 x i8*] [i8* getelementptr ([9 x i8], [9 x i8]* @.tst.str.c.0, i32 0, i32 0), i8* getelementptr ([9 x i8], [9 x i8]* @.tst.str.c.1, i32 0, i32 0)]
@.tst.str.c.2 = constant [10 x i8] c"TestClass\00"
@TestClass_vTable_Const = global %struct.TestClass_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr ([10 x i8], [10 x i8]* @.tst.str.c.2, i32 0, i32 0), void (i8*)* @TestClass_public_Die }
@.tst.str.3 = constant [2 x i8] c"a\00"

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

declare void @pArray_public_Constructor(%struct.class_pArray* noundef %0, i32 noundef %1, i32 noundef %2)

declare void @pArray_public_Die(i8* noundef %0)

declare i32 @pArray_public_GetLength(%struct.class_pArray* noundef %0)

declare i8* @pArray_public_Grow(%struct.class_pArray* noundef %0)

declare i8* @pArray_public_GetElementPtr(%struct.class_pArray* noundef %0, i32 noundef %1)

declare void @Any_public_Constructor(%struct.class_Any* noundef %0)

declare void @Any_public_Die(i8* noundef %0)

declare void @Array_public_Constructor(%struct.class_Array* noundef %0, i32 noundef %1)

declare void @Array_public_Die(i8* noundef %0)

declare %struct.class_Any* @Array_public_GetElement(%struct.class_Array* noundef %0, i32 noundef %1)

declare void @Array_public_SetElement(%struct.class_Array* noundef %0, i32 noundef %1, %struct.class_Any* noundef %2)

declare i32 @Array_public_GetLength(%struct.class_Array* noundef %0)

declare void @Array_public_Push(%struct.class_Array* noundef %0, %struct.class_Any* noundef %1)

declare void @Bool_public_Constructor(%struct.class_Bool* noundef %0, i1 noundef zeroext %1)

declare void @Bool_public_Die(i8* noundef %0)

declare i1 @Bool_public_GetValue(%struct.class_Bool* noundef %0)

declare void @Float_public_Constructor(%struct.class_Float* noundef %0, float noundef %1)

declare void @Float_public_Die(i8* noundef %0)

declare float @Float_public_GetValue(%struct.class_Float* noundef %0)

declare void @arc_RegisterReference(%struct.class_Any* noundef %0)

declare void @arc_UnregisterReference(%struct.class_Any* noundef %0)

declare void @arc_DestroyObject(%struct.class_Any* noundef %0)

declare void @arc_RegisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1)

declare void @arc_UnregisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1)

declare void @exc_Throw(i8* noundef %0)

declare void @exc_ThrowIfNull(i8* noundef %0)

declare void @sys_Print(%struct.class_String* noundef %0)

declare void @sys_Write(%struct.class_String* noundef %0)

declare %struct.class_String* @sys_Input()

declare void @sys_Clear()

declare void @sys_SetCursor(i32 noundef %0, i32 noundef %1)

declare void @sys_SetCursorVisible(i1 noundef zeroext %0)

declare i1 @sys_GetCursorVisible()

declare i32 @sys_Random(i32 noundef %0)

declare void @sys_Sleep(i32 noundef %0)

declare i32 @sys_Sqrt(i32 noundef %0)

declare i32 @sys_Now()

declare %struct.class_String* @sys_Char(i32 noundef %0)

define void @TestClass_public_Die(i8* %obj) {
0:
	%1 = bitcast i8* %obj to %struct.class_TestClass*
	; <DieARC>
	; -> destroying reference to 'VG_14 [Field 2]'
	%2 = getelementptr %struct.class_TestClass, %struct.class_TestClass* %1, i32 0, i32 2
	%3 = load %struct.class_String*, %struct.class_String** %2
	%4 = bitcast %struct.class_String* %3 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %4)
	; </DieARC>
	br label %$decl

$decl:
	ret void
}

define void @TestClass_public_Constructor(%struct.class_TestClass* %me) {
0:
	%1 = alloca %struct.class_TestClass*
	store %struct.class_TestClass* %me, %struct.class_TestClass** %1
	%2 = load %struct.class_TestClass*, %struct.class_TestClass** %1
	%3 = getelementptr %struct.class_TestClass, %struct.class_TestClass* %2, i32 0, i32 0
	store %struct.TestClass_vTable* @TestClass_vTable_Const, %struct.TestClass_vTable** %3
	%4 = getelementptr %struct.class_TestClass, %struct.class_TestClass* %2, i32 0, i32 1
	store i32 0, i32* %4
	%5 = getelementptr %struct.class_TestClass, %struct.class_TestClass* %me, i32 0, i32 2
	store %struct.class_String* null, %struct.class_String** %5
	br label %semiroot

semiroot:
	; <ReturnARC>
	; </ReturnARC>
	ret void
}

define void @tst_PrintString(%struct.class_String* %VP_16) {
0:
	br label %semiroot

semiroot:
	%1 = getelementptr [2 x i8], [2 x i8]* @.tst.str.3, i32 0, i32 0
	%2 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%3 = ptrtoint %struct.class_String* %2 to i32
	%4 = call i8* @malloc(i32 %3)
	%5 = bitcast i8* %4 to %struct.class_String*
	%6 = getelementptr %struct.class_String, %struct.class_String* %5, i32 0
	call void @String_public_Constructor(%struct.class_String* %6)
	%7 = bitcast %struct.class_String* %6 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %7)
	call void @String_public_Load(%struct.class_String* %6, i8* %1)
	call void @sys_Print(%struct.class_String* %6)
	%8 = bitcast %struct.class_String* %6 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %8)
	; <ReturnARC>
	;  -> destroying reference to '%someString'
	%9 = bitcast %struct.class_String* %VP_16 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %9)
	; </ReturnARC>
	ret void
}
