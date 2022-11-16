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
%struct.Texture = type { i32, i32, i32, i32, i32 }
%struct.Color = type { i8, i8, i8, i8 }
%struct.Image = type { i8*, i32, i32, i32, i32 }

@Any_vTable_Const = external global %struct.Any_vTable
@Array_vTable_Const = external global %struct.String_vTable
@Byte_vTable_Const = external global %struct.String_vTable
@Float_vTable_Const = external global %struct.String_vTable
@Int_vTable_Const = external global %struct.String_vTable
@Long_vTable_Const = external global %struct.String_vTable
@String_vTable_Const = external global %struct.String_vTable
@Thread_vTable_Const = external global %struct.Any_vTable
@pArray_vTable_Const = external global %struct.String_vTable
@.str.0 = constant [12 x i8] c"cool window\00"

declare i8* @malloc(i32 %len)

declare void @free(i8* %dest)

declare i32 @printf(i8* %format, ...)

declare i32 @snprintf(i8* %dest, i32 %len, i8* %format, ...)

declare i32 @atoi(i8* %str)

declare i64 @atol(i8* %str)

declare double @atof(i8* %str)

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

declare void @Any_public_Constructor(%struct.class_Any* noundef %0)

declare void @Any_public_Die(i8* noundef %0)

declare void @Byte_public_Constructor(%struct.class_Byte* noundef %0, i8 noundef signext %1)

declare void @Byte_public_Die(i8* noundef %0)

declare i8 @Byte_public_GetValue(%struct.class_Byte* noundef %0)

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
	%VL_42 = alloca %struct.Image, align 1
	%VL_43 = alloca %struct.Texture, align 1
	%1 = alloca %struct.Color
	%2 = alloca %struct.Color
	%3 = alloca %struct.Color
	br label %semiroot

semiroot:
	%4 = getelementptr [12 x i8], [12 x i8]* @.str.0, i32 0, i32 0
	call void @InitWindow(i32 1000, i32 1000, i8* %4)
	%5 = getelementptr inbounds %struct.Color, %struct.Color* %1, i32 0, i32 0
	store i8 100, i8* %5
	%6 = getelementptr inbounds %struct.Color, %struct.Color* %1, i32 0, i32 1
	store i8 100, i8* %6
	%7 = getelementptr inbounds %struct.Color, %struct.Color* %1, i32 0, i32 2
	store i8 100, i8* %7
	%8 = getelementptr inbounds %struct.Color, %struct.Color* %1, i32 0, i32 3
	store i8 255, i8* %8
	%9 = call %struct.Image @GenImageColor$ADAPTED(i32 300, i32 300, %struct.Color* %1)
	%10 = load %struct.Image, %struct.Image %9
	store %struct.Image %10, %struct.Image* %VL_42
	%11 = call %struct.Texture @LoadTextureFromImage$ADAPTED(%struct.Image* %VL_42)
	%12 = load %struct.Texture, %struct.Texture %11
	store %struct.Texture %12, %struct.Texture* %VL_43
	br label %continue1

Label1:
	call void @BeginDrawing()
	%13 = getelementptr inbounds %struct.Color, %struct.Color* %2, i32 0, i32 0
	store i8 255, i8* %13
	%14 = getelementptr inbounds %struct.Color, %struct.Color* %2, i32 0, i32 1
	store i8 255, i8* %14
	%15 = getelementptr inbounds %struct.Color, %struct.Color* %2, i32 0, i32 2
	store i8 255, i8* %15
	%16 = getelementptr inbounds %struct.Color, %struct.Color* %2, i32 0, i32 3
	store i8 255, i8* %16
	call void @ClearBackground$ADAPTED(%struct.Color* %2)
	%17 = getelementptr inbounds %struct.Color, %struct.Color* %3, i32 0, i32 0
	store i8 255, i8* %17
	%18 = getelementptr inbounds %struct.Color, %struct.Color* %3, i32 0, i32 1
	store i8 255, i8* %18
	%19 = getelementptr inbounds %struct.Color, %struct.Color* %3, i32 0, i32 2
	store i8 255, i8* %19
	%20 = getelementptr inbounds %struct.Color, %struct.Color* %3, i32 0, i32 3
	store i8 255, i8* %20
	call void @DrawTexture$ADAPTED(%struct.Texture* %VL_43, i32 0, i32 0, %struct.Color* %3)
	call void @EndDrawing()
	br label %continue1

continue1:
	%21 = call i1 @WindowShouldClose()
	%22 = icmp ne i1 %21, 0
	%23 = xor i1 %22, true
	br i1 %23, label %Label1, label %break1

break1:
	call void @UnloadTexture$ADAPTED(%struct.Texture* %VL_43)
	call void @UnloadImage$ADAPTED(%struct.Image* %VL_42)
	call void @CloseWindow()
	; <ReturnARC>
	; </ReturnARC>
	ret void
}

declare void @CloseWindow()

declare %struct.Texture @LoadTextureFromImage$ADAPTED(%struct.Image* %VP_35)

declare void @ClearBackground$ADAPTED(%struct.Color* %VP_31)

declare %struct.Image @GenImageColor$ADAPTED(i32 %VP_32, i32 %VP_33, %struct.Color* %VP_34)

declare void @UnloadImage$ADAPTED(%struct.Image* %VP_41)

declare void @EndDrawing()

declare void @DrawTexture$ADAPTED(%struct.Texture* %VP_36, i32 %VP_37, i32 %VP_38, %struct.Color* %VP_39)

declare void @InitWindow(i32 %VP_28, i32 %VP_29, i8* %VP_30)

declare i1 @WindowShouldClose()

declare void @BeginDrawing()

declare void @UnloadTexture$ADAPTED(%struct.Texture* %VP_40)
