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
%struct.SomeStruct = type { i1, i32 }

@Any_vTable_Const = external global %struct.Any_vTable
@Array_vTable_Const = external global %struct.String_vTable
@Byte_vTable_Const = external global %struct.String_vTable
@Float_vTable_Const = external global %struct.String_vTable
@Int_vTable_Const = external global %struct.String_vTable
@Long_vTable_Const = external global %struct.String_vTable
@String_vTable_Const = external global %struct.String_vTable
@Thread_vTable_Const = external global %struct.Any_vTable
@pArray_vTable_Const = external global %struct.String_vTable
@.str.0 = constant [5 x i8] c"true\00"
@.str.1 = constant [6 x i8] c"false\00"

declare i8* @malloc(i32 %len)

declare void @free(i8* %dest)

declare i32 @snprintf(i8* %dest, i32 %len, i8* %format, ...)

declare i32 @atoi(i8* %str)

declare i64 @atol(i8* %str)

declare double @atof(i8* %str)

declare void @Array_public_Constructor(%struct.class_Array* noundef %0, i32 noundef %1)

declare void @Array_public_Die(i8* noundef %0)

declare %struct.class_Any* @Array_public_GetElement(%struct.class_Array* noundef %0, i32 noundef %1)

declare void @Array_public_SetElement(%struct.class_Array* noundef %0, i32 noundef %1, %struct.class_Any* noundef %2)

declare i32 @Array_public_GetLength(%struct.class_Array* noundef %0)

declare void @Array_public_Push(%struct.class_Array* noundef %0, %struct.class_Any* noundef %1)

declare void @Byte_public_Constructor(%struct.class_Byte* noundef %0, i8 noundef signext %1)

declare void @Byte_public_Die(i8* noundef %0)

declare i8 @Byte_public_GetValue(%struct.class_Byte* noundef %0)

declare void @Int_public_Constructor(%struct.class_Int* noundef %0, i32 noundef %1)

declare void @Int_public_Die(i8* noundef %0)

declare i32 @Int_public_GetValue(%struct.class_Int* noundef %0)

declare void @Thread_public_Constructor(%struct.class_Thread* noundef %0, i8* (i8*)* noundef %1, i8* noundef %2)

declare void @Thread_public_Die(i8* noundef %0)

declare void @Thread_public_Start(%struct.class_Thread* noundef %0)

declare void @Thread_public_Join(%struct.class_Thread* noundef %0)

declare void @Thread_public_Kill(%struct.class_Thread* noundef %0)

declare void @Any_public_Constructor(%struct.class_Any* noundef %0)

declare void @Any_public_Die(i8* noundef %0)

declare void @Float_public_Constructor(%struct.class_Float* noundef %0, float noundef %1)

declare void @Float_public_Die(i8* noundef %0)

declare float @Float_public_GetValue(%struct.class_Float* noundef %0)

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

define void @main() {
0:
	%VL_17 = alloca %struct.SomeStruct
	br label %semiroot

semiroot:
	store %struct.SomeStruct { i1 false, i32 0 }, %struct.SomeStruct* %VL_17
	%1 = alloca %struct.SomeStruct
	%2 = load %struct.SomeStruct, %struct.SomeStruct* %VL_17
	%3 = call %struct.SomeStruct @"F_someOtherFunc_[T_SomeStruct_[]]SomeStruct"(%struct.SomeStruct %2)
	store %struct.SomeStruct %3, %struct.SomeStruct* %1
	%4 = getelementptr %struct.SomeStruct, %struct.SomeStruct* %1, i32 0, i32 0
	%5 = load i1, i1* %4
	%6 = getelementptr [5 x i8], [5 x i8]* @.str.0, i32 0, i32 0
	%7 = getelementptr [6 x i8], [6 x i8]* @.str.1, i32 0, i32 0
	%8 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%9 = ptrtoint %struct.class_String* %8 to i32
	%10 = call i8* @malloc(i32 %9)
	%11 = bitcast i8* %10 to %struct.class_String*
	%12 = getelementptr %struct.class_String, %struct.class_String* %11, i32 0
	call void @String_public_Constructor(%struct.class_String* %12)
	%13 = bitcast %struct.class_String* %12 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %13)
	%14 = select i1 %5, i8* %6, i8* %7
	call void @String_public_Load(%struct.class_String* %12, i8* %14)
	call void @sys_Print(%struct.class_String* %12)
	%15 = bitcast %struct.class_String* %12 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %15)
	; <ReturnARC>
	; </ReturnARC>
	ret void
}

define %struct.SomeStruct @F_someFunc_SomeStruct() {
0:
	%VL_18 = alloca %struct.SomeStruct
	br label %semiroot

semiroot:
	store %struct.SomeStruct { i1 false, i32 0 }, %struct.SomeStruct* %VL_18
	%1 = getelementptr %struct.SomeStruct, %struct.SomeStruct* %VL_18, i32 0, i32 0
	store i1 false, i1* %1
	%2 = getelementptr %struct.SomeStruct, %struct.SomeStruct* %VL_18, i32 0, i32 1
	store i32 100, i32* %2
	%3 = load %struct.SomeStruct, %struct.SomeStruct* %VL_18
	; <ReturnARC>
	; </ReturnARC>
	ret %struct.SomeStruct %3
}

define %struct.SomeStruct @"F_someOtherFunc_[T_SomeStruct_[]]SomeStruct"(%struct.SomeStruct %VP_16) {
0:
	%LVP_16 = alloca %struct.SomeStruct
	store %struct.SomeStruct %VP_16, %struct.SomeStruct* %LVP_16
	br label %semiroot

semiroot:
	%1 = alloca %struct.SomeStruct
	%2 = load %struct.SomeStruct, %struct.SomeStruct* %LVP_16
	store %struct.SomeStruct %2, %struct.SomeStruct* %1
	%3 = getelementptr %struct.SomeStruct, %struct.SomeStruct* %1, i32 0, i32 0
	store i1 true, i1* %3
	%4 = load %struct.SomeStruct, %struct.SomeStruct* %LVP_16
	; <ReturnARC>
	; </ReturnARC>
	ret %struct.SomeStruct %4
}
