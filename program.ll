%struct.Standard_vTable = type { i8*, i8*, i8* }
%struct.class_Any = type { %struct.Standard_vTable }
%struct.class_Array = type { %struct.Standard_vTable, %struct.class_Any**, i32, i32, i32 }
%struct.class_Byte = type { %struct.Standard_vTable, i8 }
%struct.class_Float = type { %struct.Standard_vTable, float }
%struct.class_Int = type { %struct.Standard_vTable, i32 }
%struct.class_Long = type { %struct.Standard_vTable, i64 }
%struct.class_String = type { %struct.Standard_vTable, i8*, i32, i32, i32 }
%struct.class_Thread = type { %struct.Standard_vTable, i8* (i8*)*, %struct.class_Array*, i64 }
%struct.class_pArray = type { %struct.Standard_vTable, i8*, i32, i32, i32, i32 }

@Any_vTable_Const = external global %struct.Standard_vTable
@Array_vTable_Const = external global %struct.Standard_vTable
@Byte_vTable_Const = external global %struct.Standard_vTable
@Float_vTable_Const = external global %struct.Standard_vTable
@Int_vTable_Const = external global %struct.Standard_vTable
@Long_vTable_Const = external global %struct.Standard_vTable
@String_vTable_Const = external global %struct.Standard_vTable
@Thread_vTable_Const = external global %struct.Standard_vTable
@pArray_vTable_Const = external global %struct.Standard_vTable
@.str.0 = constant [17 x i8] c"Got the string: \00"
@.str.c.1 = constant [7 x i8] c"String\00"
@.str.c.2 = constant [13 x i8] c"TO_string_[]\00"
@.str.3 = constant [10 x i8] c"very cool\00"
@.str.c.4 = constant [7 x i8] c"String\00"
@.str.c.5 = constant [13 x i8] c"TO_string_[]\00"
@.str.c.6 = constant [6 x i8] c"Array\00"
@.str.c.7 = constant [16 x i8] c"TO_array_[any;]\00"
@.str.8 = constant [9 x i8] c"among us\00"
@.str.c.9 = constant [7 x i8] c"String\00"
@.str.c.10 = constant [13 x i8] c"TO_string_[]\00"
@.str.c.11 = constant [5 x i8] c"Long\00"
@.str.c.12 = constant [10 x i8] c"T_long_[]\00"
@.str.c.13 = constant [13 x i8] c"TO_string_[]\00"
@.str.c.14 = constant [10 x i8] c"T_long_[]\00"
@.str.c.15 = constant [7 x i8] c"Thread\00"
@.str.c.16 = constant [13 x i8] c"TO_thread_[]\00"
@.str.17 = constant [15 x i8] c"we threadin :)\00"
@.str.c.18 = constant [7 x i8] c"String\00"
@.str.c.19 = constant [13 x i8] c"TO_string_[]\00"
@.str.20 = constant [18 x i8] c"no more thread :(\00"
@.str.c.21 = constant [7 x i8] c"String\00"
@.str.c.22 = constant [13 x i8] c"TO_string_[]\00"

declare i8* @malloc(i32 %len)

declare void @free(i8* %dest)

declare i32 @printf(i8* %format, ...)

declare i32 @snprintf(i8* %dest, i32 %len, i8* %format, ...)

declare i32 @atoi(i8* %str)

declare i64 @atol(i8* %str)

declare double @atof(i8* %str)

declare i8* @GC_init()

declare i8* @GC_malloc(i32 %len)

declare i8* @GC_realloc(i8* %ptr, i32 %len)

declare void @Array_public_Constructor(%struct.class_Array* noundef %0, i32 noundef %1)

declare %struct.class_Any* @Array_public_GetElement(%struct.class_Array* noundef %0, i32 noundef %1)

declare void @Array_public_SetElement(%struct.class_Array* noundef %0, i32 noundef %1, %struct.class_Any* noundef %2)

declare i32 @Array_public_GetLength(%struct.class_Array* noundef %0)

declare void @Array_public_Push(%struct.class_Array* noundef %0, %struct.class_Any* noundef %1)

