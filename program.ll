%struct.Standard_vTable = type { i8*, i8*, void (i8*)*, i8* }
%struct.class_Any = type { %struct.Standard_vTable, i32 }
%struct.class_Array = type { %struct.Standard_vTable, i32, %struct.class_Any**, i32, i32, i32 }
%struct.class_Byte = type { %struct.Standard_vTable, i32, i8 }
%struct.class_Float = type { %struct.Standard_vTable, i32, float }
%struct.class_Int = type { %struct.Standard_vTable, i32, i32 }
%struct.class_Long = type { %struct.Standard_vTable, i32, i64 }
%struct.class_String = type { %struct.Standard_vTable, i32, i8*, i32, i32, i32 }
%struct.class_Thread = type { %struct.Standard_vTable, i32, i8* (i8*)*, i8*, i64 }
%struct.class_pArray = type { %struct.Standard_vTable, i32, i8*, i32, i32, i32, i32 }
%struct.class_SomeClass = type { %struct.Standard_vTable, i32, %struct.class_String* }

@Any_vTable_Const = external global %struct.Standard_vTable
@Array_vTable_Const = external global %struct.Standard_vTable
@Byte_vTable_Const = external global %struct.Standard_vTable
@Float_vTable_Const = external global %struct.Standard_vTable
@Int_vTable_Const = external global %struct.Standard_vTable
@Long_vTable_Const = external global %struct.Standard_vTable
@String_vTable_Const = external global %struct.Standard_vTable
@Thread_vTable_Const = external global %struct.Standard_vTable
@pArray_vTable_Const = external global %struct.Standard_vTable
@.str.c.0 = constant [10 x i8] c"SomeClass\00"
@SomeClass_vTable_Const = global %struct.Standard_vTable { i8* null, i8* getelementptr ([10 x i8], [10 x i8]* @.str.c.0, i32 0, i32 0), void (i8*)* @SomeClass_public_Die, i8* null }
@.str.c.1 = constant [10 x i8] c"SomeClass\00"
@.str.c.2 = constant [16 x i8] c"TO_SomeClass_[]\00"
@.str.3 = constant [14 x i8] c"cooler string\00"
@.str.c.4 = constant [7 x i8] c"String\00"
@.str.c.5 = constant [13 x i8] c"TO_string_[]\00"
@.str.6 = constant [12 x i8] c"cool string\00"
@.str.c.7 = constant [7 x i8] c"String\00"
@.str.c.8 = constant [13 x i8] c"TO_string_[]\00"
@.str.c.9 = constant [5 x i8] c"Long\00"
@.str.c.10 = constant [10 x i8] c"T_long_[]\00"
@.str.c.11 = constant [10 x i8] c"T_long_[]\00"

declare i8* @malloc(i32 %len)

declare void @free(i8* %dest)

declare i32 @printf(i8* %format, ...)

declare i32 @snprintf(i8* %dest, i32 %len, i8* %format, ...)

declare i32 @atoi(i8* %str)

declare i64 @atol(i8* %str)

declare double @atof(i8* %str)

declare void @Any_public_Constructor(%struct.class_Any* noundef %0)

declare void @Any_public_Die(i8* noundef %0)

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

declare void @Float_public_Constructor(%struct.class_Float* noundef %0, float noundef %1)

declare void @Float_public_Die(i8* noundef %0)

declare float @Float_public_GetValue(%struct.class_Float* noundef %0)

declare void @Int_public_Constructor(%struct.class_Int* noundef %0, i32 noundef %1)

declare void @Int_public_Die(i8* noundef %0)

declare i32 @Int_public_GetValue(%struct.class_Int* noundef %0)

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

declare void @arc_RegisterReference(%struct.class_Any* noundef %0)

declare void @arc_UnregisterReference(%struct.class_Any* noundef %0)

declare void @arc_DestroyObject(%struct.class_Any* noundef %0)

declare void @arc_RegisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1)

declare void @arc_UnregisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1)

declare void @exc_Throw(i8* noundef %0)

declare void @exc_ThrowIfNull(i8* noundef %0)

declare void @exc_ThrowIfInvalidCast(%struct.class_Any* noundef %0, %struct.Standard_vTable* noundef %1, i8* noundef %2)

declare void @llvm.dbg.declare(metadata %0, metadata %1, metadata %2)

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

define void @SomeClass_public_Die(i8* %obj) {
0:
	%1 = bitcast i8* %obj to %struct.class_SomeClass*
	; <DieARC>
	; -> destroying reference to 'VG_14 [Field 2]'
	%2 = getelementptr %struct.class_SomeClass, %struct.class_SomeClass* %1, i32 0, i32 2
	%3 = load %struct.class_String*, %struct.class_String** %2
	%4 = bitcast %struct.class_String* %3 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %4)
	; </DieARC>
	br label %$decl

$decl:
	ret void
}

