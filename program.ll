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

@Any_vTable_Const = external global %struct.Any_vTable
@String_vTable_Const = external global %struct.String_vTable
@.str.0 = constant [3 x i8] c"%d\00"
@.str.1 = constant [3 x i8] c": \00"

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

declare void @Byte_public_Constructor(%struct.class_Byte* noundef %0, i8 noundef signext %1)

declare void @Byte_public_Die(i8* noundef %0)

declare i8 @Byte_public_GetValue(%struct.class_Byte* noundef %0)

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

declare void @Float_public_Constructor(%struct.class_Float* noundef %0, float noundef %1)

declare void @Float_public_Die(i8* noundef %0)

declare float @Float_public_GetValue(%struct.class_Float* noundef %0)

declare void @Int_public_Constructor(%struct.class_Int* noundef %0, i32 noundef %1)

declare void @Int_public_Die(i8* noundef %0)

declare i32 @Int_public_GetValue(%struct.class_Int* noundef %0)

declare void @Long_public_Constructor(%struct.class_Long* noundef %0, i64 noundef %1)

declare void @Long_public_Die(i8* noundef %0)

declare i64 @Long_public_GetValue(%struct.class_Long* noundef %0)

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

define void @main() {
0:
	%VL_15 = alloca i32
	%VL_16 = alloca i32
	%VL_17 = alloca i32
	br label %semiroot

semiroot:
	%1 = call i32 @sys_Now()
	store i32 %1, i32* %VL_15
	store i32 0, i32* %VL_16
	store i32 47, i32* %VL_17
	br label %Label1

Label2:
	%2 = load i32, i32* %VL_16
	%3 = getelementptr [3 x i8], [3 x i8]* @.str.0, i32 0, i32 0
	%4 = call i32 (i8*, i32, i8*, ...) @snprintf(i8* null, i32 0, i8* %3, i32 %2)
	%5 = add i32 %4, 1
	%6 = call i8* @malloc(i32 %5)
	%7 = add i32 %4, 1
	%8 = getelementptr [3 x i8], [3 x i8]* @.str.0, i32 0, i32 0
	%9 = call i32 (i8*, i32, i8*, ...) @snprintf(i8* %6, i32 %7, i8* %8, i32 %2)
	%10 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%11 = ptrtoint %struct.class_String* %10 to i32
	%12 = call i8* @malloc(i32 %11)
	%13 = bitcast i8* %12 to %struct.class_String*
	%14 = getelementptr %struct.class_String, %struct.class_String* %13, i32 0
	call void @String_public_Constructor(%struct.class_String* %14)
	%15 = bitcast %struct.class_String* %14 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %15)
	call void @String_public_Load(%struct.class_String* %14, i8* %6)
	call void @free(i8* %6)
	%16 = getelementptr [3 x i8], [3 x i8]* @.str.1, i32 0, i32 0
	%17 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%18 = ptrtoint %struct.class_String* %17 to i32
	%19 = call i8* @malloc(i32 %18)
	%20 = bitcast i8* %19 to %struct.class_String*
	%21 = getelementptr %struct.class_String, %struct.class_String* %20, i32 0
	call void @String_public_Constructor(%struct.class_String* %21)
	%22 = bitcast %struct.class_String* %21 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %22)
	call void @String_public_Load(%struct.class_String* %21, i8* %16)
	%23 = call %struct.class_String* @String_public_Concat(%struct.class_String* %14, %struct.class_String* %21)
	%24 = bitcast %struct.class_String* %14 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %24)
	%25 = bitcast %struct.class_String* %21 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %25)
	%26 = load i32, i32* %VL_16
	%27 = sext i32 %26 to i64
	%28 = call i64 @"F_fib_[T_long_[]]long"(i64 %27)
	%29 = getelementptr [3 x i8], [3 x i8]* @.str.0, i32 0, i32 0
	%30 = call i32 (i8*, i32, i8*, ...) @snprintf(i8* null, i32 0, i8* %29, i64 %28)
	%31 = add i32 %30, 1
	%32 = call i8* @malloc(i32 %31)
	%33 = add i32 %30, 1
	%34 = getelementptr [3 x i8], [3 x i8]* @.str.0, i32 0, i32 0
	%35 = call i32 (i8*, i32, i8*, ...) @snprintf(i8* %32, i32 %33, i8* %34, i64 %28)
	%36 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%37 = ptrtoint %struct.class_String* %36 to i32
	%38 = call i8* @malloc(i32 %37)
	%39 = bitcast i8* %38 to %struct.class_String*
	%40 = getelementptr %struct.class_String, %struct.class_String* %39, i32 0
	call void @String_public_Constructor(%struct.class_String* %40)
	%41 = bitcast %struct.class_String* %40 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %41)
	call void @String_public_Load(%struct.class_String* %40, i8* %32)
	call void @free(i8* %32)
	%42 = call %struct.class_String* @String_public_Concat(%struct.class_String* %23, %struct.class_String* %40)
	%43 = bitcast %struct.class_String* %23 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %43)
	%44 = bitcast %struct.class_String* %40 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %44)
	call void @sys_Print(%struct.class_String* %42)
	%45 = bitcast %struct.class_String* %42 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %45)
	br label %continue1

continue1:
	%46 = load i32, i32* %VL_16
	%47 = add i32 %46, 1
	store i32 %47, i32* %VL_16
	br label %Label1

Label1:
	%48 = load i32, i32* %VL_16
	%49 = load i32, i32* %VL_17
	%50 = icmp sle i32 %48, %49
	br i1 %50, label %Label2, label %break1

break1:
	; <GC - ARC>
	; </GC - ARC>
	; <ReturnARC>
	; </ReturnARC>
	ret void
}

define i64 @"F_fib_[T_long_[]]long"(i64 %VP_14) {
0:
	br label %semiroot

semiroot:
	%1 = sext i32 1 to i64
	%2 = icmp sle i64 %VP_14, %1
	br i1 %2, label %Label3, label %Label4

Label3:
	; <ReturnARC>
	; </ReturnARC>
	ret i64 %VP_14

Label4:
	%3 = sext i32 1 to i64
	%4 = sub i64 %VP_14, %3
	%5 = call i64 @"F_fib_[T_long_[]]long"(i64 %4)
	%6 = sext i32 2 to i64
	%7 = sub i64 %VP_14, %6
	%8 = call i64 @"F_fib_[T_long_[]]long"(i64 %7)
	%9 = add i64 %5, %8
	; <ReturnARC>
	; </ReturnARC>
	ret i64 %9
}