declare void @Int_public_Constructor(%struct.class_Int* noundef %0, i32 noundef %1)

declare i32 @Int_public_GetValue(%struct.class_Int* noundef %0)

declare void @Long_public_Constructor(%struct.class_Long* noundef %0, i64 noundef %1)

declare i64 @Long_public_GetValue(%struct.class_Long* noundef %0)

declare void @String_public_Constructor(%struct.class_String* noundef %0)

declare void @String_public_Load(%struct.class_String* noundef %0, i8* noundef %1)

declare void @String_public_Resize(%struct.class_String* noundef %0, i32 noundef %1)

declare void @String_public_AddChar(%struct.class_String* noundef %0, i8 noundef signext %1)

declare %struct.class_String* @String_public_Concat(%struct.class_String* noundef %0, %struct.class_String* noundef %1)

declare i1 @String_public_Equal(%struct.class_String* noundef %0, %struct.class_String* noundef %1)

declare i8* @String_public_GetBuffer(%struct.class_String* noundef %0)

declare i32 @String_public_GetLength(%struct.class_String* noundef %0)

declare %struct.class_String* @String_public_Substring(%struct.class_String* noundef %0, i32 noundef %1, i32 noundef %2)

declare void @Thread_public_Constructor(%struct.class_Thread* noundef %0, i8* (i8*)* noundef %1, %struct.class_Array* noundef %2)

declare void @Thread_public_Start(%struct.class_Thread* noundef %0)

declare void @Thread_public_Join(%struct.class_Thread* noundef %0)

declare void @Thread_public_Kill(%struct.class_Thread* noundef %0)

declare void @pArray_public_Constructor(%struct.class_pArray* noundef %0, i32 noundef %1, i32 noundef %2)

declare i32 @pArray_public_GetLength(%struct.class_pArray* noundef %0)

declare i8* @pArray_public_Grow(%struct.class_pArray* noundef %0)

declare i8* @pArray_public_GetElementPtr(%struct.class_pArray* noundef %0, i32 noundef %1)

declare void @Any_public_Constructor(%struct.class_Any* noundef %0)

declare void @Byte_public_Constructor(%struct.class_Byte* noundef %0, i8 noundef signext %1)

declare i8 @Byte_public_GetValue(%struct.class_Byte* noundef %0)

declare void @Float_public_Constructor(%struct.class_Float* noundef %0, float noundef %1)

declare float @Float_public_GetValue(%struct.class_Float* noundef %0)

declare void @exc_Throw(i8* noundef %0)

declare void @exc_ThrowIfNull(i8* noundef %0)

declare void @exc_ThrowIfInvalidCast(%struct.class_Any* noundef %0, %struct.Standard_vTable* noundef %1, i8* noundef %2)

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
	%VL_16 = alloca %struct.class_Thread*
	br label %semiroot