define void @SomeClass_public_Constructor(%struct.class_SomeClass* %me) {
0:
	%1 = alloca %struct.class_SomeClass*
	store %struct.class_SomeClass* %me, %struct.class_SomeClass** %1
	%2 = getelementptr %struct.class_SomeClass, %struct.class_SomeClass* %me, i32 0, i32 2
	store %struct.class_String* null, %struct.class_String** %2
	%VL_16 = alloca void (%struct.class_SomeClass*)*
	%VL_17 = alloca %struct.class_Any*
	br label %semiroot

semiroot:
	; <BoundAssignmentExpression>
	; <BoundLiteralExpression>
	%3 = getelementptr [12 x i8], [12 x i8]* @.str.6, i32 0, i32 0
	%4 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%5 = ptrtoint %struct.class_String* %4 to i32
	%6 = call i8* @malloc(i32 %5)
	%7 = bitcast i8* %6 to %struct.class_String*
	%8 = getelementptr %struct.class_String, %struct.class_String* %7, i32 0, i32 1
	store i32 0, i32* %8
	%9 = getelementptr %struct.class_String, %struct.class_String* %7, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([7 x i8], [7 x i8]* @.str.c.7, i32 0, i32 0), void (i8*)* @String_public_Die, i8* getelementptr ([13 x i8], [13 x i8]* @.str.c.8, i32 0, i32 0) }, %struct.Standard_vTable* %9
	%10 = bitcast %struct.class_String* %7 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %10)
	call void @String_public_Constructor(%struct.class_String* %7)
	call void @String_public_Load(%struct.class_String* %7, i8* %3)
	; </BoundLiteralExpression>
	%11 = getelementptr %struct.class_SomeClass, %struct.class_SomeClass* %me, i32 0, i32 2
	%12 = load %struct.class_String*, %struct.class_String** %11
	%13 = bitcast %struct.class_String* %12 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %13)
	store %struct.class_String* %7, %struct.class_String** %11
	%14 = bitcast %struct.class_String* %7 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %14)
	; </BoundAssignmentExpression>
	; expression value unused -> destroying reference
	%15 = bitcast %struct.class_String* %7 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %15)
	; <BoundFunctionExpression>
	; </BoundFunctionExpression>
	store void (%struct.class_SomeClass*)* @SomeClass_private_F_someOtherFunction_void, void (%struct.class_SomeClass*)** %VL_16
	; <BoundConversionExpression>
	; <BoundVariableExpression>
	%16 = load void (%struct.class_SomeClass*)*, void (%struct.class_SomeClass*)** %VL_16
	; </BoundVariableExpression>
	%17 = ptrtoint void (%struct.class_SomeClass*)* %16 to i64
	%18 = getelementptr %struct.class_Long, %struct.class_Long* null, i32 1
	%19 = ptrtoint %struct.class_Long* %18 to i32
	%20 = call i8* @malloc(i32 %19)
	%21 = bitcast i8* %20 to %struct.class_Long*
	%22 = getelementptr %struct.class_Long, %struct.class_Long* %21, i32 0, i32 1
	store i32 0, i32* %22
	%23 = getelementptr %struct.class_Long, %struct.class_Long* %21, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([5 x i8], [5 x i8]* @.str.c.9, i32 0, i32 0), void (i8*)* @Long_public_Die, i8* getelementptr ([10 x i8], [10 x i8]* @.str.c.10, i32 0, i32 0) }, %struct.Standard_vTable* %23
	%24 = bitcast %struct.class_Long* %21 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %24)
	call void @Long_public_Constructor(%struct.class_Long* %21, i64 %17)
	%25 = bitcast %struct.class_Long* %21 to %struct.class_Any*
	; </BoundConversionExpression>
	store %struct.class_Any* %25, %struct.class_Any** %VL_17
	; <BoundTypeCallExpression>
	; <BoundConversionExpression>
	; <BoundVariableExpression>
	%26 = load %struct.class_Any*, %struct.class_Any** %VL_17
	; </BoundVariableExpression>
	%27 = bitcast %struct.class_Any* %26 to %struct.class_Any*
	%28 = bitcast %struct.Standard_vTable* @Long_vTable_Const to %struct.Standard_vTable*
	call void @exc_ThrowIfInvalidCast(%struct.class_Any* %27, %struct.Standard_vTable* %28, i8* getelementptr ([10 x i8], [10 x i8]* @.str.c.11, i32 0, i32 0))
	%29 = bitcast %struct.class_Any* %26 to %struct.class_Long*
	%30 = call i64 @Long_public_GetValue(%struct.class_Long* %29)
	%31 = inttoptr i64 %30 to void (%struct.class_SomeClass*)*
	; </BoundConversionExpression>
	; <BoundThisExpression>
	; </BoundThisExpression>
	%32 = bitcast %struct.class_SomeClass* %me to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %32)
	call void %31(%struct.class_SomeClass* %me)
	%33 = bitcast %struct.class_SomeClass* %me to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %33)
	%34 = bitcast void (%struct.class_SomeClass*)* %31 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %34)
	; </BoundTypeCallExpression>
	; <ReturnARC>
	;  -> destroying reference to '%VL_17'
	%35 = load %struct.class_Any*, %struct.class_Any** %VL_17
	%36 = bitcast %struct.class_Any* %35 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %36)
	; </ReturnARC>
	ret void
}