semiroot:
	%1 = call i8* @GC_init()
	%2 = getelementptr %struct.class_Array, %struct.class_Array* null, i32 1
	%3 = ptrtoint %struct.class_Array* %2 to i32
	%4 = call i8* @GC_malloc(i32 %3)
	%5 = bitcast i8* %4 to %struct.class_Array*
	%6 = getelementptr %struct.class_Array, %struct.class_Array* %5, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([6 x i8], [6 x i8]* @.str.c.6, i32 0, i32 0), i8* getelementptr ([16 x i8], [16 x i8]* @.str.c.7, i32 0, i32 0) }, %struct.Standard_vTable* %6
	call void @Array_public_Constructor(%struct.class_Array* %5, i32 2)
	%7 = getelementptr [9 x i8], [9 x i8]* @.str.8, i32 0, i32 0
	%8 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%9 = ptrtoint %struct.class_String* %8 to i32
	%10 = call i8* @GC_malloc(i32 %9)
	%11 = bitcast i8* %10 to %struct.class_String*
	%12 = getelementptr %struct.class_String, %struct.class_String* %11, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([7 x i8], [7 x i8]* @.str.c.9, i32 0, i32 0), i8* getelementptr ([13 x i8], [13 x i8]* @.str.c.10, i32 0, i32 0) }, %struct.Standard_vTable* %12
	call void @String_public_Constructor(%struct.class_String* %11)
	call void @String_public_Load(%struct.class_String* %11, i8* %7)
	%13 = bitcast %struct.class_String* %11 to %struct.class_Any*
	call void @Array_public_SetElement(%struct.class_Array* %5, i32 0, %struct.class_Any* %13)
	%14 = ptrtoint void (%struct.class_String*)* @"F_LAMBDA_1_[TO_string_[]]void" to i64
	%15 = getelementptr %struct.class_Long, %struct.class_Long* null, i32 1
	%16 = ptrtoint %struct.class_Long* %15 to i32
	%17 = call i8* @GC_malloc(i32 %16)
	%18 = bitcast i8* %17 to %struct.class_Long*
	%19 = getelementptr %struct.class_Long, %struct.class_Long* %18, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([5 x i8], [5 x i8]* @.str.c.11, i32 0, i32 0), i8* getelementptr ([10 x i8], [10 x i8]* @.str.c.12, i32 0, i32 0) }, %struct.Standard_vTable* %19
	call void @Long_public_Constructor(%struct.class_Long* %18, i64 %14)
	%20 = bitcast %struct.class_Long* %18 to %struct.class_Any*
	call void @Array_public_SetElement(%struct.class_Array* %5, i32 1, %struct.class_Any* %20)
	%21 = getelementptr %struct.class_Thread, %struct.class_Thread* null, i32 1
	%22 = ptrtoint %struct.class_Thread* %21 to i32
	%23 = call i8* @GC_malloc(i32 %22)
	%24 = bitcast i8* %23 to %struct.class_Thread*
	%25 = getelementptr %struct.class_Thread, %struct.class_Thread* %24, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([7 x i8], [7 x i8]* @.str.c.15, i32 0, i32 0), i8* getelementptr ([13 x i8], [13 x i8]* @.str.c.16, i32 0, i32 0) }, %struct.Standard_vTable* %25
	call void @Thread_public_Constructor(%struct.class_Thread* %24, i8* (i8*)* @"T_action_[string;void;]_ThreadWrapper", %struct.class_Array* %5)
	call void @Thread_public_Start(%struct.class_Thread* %24)
	store %struct.class_Thread* %24, %struct.class_Thread** %VL_16
	%26 = getelementptr [15 x i8], [15 x i8]* @.str.17, i32 0, i32 0
	%27 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%28 = ptrtoint %struct.class_String* %27 to i32
	%29 = call i8* @GC_malloc(i32 %28)
	%30 = bitcast i8* %29 to %struct.class_String*
	%31 = getelementptr %struct.class_String, %struct.class_String* %30, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([7 x i8], [7 x i8]* @.str.c.18, i32 0, i32 0), i8* getelementptr ([13 x i8], [13 x i8]* @.str.c.19, i32 0, i32 0) }, %struct.Standard_vTable* %31
	call void @String_public_Constructor(%struct.class_String* %30)
	call void @String_public_Load(%struct.class_String* %30, i8* %26)
	call void @sys_Print(%struct.class_String* %30)
	%32 = load %struct.class_Thread*, %struct.class_Thread** %VL_16
	%33 = bitcast %struct.class_Thread* %32 to i8*
	call void @exc_ThrowIfNull(i8* %33)
	call void @Thread_public_Join(%struct.class_Thread* %32)
	%34 = getelementptr [18 x i8], [18 x i8]* @.str.20, i32 0, i32 0
	%35 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%36 = ptrtoint %struct.class_String* %35 to i32
	%37 = call i8* @GC_malloc(i32 %36)
	%38 = bitcast i8* %37 to %struct.class_String*
	%39 = getelementptr %struct.class_String, %struct.class_String* %38, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([7 x i8], [7 x i8]* @.str.c.21, i32 0, i32 0), i8* getelementptr ([13 x i8], [13 x i8]* @.str.c.22, i32 0, i32 0) }, %struct.Standard_vTable* %39
	call void @String_public_Constructor(%struct.class_String* %38)
	call void @String_public_Load(%struct.class_String* %38, i8* %34)
	call void @sys_Print(%struct.class_String* %38)
	ret void
}

define void @"F_LAMBDA_1_[TO_string_[]]void"(%struct.class_String* %VP_14) {
0:
	%LVP_14 = alloca %struct.class_String*
	store %struct.class_String* %VP_14, %struct.class_String** %LVP_14
	br label %semiroot

semiroot:
	%1 = getelementptr [17 x i8], [17 x i8]* @.str.0, i32 0, i32 0
	%2 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%3 = ptrtoint %struct.class_String* %2 to i32
	%4 = call i8* @GC_malloc(i32 %3)
	%5 = bitcast i8* %4 to %struct.class_String*
	%6 = getelementptr %struct.class_String, %struct.class_String* %5, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([7 x i8], [7 x i8]* @.str.c.1, i32 0, i32 0), i8* getelementptr ([13 x i8], [13 x i8]* @.str.c.2, i32 0, i32 0) }, %struct.Standard_vTable* %6
	call void @String_public_Constructor(%struct.class_String* %5)
	call void @String_public_Load(%struct.class_String* %5, i8* %1)
	%7 = load %struct.class_String*, %struct.class_String** %LVP_14
	%8 = call %struct.class_String* @String_public_Concat(%struct.class_String* %5, %struct.class_String* %7)
	call void @sys_Print(%struct.class_String* %8)
	%9 = getelementptr [10 x i8], [10 x i8]* @.str.3, i32 0, i32 0
	%10 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%11 = ptrtoint %struct.class_String* %10 to i32
	%12 = call i8* @GC_malloc(i32 %11)
	%13 = bitcast i8* %12 to %struct.class_String*
	%14 = getelementptr %struct.class_String, %struct.class_String* %13, i32 0, i32 0
	store %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr ([7 x i8], [7 x i8]* @.str.c.4, i32 0, i32 0), i8* getelementptr ([13 x i8], [13 x i8]* @.str.c.5, i32 0, i32 0) }, %struct.Standard_vTable* %14
	call void @String_public_Constructor(%struct.class_String* %13)
	call void @String_public_Load(%struct.class_String* %13, i8* %9)
	call void @sys_Print(%struct.class_String* %13)
	ret void
}

define i8* @"T_action_[string;void;]_ThreadWrapper"(i8* %param) {
0:
	%1 = bitcast i8* %param to %struct.class_Array*
	%2 = call %struct.class_Any* @Array_public_GetElement(%struct.class_Array* %1, i32 0)
	%3 = bitcast %struct.class_Any* %2 to %struct.class_Any*
	%4 = bitcast %struct.Standard_vTable* @String_vTable_Const to %struct.Standard_vTable*
	call void @exc_ThrowIfInvalidCast(%struct.class_Any* %3, %struct.Standard_vTable* %4, i8* getelementptr ([13 x i8], [13 x i8]* @.str.c.13, i32 0, i32 0))
	%5 = bitcast %struct.class_Any* %2 to %struct.class_String*
	%6 = call %struct.class_Any* @Array_public_GetElement(%struct.class_Array* %1, i32 1)
	%7 = bitcast %struct.class_Any* %6 to %struct.class_Any*
	%8 = bitcast %struct.Standard_vTable* @Long_vTable_Const to %struct.Standard_vTable*
	call void @exc_ThrowIfInvalidCast(%struct.class_Any* %7, %struct.Standard_vTable* %8, i8* getelementptr ([10 x i8], [10 x i8]* @.str.c.14, i32 0, i32 0))
	%9 = bitcast %struct.class_Any* %6 to %struct.class_Long*
	%10 = call i64 @Long_public_GetValue(%struct.class_Long* %9)
	%11 = inttoptr i64 %10 to void (%struct.class_String*)*
	call void %11(%struct.class_String* %5)
	ret i8* null
}