define void @SomeClass_private_F_someOtherFunction_void(%struct.class_SomeClass* %$me) {
0:
	br label %semiroot

semiroot:
	; <BoundAssignmentExpression>
	; <BoundLiteralExpression>
	%1 = getelementptr [14 x i8], [14 x i8]* @.str.3, i32 0, i32 0
	%2 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%3 = ptrtoint %struct.class_String* %2 to i32
	%4 = call i8* @malloc(i32 %3)
	%5 = bitcast i8* %4 to %struct.class_String*
	%6 = getelementptr %struct.class_String, %struct.class_String* %5, i32 0, i32 1
	store i32 0, i32* %6
	%7 = getelementptr %struct.class_String, %struct.class_String* %5, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([7 x i8], [7 x i8]* @.str.c.4, i32 0, i32 0), void (i8*)* @String_public_Die, i8* getelementptr ([13 x i8], [13 x i8]* @.str.c.5, i32 0, i32 0) }, %struct.Standard_vTable* %7
	%8 = bitcast %struct.class_String* %5 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %8)
	call void @String_public_Constructor(%struct.class_String* %5)
	call void @String_public_Load(%struct.class_String* %5, i8* %1)
	; </BoundLiteralExpression>
	%9 = getelementptr %struct.class_SomeClass, %struct.class_SomeClass* %$me, i32 0, i32 2
	%10 = load %struct.class_String*, %struct.class_String** %9
	%11 = bitcast %struct.class_String* %10 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %11)
	store %struct.class_String* %5, %struct.class_String** %9
	%12 = bitcast %struct.class_String* %5 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %12)
	; </BoundAssignmentExpression>
	; expression value unused -> destroying reference
	%13 = bitcast %struct.class_String* %5 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %13)
	; <ReturnARC>
	; </ReturnARC>
	ret void
}

define void @main() {
0:
	%VL_15 = alloca %struct.class_SomeClass*
	br label %semiroot

semiroot:
	; <BoundMakeExpression>
	%1 = getelementptr %struct.class_SomeClass, %struct.class_SomeClass* null, i32 1
	%2 = ptrtoint %struct.class_SomeClass* %1 to i32
	%3 = call i8* @malloc(i32 %2)
	%4 = bitcast i8* %3 to %struct.class_SomeClass*
	%5 = getelementptr %struct.class_SomeClass, %struct.class_SomeClass* %4, i32 0, i32 1
	store i32 0, i32* %5
	%6 = getelementptr %struct.class_SomeClass, %struct.class_SomeClass* %4, i32 0, i32 0
	store %struct.Standard_vTable { i8* null, i8* getelementptr ([10 x i8], [10 x i8]* @.str.c.1, i32 0, i32 0), void (i8*)* @SomeClass_public_Die, i8* getelementptr ([16 x i8], [16 x i8]* @.str.c.2, i32 0, i32 0) }, %struct.Standard_vTable* %6
	%7 = bitcast %struct.class_SomeClass* %4 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %7)
	call void @SomeClass_public_Constructor(%struct.class_SomeClass* %4)
	; </BoundMakeExpression>
	store %struct.class_SomeClass* %4, %struct.class_SomeClass** %VL_15
	; <BoundPackageCallExpression>
	; <BoundClassFieldAccessExpression>
	; <BoundVariableExpression>
	%8 = load %struct.class_SomeClass*, %struct.class_SomeClass** %VL_15
	; </BoundVariableExpression>
	%9 = bitcast %struct.class_SomeClass* %8 to i8*
	call void @exc_ThrowIfNull(i8* %9)
	%10 = getelementptr %struct.class_SomeClass, %struct.class_SomeClass* %8, i32 0, i32 2
	%11 = load %struct.class_String*, %struct.class_String** %10
	; </BoundClassFieldAccessExpression>
	%12 = bitcast %struct.class_String* %11 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %12)
	call void @sys_Print(%struct.class_String* %11)
	%13 = bitcast %struct.class_String* %11 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %13)
	; </BoundPackageCallExpression>
	; <ReturnARC>
	;  -> destroying reference to '%VL_15'
	%14 = load %struct.class_SomeClass*, %struct.class_SomeClass** %VL_15
	%15 = bitcast %struct.class_SomeClass* %14 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %15)
	; </ReturnARC>
	ret void
}
