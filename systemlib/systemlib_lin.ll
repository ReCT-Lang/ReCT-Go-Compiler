; ModuleID = './systemlib_lin.bc'
source_filename = "llvm-link"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }
%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.class_Array = type { %struct.String_vTable*, i32, %struct.class_Any**, i32, i32, i32 }
%struct.class_pArray = type { %struct.String_vTable*, i32, i8*, i32, i32, i32, i32 }
%struct.class_Int = type { %struct.String_vTable*, i32, i32 }
%struct.class_Float = type { %struct.String_vTable*, i32, float }
%struct.class_Bool = type { %struct.String_vTable*, i32, i8 }
%struct.class_Action = type { %struct.Any_vTable*, i32, i8* (i8*)*, i8*, i64 }
%union.pthread_attr_t = type { i64, [48 x i8] }

@.str = private unnamed_addr constant [59 x i8] c"\1B[36mARC \1B[0m- \1B[32mRegistered %s reference [%d] - %s\1B[0m\0A\00", align 1
@.str.1 = private unnamed_addr constant [61 x i8] c"\1B[36mARC \1B[0m- \1B[33mUnregistered %s reference [%d] - %s\1B[0m\0A\00", align 1
@.str.2 = private unnamed_addr constant [53 x i8] c"\1B[36mARC \1B[0m- \1B[31mDestroying %s instance - %s\1B[0m\0A\00", align 1
@.str.3 = private unnamed_addr constant [44 x i8] c"\1B[36mARC \1B[0m- \1B[0;35mWhat?? [%d] - %s\1B[0m\0A\00", align 1
@.str.4 = private unnamed_addr constant [4 x i8] c"Any\00", align 1
@Any_vTable_Const = dso_local constant %struct.Any_vTable { i8* null, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.4, i32 0, i32 0), void (i8*)* @Any_public_Die }, align 8
@.str.1.5 = private unnamed_addr constant [7 x i8] c"String\00", align 1
@String_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.1.5, i32 0, i32 0), void (i8*)* @String_public_Die }, align 8
@.str.2.8 = private unnamed_addr constant [44 x i8] c"https://www.youtube.com/watch?v=dQw4w9WgXcQ\00", align 1
@.str.3.6 = private unnamed_addr constant [4 x i8] c"Int\00", align 1
@Int_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.3.6, i32 0, i32 0), void (i8*)* @Int_public_Die }, align 8
@.str.4.7 = private unnamed_addr constant [6 x i8] c"Float\00", align 1
@Float_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.4.7, i32 0, i32 0), void (i8*)* @Float_public_Die }, align 8
@.str.5 = private unnamed_addr constant [5 x i8] c"Bool\00", align 1
@Bool_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.str.5, i32 0, i32 0), void (i8*)* @Bool_public_Die }, align 8
@.str.6 = private unnamed_addr constant [6 x i8] c"Array\00", align 1
@Array_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.6, i32 0, i32 0), void (i8*)* @Array_public_Die }, align 8
@.str.7 = private unnamed_addr constant [7 x i8] c"pArray\00", align 1
@pArray_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.7, i32 0, i32 0), void (i8*)* @pArray_public_Die }, align 8
@.str.8 = private unnamed_addr constant [7 x i8] c"Action\00", align 1
@Action_vTable_Const = dso_local constant %struct.Any_vTable { i8* bitcast (%struct.Any_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.8, i32 0, i32 0), void (i8*)* @Action_public_Die }, align 8
@isCursorVisible = dso_local global i8 1, align 1
@.str.9 = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.1.10 = private unnamed_addr constant [3 x i8] c"%s\00", align 1
@.str.2.11 = private unnamed_addr constant [8 x i8] c"\1B[2J\1B[H\00", align 1
@.str.3.12 = private unnamed_addr constant [10 x i8] c"%c[%d;%df\00", align 1
@.str.4.13 = private unnamed_addr constant [8 x i8] c"\1B[?251]\00", align 1

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @arc_RegisterReference(%struct.class_Any* %0) #0 {
  %2 = alloca %struct.class_Any*, align 8
  store %struct.class_Any* %0, %struct.class_Any** %2, align 8
  %3 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %4 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %3, i32 0, i32 1
  %5 = load i32, i32* %4, align 8
  %6 = add nsw i32 %5, 1
  store i32 %6, i32* %4, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @arc_UnregisterReference(%struct.class_Any* %0) #0 {
  %2 = alloca %struct.class_Any*, align 8
  store %struct.class_Any* %0, %struct.class_Any** %2, align 8
  %3 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %4 = icmp eq %struct.class_Any* %3, null
  br i1 %4, label %5, label %6

5:                                                ; preds = %1
  br label %25

6:                                                ; preds = %1
  %7 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %8 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %7, i32 0, i32 1
  %9 = load i32, i32* %8, align 8
  %10 = add nsw i32 %9, -1
  store i32 %10, i32* %8, align 8
  %11 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %12 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %11, i32 0, i32 1
  %13 = load i32, i32* %12, align 8
  %14 = icmp sle i32 %13, 0
  br i1 %14, label %15, label %25

15:                                               ; preds = %6
  %16 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %17 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %16, i32 0, i32 0
  %18 = load %struct.Any_vTable*, %struct.Any_vTable** %17, align 8
  %19 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %18, i32 0, i32 2
  %20 = load void (i8*)*, void (i8*)** %19, align 8
  %21 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %22 = bitcast %struct.class_Any* %21 to i8*
  call void %20(i8* %22)
  %23 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %24 = bitcast %struct.class_Any* %23 to i8*
  call void @free(i8* %24) #7
  br label %25

25:                                               ; preds = %15, %6, %5
  ret void
}

; Function Attrs: nounwind
declare dso_local void @free(i8*) #1

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @arc_RegisterReferenceVerbose(%struct.class_Any* %0, i8* %1) #0 {
  %3 = alloca %struct.class_Any*, align 8
  %4 = alloca i8*, align 8
  store %struct.class_Any* %0, %struct.class_Any** %3, align 8
  store i8* %1, i8** %4, align 8
  %5 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %6 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %5, i32 0, i32 1
  %7 = load i32, i32* %6, align 8
  %8 = add nsw i32 %7, 1
  store i32 %8, i32* %6, align 8
  %9 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %10 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %9, i32 0, i32 0
  %11 = load %struct.Any_vTable*, %struct.Any_vTable** %10, align 8
  %12 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %11, i32 0, i32 1
  %13 = load i8*, i8** %12, align 8
  %14 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %15 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %14, i32 0, i32 1
  %16 = load i32, i32* %15, align 8
  %17 = load i8*, i8** %4, align 8
  %18 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([59 x i8], [59 x i8]* @.str, i64 0, i64 0), i8* %13, i32 %16, i8* %17)
  ret void
}

declare dso_local i32 @printf(i8*, ...) #2

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @arc_UnregisterReferenceVerbose(%struct.class_Any* %0, i8* %1) #0 {
  %3 = alloca %struct.class_Any*, align 8
  %4 = alloca i8*, align 8
  store %struct.class_Any* %0, %struct.class_Any** %3, align 8
  store i8* %1, i8** %4, align 8
  %5 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %6 = icmp eq %struct.class_Any* %5, null
  br i1 %6, label %7, label %8

7:                                                ; preds = %2
  br label %56

8:                                                ; preds = %2
  %9 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %10 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %9, i32 0, i32 1
  %11 = load i32, i32* %10, align 8
  %12 = add nsw i32 %11, -1
  store i32 %12, i32* %10, align 8
  %13 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %14 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %13, i32 0, i32 0
  %15 = load %struct.Any_vTable*, %struct.Any_vTable** %14, align 8
  %16 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %15, i32 0, i32 1
  %17 = load i8*, i8** %16, align 8
  %18 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %19 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %18, i32 0, i32 1
  %20 = load i32, i32* %19, align 8
  %21 = load i8*, i8** %4, align 8
  %22 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([61 x i8], [61 x i8]* @.str.1, i64 0, i64 0), i8* %17, i32 %20, i8* %21)
  %23 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %24 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %23, i32 0, i32 1
  %25 = load i32, i32* %24, align 8
  %26 = icmp eq i32 %25, 0
  br i1 %26, label %27, label %44

27:                                               ; preds = %8
  %28 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %29 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %28, i32 0, i32 0
  %30 = load %struct.Any_vTable*, %struct.Any_vTable** %29, align 8
  %31 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %30, i32 0, i32 1
  %32 = load i8*, i8** %31, align 8
  %33 = load i8*, i8** %4, align 8
  %34 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([53 x i8], [53 x i8]* @.str.2, i64 0, i64 0), i8* %32, i8* %33)
  %35 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %36 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %35, i32 0, i32 0
  %37 = load %struct.Any_vTable*, %struct.Any_vTable** %36, align 8
  %38 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %37, i32 0, i32 2
  %39 = load void (i8*)*, void (i8*)** %38, align 8
  %40 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %41 = bitcast %struct.class_Any* %40 to i8*
  call void %39(i8* %41)
  %42 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %43 = bitcast %struct.class_Any* %42 to i8*
  call void @free(i8* %43) #7
  br label %56

44:                                               ; preds = %8
  %45 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %46 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %45, i32 0, i32 1
  %47 = load i32, i32* %46, align 8
  %48 = icmp slt i32 %47, 0
  br i1 %48, label %49, label %55

49:                                               ; preds = %44
  %50 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %51 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %50, i32 0, i32 1
  %52 = load i32, i32* %51, align 8
  %53 = load i8*, i8** %4, align 8
  %54 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([44 x i8], [44 x i8]* @.str.3, i64 0, i64 0), i32 %52, i8* %53)
  br label %55

55:                                               ; preds = %49, %44
  br label %56

56:                                               ; preds = %55, %27, %7
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Any_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @String_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  %3 = alloca %struct.class_String*, align 8
  store i8* %0, i8** %2, align 8
  %4 = load i8*, i8** %2, align 8
  %5 = bitcast i8* %4 to %struct.class_String*
  store %struct.class_String* %5, %struct.class_String** %3, align 8
  %6 = load %struct.class_String*, %struct.class_String** %3, align 8
  %7 = getelementptr inbounds %struct.class_String, %struct.class_String* %6, i32 0, i32 2
  %8 = load i8*, i8** %7, align 8
  %9 = icmp ne i8* %8, null
  br i1 %9, label %10, label %14

10:                                               ; preds = %1
  %11 = load %struct.class_String*, %struct.class_String** %3, align 8
  %12 = getelementptr inbounds %struct.class_String, %struct.class_String* %11, i32 0, i32 2
  %13 = load i8*, i8** %12, align 8
  call void @free(i8* %13) #7
  br label %14

14:                                               ; preds = %10, %1
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Int_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Float_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Bool_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Array_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  %3 = alloca %struct.class_Array*, align 8
  %4 = alloca i32, align 4
  store i8* %0, i8** %2, align 8
  %5 = load i8*, i8** %2, align 8
  %6 = bitcast i8* %5 to %struct.class_Array*
  store %struct.class_Array* %6, %struct.class_Array** %3, align 8
  store i32 0, i32* %4, align 4
  br label %7

7:                                                ; preds = %21, %1
  %8 = load i32, i32* %4, align 4
  %9 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %10 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %9, i32 0, i32 3
  %11 = load i32, i32* %10, align 8
  %12 = icmp slt i32 %8, %11
  br i1 %12, label %13, label %24

13:                                               ; preds = %7
  %14 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %15 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %14, i32 0, i32 2
  %16 = load %struct.class_Any**, %struct.class_Any*** %15, align 8
  %17 = load i32, i32* %4, align 4
  %18 = sext i32 %17 to i64
  %19 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %16, i64 %18
  %20 = load %struct.class_Any*, %struct.class_Any** %19, align 8
  call void @arc_UnregisterReference(%struct.class_Any* %20)
  br label %21

21:                                               ; preds = %13
  %22 = load i32, i32* %4, align 4
  %23 = add nsw i32 %22, 1
  store i32 %23, i32* %4, align 4
  br label %7, !llvm.loop !4

24:                                               ; preds = %7
  %25 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %26 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %25, i32 0, i32 2
  %27 = load %struct.class_Any**, %struct.class_Any*** %26, align 8
  %28 = bitcast %struct.class_Any** %27 to i8*
  call void @free(i8* %28) #7
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @pArray_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  %3 = alloca %struct.class_pArray*, align 8
  store i8* %0, i8** %2, align 8
  %4 = load i8*, i8** %2, align 8
  %5 = bitcast i8* %4 to %struct.class_pArray*
  store %struct.class_pArray* %5, %struct.class_pArray** %3, align 8
  %6 = load %struct.class_pArray*, %struct.class_pArray** %3, align 8
  %7 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %6, i32 0, i32 2
  %8 = load i8*, i8** %7, align 8
  call void @free(i8* %8) #7
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Action_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Any_public_Constructor(%struct.class_Any* %0) #0 {
  %2 = alloca %struct.class_Any*, align 8
  store %struct.class_Any* %0, %struct.class_Any** %2, align 8
  %3 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %4 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %3, i32 0, i32 0
  store %struct.Any_vTable* @Any_vTable_Const, %struct.Any_vTable** %4, align 8
  %5 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %6 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %5, i32 0, i32 1
  store i32 0, i32* %6, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @String_public_Constructor(%struct.class_String* %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = getelementptr inbounds %struct.class_String, %struct.class_String* %3, i32 0, i32 0
  store %struct.String_vTable* @String_vTable_Const, %struct.String_vTable** %4, align 8
  %5 = load %struct.class_String*, %struct.class_String** %2, align 8
  %6 = getelementptr inbounds %struct.class_String, %struct.class_String* %5, i32 0, i32 1
  store i32 0, i32* %6, align 8
  %7 = load %struct.class_String*, %struct.class_String** %2, align 8
  %8 = getelementptr inbounds %struct.class_String, %struct.class_String* %7, i32 0, i32 2
  store i8* null, i8** %8, align 8
  %9 = load %struct.class_String*, %struct.class_String** %2, align 8
  %10 = getelementptr inbounds %struct.class_String, %struct.class_String* %9, i32 0, i32 3
  store i32 0, i32* %10, align 8
  %11 = load %struct.class_String*, %struct.class_String** %2, align 8
  %12 = getelementptr inbounds %struct.class_String, %struct.class_String* %11, i32 0, i32 4
  store i32 0, i32* %12, align 4
  %13 = load %struct.class_String*, %struct.class_String** %2, align 8
  %14 = getelementptr inbounds %struct.class_String, %struct.class_String* %13, i32 0, i32 5
  store i32 16, i32* %14, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @String_public_Load(%struct.class_String* %0, i8* %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i8*, align 8
  %5 = alloca i32, align 4
  %6 = alloca i8*, align 8
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i8* %1, i8** %4, align 8
  %7 = load i8*, i8** %4, align 8
  %8 = call i64 @strlen(i8* %7) #8
  %9 = trunc i64 %8 to i32
  store i32 %9, i32* %5, align 4
  %10 = load i32, i32* %5, align 4
  %11 = add nsw i32 %10, 1
  %12 = sext i32 %11 to i64
  %13 = call noalias align 16 i8* @malloc(i64 %12) #7
  store i8* %13, i8** %6, align 8
  %14 = load i8*, i8** %6, align 8
  %15 = load i8*, i8** %4, align 8
  %16 = load i32, i32* %5, align 4
  %17 = add nsw i32 %16, 1
  %18 = sext i32 %17 to i64
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 1 %14, i8* align 1 %15, i64 %18, i1 false)
  %19 = load %struct.class_String*, %struct.class_String** %3, align 8
  %20 = getelementptr inbounds %struct.class_String, %struct.class_String* %19, i32 0, i32 2
  %21 = load i8*, i8** %20, align 8
  %22 = icmp ne i8* %21, null
  br i1 %22, label %23, label %27

23:                                               ; preds = %2
  %24 = load %struct.class_String*, %struct.class_String** %3, align 8
  %25 = getelementptr inbounds %struct.class_String, %struct.class_String* %24, i32 0, i32 2
  %26 = load i8*, i8** %25, align 8
  call void @free(i8* %26) #7
  br label %27

27:                                               ; preds = %23, %2
  %28 = load i8*, i8** %6, align 8
  %29 = load %struct.class_String*, %struct.class_String** %3, align 8
  %30 = getelementptr inbounds %struct.class_String, %struct.class_String* %29, i32 0, i32 2
  store i8* %28, i8** %30, align 8
  %31 = load i32, i32* %5, align 4
  %32 = load %struct.class_String*, %struct.class_String** %3, align 8
  %33 = getelementptr inbounds %struct.class_String, %struct.class_String* %32, i32 0, i32 3
  store i32 %31, i32* %33, align 8
  %34 = load i32, i32* %5, align 4
  %35 = load %struct.class_String*, %struct.class_String** %3, align 8
  %36 = getelementptr inbounds %struct.class_String, %struct.class_String* %35, i32 0, i32 4
  store i32 %34, i32* %36, align 4
  ret void
}

; Function Attrs: nounwind readonly willreturn
declare dso_local i64 @strlen(i8*) #3

; Function Attrs: nounwind
declare dso_local noalias align 16 i8* @malloc(i64) #1

; Function Attrs: argmemonly nofree nounwind willreturn
declare void @llvm.memcpy.p0i8.p0i8.i64(i8* noalias nocapture writeonly, i8* noalias nocapture readonly, i64, i1 immarg) #4

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @String_public_Resize(%struct.class_String* %0, i32 %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i32, align 4
  %5 = alloca i8*, align 8
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i32 %1, i32* %4, align 4
  %6 = load i32, i32* %4, align 4
  %7 = sext i32 %6 to i64
  %8 = call noalias align 16 i8* @malloc(i64 %7) #7
  store i8* %8, i8** %5, align 8
  %9 = load i8*, i8** %5, align 8
  %10 = load %struct.class_String*, %struct.class_String** %3, align 8
  %11 = getelementptr inbounds %struct.class_String, %struct.class_String* %10, i32 0, i32 2
  %12 = load i8*, i8** %11, align 8
  %13 = load %struct.class_String*, %struct.class_String** %3, align 8
  %14 = getelementptr inbounds %struct.class_String, %struct.class_String* %13, i32 0, i32 3
  %15 = load i32, i32* %14, align 8
  %16 = sext i32 %15 to i64
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 1 %9, i8* align 1 %12, i64 %16, i1 false)
  %17 = load %struct.class_String*, %struct.class_String** %3, align 8
  %18 = getelementptr inbounds %struct.class_String, %struct.class_String* %17, i32 0, i32 2
  %19 = load i8*, i8** %18, align 8
  call void @free(i8* %19) #7
  %20 = load i8*, i8** %5, align 8
  %21 = load %struct.class_String*, %struct.class_String** %3, align 8
  %22 = getelementptr inbounds %struct.class_String, %struct.class_String* %21, i32 0, i32 2
  store i8* %20, i8** %22, align 8
  %23 = load i32, i32* %4, align 4
  %24 = load %struct.class_String*, %struct.class_String** %3, align 8
  %25 = getelementptr inbounds %struct.class_String, %struct.class_String* %24, i32 0, i32 4
  store i32 %23, i32* %25, align 4
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @String_public_AddChar(%struct.class_String* %0, i8 signext %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i8, align 1
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i8 %1, i8* %4, align 1
  %5 = load %struct.class_String*, %struct.class_String** %3, align 8
  %6 = getelementptr inbounds %struct.class_String, %struct.class_String* %5, i32 0, i32 3
  %7 = load i32, i32* %6, align 8
  %8 = load %struct.class_String*, %struct.class_String** %3, align 8
  %9 = getelementptr inbounds %struct.class_String, %struct.class_String* %8, i32 0, i32 4
  %10 = load i32, i32* %9, align 4
  %11 = icmp eq i32 %7, %10
  br i1 %11, label %12, label %21

12:                                               ; preds = %2
  %13 = load %struct.class_String*, %struct.class_String** %3, align 8
  %14 = load %struct.class_String*, %struct.class_String** %3, align 8
  %15 = getelementptr inbounds %struct.class_String, %struct.class_String* %14, i32 0, i32 4
  %16 = load i32, i32* %15, align 4
  %17 = load %struct.class_String*, %struct.class_String** %3, align 8
  %18 = getelementptr inbounds %struct.class_String, %struct.class_String* %17, i32 0, i32 5
  %19 = load i32, i32* %18, align 8
  %20 = add nsw i32 %16, %19
  call void @String_public_Resize(%struct.class_String* %13, i32 %20)
  br label %21

21:                                               ; preds = %12, %2
  %22 = load i8, i8* %4, align 1
  %23 = load %struct.class_String*, %struct.class_String** %3, align 8
  %24 = getelementptr inbounds %struct.class_String, %struct.class_String* %23, i32 0, i32 2
  %25 = load i8*, i8** %24, align 8
  %26 = load %struct.class_String*, %struct.class_String** %3, align 8
  %27 = getelementptr inbounds %struct.class_String, %struct.class_String* %26, i32 0, i32 3
  %28 = load i32, i32* %27, align 8
  %29 = sext i32 %28 to i64
  %30 = getelementptr inbounds i8, i8* %25, i64 %29
  store i8 %22, i8* %30, align 1
  %31 = load %struct.class_String*, %struct.class_String** %3, align 8
  %32 = getelementptr inbounds %struct.class_String, %struct.class_String* %31, i32 0, i32 3
  %33 = load i32, i32* %32, align 8
  %34 = add nsw i32 %33, 1
  store i32 %34, i32* %32, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local %struct.class_String* @String_public_Concat(%struct.class_String* %0, %struct.class_String* %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca %struct.class_String*, align 8
  %5 = alloca i8*, align 8
  %6 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store %struct.class_String* %1, %struct.class_String** %4, align 8
  %7 = load %struct.class_String*, %struct.class_String** %3, align 8
  %8 = getelementptr inbounds %struct.class_String, %struct.class_String* %7, i32 0, i32 3
  %9 = load i32, i32* %8, align 8
  %10 = load %struct.class_String*, %struct.class_String** %4, align 8
  %11 = getelementptr inbounds %struct.class_String, %struct.class_String* %10, i32 0, i32 3
  %12 = load i32, i32* %11, align 8
  %13 = add nsw i32 %9, %12
  %14 = add nsw i32 %13, 1
  %15 = sext i32 %14 to i64
  %16 = call noalias align 16 i8* @malloc(i64 %15) #7
  store i8* %16, i8** %5, align 8
  %17 = load i8*, i8** %5, align 8
  %18 = load %struct.class_String*, %struct.class_String** %3, align 8
  %19 = getelementptr inbounds %struct.class_String, %struct.class_String* %18, i32 0, i32 2
  %20 = load i8*, i8** %19, align 8
  %21 = call i8* @strcpy(i8* %17, i8* %20) #7
  %22 = load i8*, i8** %5, align 8
  %23 = load %struct.class_String*, %struct.class_String** %4, align 8
  %24 = getelementptr inbounds %struct.class_String, %struct.class_String* %23, i32 0, i32 2
  %25 = load i8*, i8** %24, align 8
  %26 = call i8* @strcat(i8* %22, i8* %25) #7
  %27 = call noalias align 16 i8* @malloc(i64 40) #7
  %28 = bitcast i8* %27 to %struct.class_String*
  store %struct.class_String* %28, %struct.class_String** %6, align 8
  %29 = load %struct.class_String*, %struct.class_String** %6, align 8
  call void @String_public_Constructor(%struct.class_String* %29)
  %30 = load %struct.class_String*, %struct.class_String** %6, align 8
  %31 = load i8*, i8** %5, align 8
  call void @String_public_Load(%struct.class_String* %30, i8* %31)
  %32 = load %struct.class_String*, %struct.class_String** %6, align 8
  %33 = bitcast %struct.class_String* %32 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* %33)
  %34 = load i8*, i8** %5, align 8
  call void @free(i8* %34) #7
  %35 = load %struct.class_String*, %struct.class_String** %6, align 8
  ret %struct.class_String* %35
}

; Function Attrs: nounwind
declare dso_local i8* @strcpy(i8*, i8*) #1

; Function Attrs: nounwind
declare dso_local i8* @strcat(i8*, i8*) #1

; Function Attrs: noinline nounwind optnone uwtable
define dso_local zeroext i1 @String_public_Equal(%struct.class_String* %0, %struct.class_String* %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca %struct.class_String*, align 8
  %5 = alloca i32, align 4
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store %struct.class_String* %1, %struct.class_String** %4, align 8
  %6 = load %struct.class_String*, %struct.class_String** %3, align 8
  %7 = getelementptr inbounds %struct.class_String, %struct.class_String* %6, i32 0, i32 2
  %8 = load i8*, i8** %7, align 8
  %9 = load %struct.class_String*, %struct.class_String** %4, align 8
  %10 = getelementptr inbounds %struct.class_String, %struct.class_String* %9, i32 0, i32 2
  %11 = load i8*, i8** %10, align 8
  %12 = call i32 @strcmp(i8* %8, i8* %11) #8
  store i32 %12, i32* %5, align 4
  %13 = load i32, i32* %5, align 4
  %14 = icmp eq i32 %13, 0
  ret i1 %14
}

; Function Attrs: nounwind readonly willreturn
declare dso_local i32 @strcmp(i8*, i8*) #3

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i8* @String_public_GetBuffer(%struct.class_String* %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = getelementptr inbounds %struct.class_String, %struct.class_String* %3, i32 0, i32 2
  %5 = load i8*, i8** %4, align 8
  ret i8* %5
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i32 @String_public_GetLength(%struct.class_String* %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = getelementptr inbounds %struct.class_String, %struct.class_String* %3, i32 0, i32 3
  %5 = load i32, i32* %4, align 8
  ret i32 %5
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local %struct.class_String* @String_public_Substring(%struct.class_String* %0, i32 %1, i32 %2) #0 {
  %4 = alloca %struct.class_String*, align 8
  %5 = alloca i32, align 4
  %6 = alloca i32, align 4
  %7 = alloca i8*, align 8
  %8 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %4, align 8
  store i32 %1, i32* %5, align 4
  store i32 %2, i32* %6, align 4
  %9 = load i32, i32* %5, align 4
  %10 = icmp slt i32 %9, 0
  br i1 %10, label %22, label %11

11:                                               ; preds = %3
  %12 = load i32, i32* %6, align 4
  %13 = icmp slt i32 %12, 0
  br i1 %13, label %22, label %14

14:                                               ; preds = %11
  %15 = load i32, i32* %5, align 4
  %16 = load i32, i32* %6, align 4
  %17 = add nsw i32 %15, %16
  %18 = load %struct.class_String*, %struct.class_String** %4, align 8
  %19 = getelementptr inbounds %struct.class_String, %struct.class_String* %18, i32 0, i32 3
  %20 = load i32, i32* %19, align 8
  %21 = icmp sgt i32 %17, %20
  br i1 %21, label %22, label %23

22:                                               ; preds = %14, %11, %3
  store i8* getelementptr inbounds ([44 x i8], [44 x i8]* @.str.2.8, i64 0, i64 0), i8** %7, align 8
  br label %41

23:                                               ; preds = %14
  %24 = load i32, i32* %6, align 4
  %25 = add nsw i32 %24, 1
  %26 = sext i32 %25 to i64
  %27 = call noalias align 16 i8* @malloc(i64 %26) #7
  store i8* %27, i8** %7, align 8
  %28 = load i8*, i8** %7, align 8
  %29 = load %struct.class_String*, %struct.class_String** %4, align 8
  %30 = getelementptr inbounds %struct.class_String, %struct.class_String* %29, i32 0, i32 2
  %31 = load i8*, i8** %30, align 8
  %32 = load i32, i32* %5, align 4
  %33 = sext i32 %32 to i64
  %34 = getelementptr inbounds i8, i8* %31, i64 %33
  %35 = load i32, i32* %6, align 4
  %36 = sext i32 %35 to i64
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 1 %28, i8* align 1 %34, i64 %36, i1 false)
  %37 = load i8*, i8** %7, align 8
  %38 = load i32, i32* %6, align 4
  %39 = sext i32 %38 to i64
  %40 = getelementptr inbounds i8, i8* %37, i64 %39
  store i8 0, i8* %40, align 1
  br label %41

41:                                               ; preds = %23, %22
  %42 = call noalias align 16 i8* @malloc(i64 40) #7
  %43 = bitcast i8* %42 to %struct.class_String*
  store %struct.class_String* %43, %struct.class_String** %8, align 8
  %44 = load %struct.class_String*, %struct.class_String** %8, align 8
  call void @String_public_Constructor(%struct.class_String* %44)
  %45 = load %struct.class_String*, %struct.class_String** %8, align 8
  %46 = load i8*, i8** %7, align 8
  call void @String_public_Load(%struct.class_String* %45, i8* %46)
  %47 = load %struct.class_String*, %struct.class_String** %8, align 8
  %48 = bitcast %struct.class_String* %47 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* %48)
  %49 = load i8*, i8** %7, align 8
  call void @free(i8* %49) #7
  %50 = load %struct.class_String*, %struct.class_String** %8, align 8
  ret %struct.class_String* %50
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Int_public_Constructor(%struct.class_Int* %0, i32 %1) #0 {
  %3 = alloca %struct.class_Int*, align 8
  %4 = alloca i32, align 4
  store %struct.class_Int* %0, %struct.class_Int** %3, align 8
  store i32 %1, i32* %4, align 4
  %5 = load %struct.class_Int*, %struct.class_Int** %3, align 8
  %6 = getelementptr inbounds %struct.class_Int, %struct.class_Int* %5, i32 0, i32 0
  store %struct.String_vTable* @Int_vTable_Const, %struct.String_vTable** %6, align 8
  %7 = load %struct.class_Int*, %struct.class_Int** %3, align 8
  %8 = getelementptr inbounds %struct.class_Int, %struct.class_Int* %7, i32 0, i32 1
  store i32 0, i32* %8, align 8
  %9 = load i32, i32* %4, align 4
  %10 = load %struct.class_Int*, %struct.class_Int** %3, align 8
  %11 = getelementptr inbounds %struct.class_Int, %struct.class_Int* %10, i32 0, i32 2
  store i32 %9, i32* %11, align 4
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i32 @Int_public_GetValue(%struct.class_Int* %0) #0 {
  %2 = alloca i32, align 4
  %3 = alloca %struct.class_Int*, align 8
  store %struct.class_Int* %0, %struct.class_Int** %3, align 8
  %4 = load %struct.class_Int*, %struct.class_Int** %3, align 8
  %5 = icmp eq %struct.class_Int* %4, null
  br i1 %5, label %6, label %7

6:                                                ; preds = %1
  store i32 0, i32* %2, align 4
  br label %11

7:                                                ; preds = %1
  %8 = load %struct.class_Int*, %struct.class_Int** %3, align 8
  %9 = getelementptr inbounds %struct.class_Int, %struct.class_Int* %8, i32 0, i32 2
  %10 = load i32, i32* %9, align 4
  store i32 %10, i32* %2, align 4
  br label %11

11:                                               ; preds = %7, %6
  %12 = load i32, i32* %2, align 4
  ret i32 %12
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Float_public_Constructor(%struct.class_Float* %0, float %1) #0 {
  %3 = alloca %struct.class_Float*, align 8
  %4 = alloca float, align 4
  store %struct.class_Float* %0, %struct.class_Float** %3, align 8
  store float %1, float* %4, align 4
  %5 = load %struct.class_Float*, %struct.class_Float** %3, align 8
  %6 = getelementptr inbounds %struct.class_Float, %struct.class_Float* %5, i32 0, i32 0
  store %struct.String_vTable* @Float_vTable_Const, %struct.String_vTable** %6, align 8
  %7 = load %struct.class_Float*, %struct.class_Float** %3, align 8
  %8 = getelementptr inbounds %struct.class_Float, %struct.class_Float* %7, i32 0, i32 1
  store i32 0, i32* %8, align 8
  %9 = load float, float* %4, align 4
  %10 = load %struct.class_Float*, %struct.class_Float** %3, align 8
  %11 = getelementptr inbounds %struct.class_Float, %struct.class_Float* %10, i32 0, i32 2
  store float %9, float* %11, align 4
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local float @Float_public_GetValue(%struct.class_Float* %0) #0 {
  %2 = alloca float, align 4
  %3 = alloca %struct.class_Float*, align 8
  store %struct.class_Float* %0, %struct.class_Float** %3, align 8
  %4 = load %struct.class_Float*, %struct.class_Float** %3, align 8
  %5 = icmp eq %struct.class_Float* %4, null
  br i1 %5, label %6, label %7

6:                                                ; preds = %1
  store float 0.000000e+00, float* %2, align 4
  br label %11

7:                                                ; preds = %1
  %8 = load %struct.class_Float*, %struct.class_Float** %3, align 8
  %9 = getelementptr inbounds %struct.class_Float, %struct.class_Float* %8, i32 0, i32 2
  %10 = load float, float* %9, align 4
  store float %10, float* %2, align 4
  br label %11

11:                                               ; preds = %7, %6
  %12 = load float, float* %2, align 4
  ret float %12
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Bool_public_Constructor(%struct.class_Bool* %0, i1 zeroext %1) #0 {
  %3 = alloca %struct.class_Bool*, align 8
  %4 = alloca i8, align 1
  store %struct.class_Bool* %0, %struct.class_Bool** %3, align 8
  %5 = zext i1 %1 to i8
  store i8 %5, i8* %4, align 1
  %6 = load %struct.class_Bool*, %struct.class_Bool** %3, align 8
  %7 = getelementptr inbounds %struct.class_Bool, %struct.class_Bool* %6, i32 0, i32 0
  store %struct.String_vTable* @Bool_vTable_Const, %struct.String_vTable** %7, align 8
  %8 = load %struct.class_Bool*, %struct.class_Bool** %3, align 8
  %9 = getelementptr inbounds %struct.class_Bool, %struct.class_Bool* %8, i32 0, i32 1
  store i32 0, i32* %9, align 8
  %10 = load i8, i8* %4, align 1
  %11 = trunc i8 %10 to i1
  %12 = load %struct.class_Bool*, %struct.class_Bool** %3, align 8
  %13 = getelementptr inbounds %struct.class_Bool, %struct.class_Bool* %12, i32 0, i32 2
  %14 = zext i1 %11 to i8
  store i8 %14, i8* %13, align 4
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local zeroext i1 @Bool_public_GetValue(%struct.class_Bool* %0) #0 {
  %2 = alloca i1, align 1
  %3 = alloca %struct.class_Bool*, align 8
  store %struct.class_Bool* %0, %struct.class_Bool** %3, align 8
  %4 = load %struct.class_Bool*, %struct.class_Bool** %3, align 8
  %5 = icmp eq %struct.class_Bool* %4, null
  br i1 %5, label %6, label %7

6:                                                ; preds = %1
  store i1 false, i1* %2, align 1
  br label %12

7:                                                ; preds = %1
  %8 = load %struct.class_Bool*, %struct.class_Bool** %3, align 8
  %9 = getelementptr inbounds %struct.class_Bool, %struct.class_Bool* %8, i32 0, i32 2
  %10 = load i8, i8* %9, align 4
  %11 = trunc i8 %10 to i1
  store i1 %11, i1* %2, align 1
  br label %12

12:                                               ; preds = %7, %6
  %13 = load i1, i1* %2, align 1
  ret i1 %13
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Array_public_Constructor(%struct.class_Array* %0, i32 %1) #0 {
  %3 = alloca %struct.class_Array*, align 8
  %4 = alloca i32, align 4
  store %struct.class_Array* %0, %struct.class_Array** %3, align 8
  store i32 %1, i32* %4, align 4
  %5 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %6 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %5, i32 0, i32 0
  store %struct.String_vTable* @Array_vTable_Const, %struct.String_vTable** %6, align 8
  %7 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %8 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %7, i32 0, i32 1
  store i32 0, i32* %8, align 8
  %9 = load i32, i32* %4, align 4
  %10 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %11 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %10, i32 0, i32 3
  store i32 %9, i32* %11, align 8
  %12 = load i32, i32* %4, align 4
  %13 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %14 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %13, i32 0, i32 4
  store i32 %12, i32* %14, align 4
  %15 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %16 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %15, i32 0, i32 5
  store i32 5, i32* %16, align 8
  %17 = load i32, i32* %4, align 4
  %18 = sext i32 %17 to i64
  %19 = call noalias align 16 i8* @calloc(i64 %18, i64 8) #7
  %20 = bitcast i8* %19 to %struct.class_Any**
  %21 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %22 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %21, i32 0, i32 2
  store %struct.class_Any** %20, %struct.class_Any*** %22, align 8
  ret void
}

; Function Attrs: nounwind
declare dso_local noalias align 16 i8* @calloc(i64, i64) #1

; Function Attrs: noinline nounwind optnone uwtable
define dso_local %struct.class_Any* @Array_public_GetElement(%struct.class_Array* %0, i32 %1) #0 {
  %3 = alloca %struct.class_Any*, align 8
  %4 = alloca %struct.class_Array*, align 8
  %5 = alloca i32, align 4
  store %struct.class_Array* %0, %struct.class_Array** %4, align 8
  store i32 %1, i32* %5, align 4
  %6 = load i32, i32* %5, align 4
  %7 = icmp slt i32 %6, 0
  br i1 %7, label %14, label %8

8:                                                ; preds = %2
  %9 = load i32, i32* %5, align 4
  %10 = load %struct.class_Array*, %struct.class_Array** %4, align 8
  %11 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %10, i32 0, i32 3
  %12 = load i32, i32* %11, align 8
  %13 = icmp sge i32 %9, %12
  br i1 %13, label %14, label %15

14:                                               ; preds = %8, %2
  store %struct.class_Any* null, %struct.class_Any** %3, align 8
  br label %23

15:                                               ; preds = %8
  %16 = load %struct.class_Array*, %struct.class_Array** %4, align 8
  %17 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %16, i32 0, i32 2
  %18 = load %struct.class_Any**, %struct.class_Any*** %17, align 8
  %19 = load i32, i32* %5, align 4
  %20 = sext i32 %19 to i64
  %21 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %18, i64 %20
  %22 = load %struct.class_Any*, %struct.class_Any** %21, align 8
  store %struct.class_Any* %22, %struct.class_Any** %3, align 8
  br label %23

23:                                               ; preds = %15, %14
  %24 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  ret %struct.class_Any* %24
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Array_public_SetElement(%struct.class_Array* %0, i32 %1, %struct.class_Any* %2) #0 {
  %4 = alloca %struct.class_Array*, align 8
  %5 = alloca i32, align 4
  %6 = alloca %struct.class_Any*, align 8
  store %struct.class_Array* %0, %struct.class_Array** %4, align 8
  store i32 %1, i32* %5, align 4
  store %struct.class_Any* %2, %struct.class_Any** %6, align 8
  %7 = load i32, i32* %5, align 4
  %8 = icmp slt i32 %7, 0
  br i1 %8, label %15, label %9

9:                                                ; preds = %3
  %10 = load i32, i32* %5, align 4
  %11 = load %struct.class_Array*, %struct.class_Array** %4, align 8
  %12 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %11, i32 0, i32 3
  %13 = load i32, i32* %12, align 8
  %14 = icmp sge i32 %10, %13
  br i1 %14, label %15, label %16

15:                                               ; preds = %9, %3
  br label %32

16:                                               ; preds = %9
  %17 = load %struct.class_Any*, %struct.class_Any** %6, align 8
  call void @arc_RegisterReference(%struct.class_Any* %17)
  %18 = load %struct.class_Array*, %struct.class_Array** %4, align 8
  %19 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %18, i32 0, i32 2
  %20 = load %struct.class_Any**, %struct.class_Any*** %19, align 8
  %21 = load i32, i32* %5, align 4
  %22 = sext i32 %21 to i64
  %23 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %20, i64 %22
  %24 = load %struct.class_Any*, %struct.class_Any** %23, align 8
  call void @arc_UnregisterReference(%struct.class_Any* %24)
  %25 = load %struct.class_Any*, %struct.class_Any** %6, align 8
  %26 = load %struct.class_Array*, %struct.class_Array** %4, align 8
  %27 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %26, i32 0, i32 2
  %28 = load %struct.class_Any**, %struct.class_Any*** %27, align 8
  %29 = load i32, i32* %5, align 4
  %30 = sext i32 %29 to i64
  %31 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %28, i64 %30
  store %struct.class_Any* %25, %struct.class_Any** %31, align 8
  br label %32

32:                                               ; preds = %16, %15
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i32 @Array_public_GetLength(%struct.class_Array* %0) #0 {
  %2 = alloca %struct.class_Array*, align 8
  store %struct.class_Array* %0, %struct.class_Array** %2, align 8
  %3 = load %struct.class_Array*, %struct.class_Array** %2, align 8
  %4 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %3, i32 0, i32 3
  %5 = load i32, i32* %4, align 8
  ret i32 %5
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Array_public_Push(%struct.class_Array* %0, %struct.class_Any* %1) #0 {
  %3 = alloca %struct.class_Array*, align 8
  %4 = alloca %struct.class_Any*, align 8
  %5 = alloca i32, align 4
  %6 = alloca %struct.class_Any**, align 8
  %7 = alloca i32, align 4
  store %struct.class_Array* %0, %struct.class_Array** %3, align 8
  store %struct.class_Any* %1, %struct.class_Any** %4, align 8
  %8 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %9 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %8, i32 0, i32 3
  %10 = load i32, i32* %9, align 8
  %11 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %12 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %11, i32 0, i32 4
  %13 = load i32, i32* %12, align 4
  %14 = icmp eq i32 %10, %13
  br i1 %14, label %15, label %71

15:                                               ; preds = %2
  %16 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %17 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %16, i32 0, i32 3
  %18 = load i32, i32* %17, align 8
  %19 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %20 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %19, i32 0, i32 5
  %21 = load i32, i32* %20, align 8
  %22 = add nsw i32 %18, %21
  store i32 %22, i32* %5, align 4
  %23 = load i32, i32* %5, align 4
  %24 = sext i32 %23 to i64
  %25 = mul i64 8, %24
  %26 = call noalias align 16 i8* @malloc(i64 %25) #7
  %27 = bitcast i8* %26 to %struct.class_Any**
  store %struct.class_Any** %27, %struct.class_Any*** %6, align 8
  %28 = load %struct.class_Any**, %struct.class_Any*** %6, align 8
  %29 = bitcast %struct.class_Any** %28 to i8*
  %30 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %31 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %30, i32 0, i32 2
  %32 = load %struct.class_Any**, %struct.class_Any*** %31, align 8
  %33 = bitcast %struct.class_Any** %32 to i8*
  %34 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %35 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %34, i32 0, i32 3
  %36 = load i32, i32* %35, align 8
  %37 = sext i32 %36 to i64
  %38 = mul i64 8, %37
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 8 %29, i8* align 8 %33, i64 %38, i1 false)
  %39 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %40 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %39, i32 0, i32 2
  %41 = load %struct.class_Any**, %struct.class_Any*** %40, align 8
  %42 = bitcast %struct.class_Any** %41 to i8*
  call void @free(i8* %42) #7
  %43 = load %struct.class_Any**, %struct.class_Any*** %6, align 8
  %44 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %45 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %44, i32 0, i32 2
  store %struct.class_Any** %43, %struct.class_Any*** %45, align 8
  store i32 0, i32* %7, align 4
  br label %46

46:                                               ; preds = %64, %15
  %47 = load i32, i32* %7, align 4
  %48 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %49 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %48, i32 0, i32 5
  %50 = load i32, i32* %49, align 8
  %51 = icmp slt i32 %47, %50
  br i1 %51, label %52, label %67

52:                                               ; preds = %46
  %53 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %54 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %53, i32 0, i32 2
  %55 = load %struct.class_Any**, %struct.class_Any*** %54, align 8
  %56 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %57 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %56, i32 0, i32 3
  %58 = load i32, i32* %57, align 8
  %59 = sext i32 %58 to i64
  %60 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %55, i64 %59
  %61 = load i32, i32* %7, align 4
  %62 = sext i32 %61 to i64
  %63 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %60, i64 %62
  store %struct.class_Any* null, %struct.class_Any** %63, align 8
  br label %64

64:                                               ; preds = %52
  %65 = load i32, i32* %7, align 4
  %66 = add nsw i32 %65, 1
  store i32 %66, i32* %7, align 4
  br label %46, !llvm.loop !6

67:                                               ; preds = %46
  %68 = load i32, i32* %5, align 4
  %69 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %70 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %69, i32 0, i32 4
  store i32 %68, i32* %70, align 4
  br label %71

71:                                               ; preds = %67, %2
  %72 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %73 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %72, i32 0, i32 3
  %74 = load i32, i32* %73, align 8
  %75 = add nsw i32 %74, 1
  store i32 %75, i32* %73, align 8
  %76 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %77 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %78 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %77, i32 0, i32 3
  %79 = load i32, i32* %78, align 8
  %80 = sub nsw i32 %79, 1
  %81 = load %struct.class_Any*, %struct.class_Any** %4, align 8
  call void @Array_public_SetElement(%struct.class_Array* %76, i32 %80, %struct.class_Any* %81)
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @pArray_public_Constructor(%struct.class_pArray* %0, i32 %1, i32 %2) #0 {
  %4 = alloca %struct.class_pArray*, align 8
  %5 = alloca i32, align 4
  %6 = alloca i32, align 4
  store %struct.class_pArray* %0, %struct.class_pArray** %4, align 8
  store i32 %1, i32* %5, align 4
  store i32 %2, i32* %6, align 4
  %7 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %8 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %7, i32 0, i32 0
  store %struct.String_vTable* @pArray_vTable_Const, %struct.String_vTable** %8, align 8
  %9 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %10 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %9, i32 0, i32 1
  store i32 0, i32* %10, align 8
  %11 = load i32, i32* %5, align 4
  %12 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %13 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %12, i32 0, i32 3
  store i32 %11, i32* %13, align 8
  %14 = load i32, i32* %5, align 4
  %15 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %16 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %15, i32 0, i32 4
  store i32 %14, i32* %16, align 4
  %17 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %18 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %17, i32 0, i32 5
  store i32 5, i32* %18, align 8
  %19 = load i32, i32* %6, align 4
  %20 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %21 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %20, i32 0, i32 6
  store i32 %19, i32* %21, align 4
  %22 = load i32, i32* %5, align 4
  %23 = sext i32 %22 to i64
  %24 = load i32, i32* %6, align 4
  %25 = sext i32 %24 to i64
  %26 = call noalias align 16 i8* @calloc(i64 %23, i64 %25) #7
  %27 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %28 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %27, i32 0, i32 2
  store i8* %26, i8** %28, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i32 @pArray_public_GetLength(%struct.class_pArray* %0) #0 {
  %2 = alloca %struct.class_pArray*, align 8
  store %struct.class_pArray* %0, %struct.class_pArray** %2, align 8
  %3 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %4 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %3, i32 0, i32 3
  %5 = load i32, i32* %4, align 8
  ret i32 %5
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i8* @pArray_public_Grow(%struct.class_pArray* %0) #0 {
  %2 = alloca %struct.class_pArray*, align 8
  %3 = alloca i32, align 4
  %4 = alloca i8*, align 8
  store %struct.class_pArray* %0, %struct.class_pArray** %2, align 8
  %5 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %6 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %5, i32 0, i32 3
  %7 = load i32, i32* %6, align 8
  %8 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %9 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %8, i32 0, i32 4
  %10 = load i32, i32* %9, align 4
  %11 = icmp eq i32 %7, %10
  br i1 %11, label %12, label %54

12:                                               ; preds = %1
  %13 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %14 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %13, i32 0, i32 3
  %15 = load i32, i32* %14, align 8
  %16 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %17 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %16, i32 0, i32 5
  %18 = load i32, i32* %17, align 8
  %19 = add nsw i32 %15, %18
  %20 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %21 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %20, i32 0, i32 6
  %22 = load i32, i32* %21, align 4
  %23 = mul nsw i32 %19, %22
  store i32 %23, i32* %3, align 4
  %24 = load i32, i32* %3, align 4
  %25 = sext i32 %24 to i64
  %26 = call noalias align 16 i8* @malloc(i64 %25) #7
  store i8* %26, i8** %4, align 8
  %27 = load i8*, i8** %4, align 8
  %28 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %29 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %28, i32 0, i32 2
  %30 = load i8*, i8** %29, align 8
  %31 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %32 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %31, i32 0, i32 3
  %33 = load i32, i32* %32, align 8
  %34 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %35 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %34, i32 0, i32 6
  %36 = load i32, i32* %35, align 4
  %37 = mul nsw i32 %33, %36
  %38 = sext i32 %37 to i64
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 1 %27, i8* align 1 %30, i64 %38, i1 false)
  %39 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %40 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %39, i32 0, i32 2
  %41 = load i8*, i8** %40, align 8
  call void @free(i8* %41) #7
  %42 = load i8*, i8** %4, align 8
  %43 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %44 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %43, i32 0, i32 2
  store i8* %42, i8** %44, align 8
  %45 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %46 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %45, i32 0, i32 3
  %47 = load i32, i32* %46, align 8
  %48 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %49 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %48, i32 0, i32 5
  %50 = load i32, i32* %49, align 8
  %51 = add nsw i32 %47, %50
  %52 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %53 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %52, i32 0, i32 4
  store i32 %51, i32* %53, align 4
  br label %54

54:                                               ; preds = %12, %1
  %55 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %56 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %55, i32 0, i32 3
  %57 = load i32, i32* %56, align 8
  %58 = add nsw i32 %57, 1
  store i32 %58, i32* %56, align 8
  %59 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %60 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %59, i32 0, i32 2
  %61 = load i8*, i8** %60, align 8
  %62 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %63 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %62, i32 0, i32 3
  %64 = load i32, i32* %63, align 8
  %65 = sub nsw i32 %64, 1
  %66 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %67 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %66, i32 0, i32 6
  %68 = load i32, i32* %67, align 4
  %69 = mul nsw i32 %65, %68
  %70 = sext i32 %69 to i64
  %71 = getelementptr i8, i8* %61, i64 %70
  ret i8* %71
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i8* @pArray_public_GetElementPtr(%struct.class_pArray* %0, i32 %1) #0 {
  %3 = alloca i8*, align 8
  %4 = alloca %struct.class_pArray*, align 8
  %5 = alloca i32, align 4
  store %struct.class_pArray* %0, %struct.class_pArray** %4, align 8
  store i32 %1, i32* %5, align 4
  %6 = load i32, i32* %5, align 4
  %7 = icmp slt i32 %6, 0
  br i1 %7, label %14, label %8

8:                                                ; preds = %2
  %9 = load i32, i32* %5, align 4
  %10 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %11 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %10, i32 0, i32 3
  %12 = load i32, i32* %11, align 8
  %13 = icmp sge i32 %9, %12
  br i1 %13, label %14, label %15

14:                                               ; preds = %8, %2
  store i8* null, i8** %3, align 8
  br label %26

15:                                               ; preds = %8
  %16 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %17 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %16, i32 0, i32 2
  %18 = load i8*, i8** %17, align 8
  %19 = load i32, i32* %5, align 4
  %20 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %21 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %20, i32 0, i32 6
  %22 = load i32, i32* %21, align 4
  %23 = mul nsw i32 %19, %22
  %24 = sext i32 %23 to i64
  %25 = getelementptr i8, i8* %18, i64 %24
  store i8* %25, i8** %3, align 8
  br label %26

26:                                               ; preds = %15, %14
  %27 = load i8*, i8** %3, align 8
  ret i8* %27
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Action_public_Constructor(%struct.class_Action* %0, i8* (i8*)* %1, i8* %2) #0 {
  %4 = alloca %struct.class_Action*, align 8
  %5 = alloca i8* (i8*)*, align 8
  %6 = alloca i8*, align 8
  store %struct.class_Action* %0, %struct.class_Action** %4, align 8
  store i8* (i8*)* %1, i8* (i8*)** %5, align 8
  store i8* %2, i8** %6, align 8
  %7 = load %struct.class_Action*, %struct.class_Action** %4, align 8
  %8 = getelementptr inbounds %struct.class_Action, %struct.class_Action* %7, i32 0, i32 0
  store %struct.Any_vTable* @Action_vTable_Const, %struct.Any_vTable** %8, align 8
  %9 = load %struct.class_Action*, %struct.class_Action** %4, align 8
  %10 = getelementptr inbounds %struct.class_Action, %struct.class_Action* %9, i32 0, i32 1
  store i32 0, i32* %10, align 8
  %11 = load i8* (i8*)*, i8* (i8*)** %5, align 8
  %12 = load %struct.class_Action*, %struct.class_Action** %4, align 8
  %13 = getelementptr inbounds %struct.class_Action, %struct.class_Action* %12, i32 0, i32 2
  store i8* (i8*)* %11, i8* (i8*)** %13, align 8
  %14 = load i8*, i8** %6, align 8
  %15 = load %struct.class_Action*, %struct.class_Action** %4, align 8
  %16 = getelementptr inbounds %struct.class_Action, %struct.class_Action* %15, i32 0, i32 3
  store i8* %14, i8** %16, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Action_public_Start(%struct.class_Action* %0) #0 {
  %2 = alloca %struct.class_Action*, align 8
  store %struct.class_Action* %0, %struct.class_Action** %2, align 8
  %3 = load %struct.class_Action*, %struct.class_Action** %2, align 8
  %4 = getelementptr inbounds %struct.class_Action, %struct.class_Action* %3, i32 0, i32 4
  %5 = load %struct.class_Action*, %struct.class_Action** %2, align 8
  %6 = getelementptr inbounds %struct.class_Action, %struct.class_Action* %5, i32 0, i32 2
  %7 = load i8* (i8*)*, i8* (i8*)** %6, align 8
  %8 = load %struct.class_Action*, %struct.class_Action** %2, align 8
  %9 = getelementptr inbounds %struct.class_Action, %struct.class_Action* %8, i32 0, i32 3
  %10 = load i8*, i8** %9, align 8
  %11 = call i32 @pthread_create(i64* %4, %union.pthread_attr_t* null, i8* (i8*)* %7, i8* %10) #7
  %12 = load %struct.class_Action*, %struct.class_Action** %2, align 8
  %13 = getelementptr inbounds %struct.class_Action, %struct.class_Action* %12, i32 0, i32 4
  %14 = load i64, i64* %13, align 8
  %15 = call i32 @pthread_join(i64 %14, i8** null)
  ret void
}

; Function Attrs: nounwind
declare dso_local i32 @pthread_create(i64*, %union.pthread_attr_t*, i8* (i8*)*, i8*) #1

declare dso_local i32 @pthread_join(i64, i8**) #2

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @Action_public_Kill(%struct.class_Action* %0) #0 {
  %2 = alloca %struct.class_Action*, align 8
  store %struct.class_Action* %0, %struct.class_Action** %2, align 8
  call void @pthread_exit(i8* null) #9
  unreachable
}

; Function Attrs: noreturn
declare dso_local void @pthread_exit(i8*) #5

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @rct_Print(%struct.class_String* %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = getelementptr inbounds %struct.class_String, %struct.class_String* %3, i32 0, i32 2
  %5 = load i8*, i8** %4, align 8
  %6 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.9, i64 0, i64 0), i8* %5)
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @rct_Write(%struct.class_String* %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = getelementptr inbounds %struct.class_String, %struct.class_String* %3, i32 0, i32 2
  %5 = load i8*, i8** %4, align 8
  %6 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.str.1.10, i64 0, i64 0), i8* %5)
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local %struct.class_String* @rct_Input() #0 {
  %1 = alloca i8*, align 8
  %2 = alloca i8*, align 8
  %3 = alloca i32, align 4
  %4 = alloca %struct.class_String*, align 8
  %5 = call noalias align 16 i8* @malloc(i64 1042) #7
  store i8* %5, i8** %1, align 8
  store i32 0, i32* %3, align 4
  br label %6

6:                                                ; preds = %38, %0
  %7 = load i8*, i8** %1, align 8
  %8 = icmp ne i8* %7, null
  br i1 %8, label %9, label %18

9:                                                ; preds = %6
  %10 = call i32 @getchar()
  %11 = trunc i32 %10 to i8
  %12 = load i8*, i8** %1, align 8
  %13 = load i32, i32* %3, align 4
  %14 = sext i32 %13 to i64
  %15 = getelementptr inbounds i8, i8* %12, i64 %14
  store i8 %11, i8* %15, align 1
  %16 = sext i8 %11 to i32
  %17 = icmp ne i32 %16, 10
  br label %18

18:                                               ; preds = %9, %6
  %19 = phi i1 [ false, %6 ], [ %17, %9 ]
  br i1 %19, label %20, label %41

20:                                               ; preds = %18
  %21 = load i32, i32* %3, align 4
  %22 = srem i32 %21, 1042
  %23 = icmp eq i32 %22, 1041
  br i1 %23, label %24, label %37

24:                                               ; preds = %20
  %25 = load i8*, i8** %1, align 8
  %26 = load i32, i32* %3, align 4
  %27 = add nsw i32 1042, %26
  %28 = add nsw i32 %27, 1
  %29 = sext i32 %28 to i64
  %30 = mul i64 1, %29
  %31 = call align 16 i8* @realloc(i8* %25, i64 %30) #7
  store i8* %31, i8** %2, align 8
  %32 = icmp eq i8* %31, null
  br i1 %32, label %33, label %35

33:                                               ; preds = %24
  %34 = load i8*, i8** %1, align 8
  call void @free(i8* %34) #7
  br label %35

35:                                               ; preds = %33, %24
  %36 = load i8*, i8** %2, align 8
  store i8* %36, i8** %1, align 8
  br label %37

37:                                               ; preds = %35, %20
  br label %38

38:                                               ; preds = %37
  %39 = load i32, i32* %3, align 4
  %40 = add nsw i32 %39, 1
  store i32 %40, i32* %3, align 4
  br label %6, !llvm.loop !7

41:                                               ; preds = %18
  %42 = load i8*, i8** %1, align 8
  %43 = icmp ne i8* %42, null
  br i1 %43, label %44, label %49

44:                                               ; preds = %41
  %45 = load i8*, i8** %1, align 8
  %46 = load i32, i32* %3, align 4
  %47 = sext i32 %46 to i64
  %48 = getelementptr inbounds i8, i8* %45, i64 %47
  store i8 0, i8* %48, align 1
  br label %49

49:                                               ; preds = %44, %41
  %50 = call noalias align 16 i8* @malloc(i64 40) #7
  %51 = bitcast i8* %50 to %struct.class_String*
  store %struct.class_String* %51, %struct.class_String** %4, align 8
  %52 = load %struct.class_String*, %struct.class_String** %4, align 8
  call void @String_public_Constructor(%struct.class_String* %52)
  %53 = load %struct.class_String*, %struct.class_String** %4, align 8
  %54 = load i8*, i8** %1, align 8
  call void @String_public_Load(%struct.class_String* %53, i8* %54)
  %55 = load %struct.class_String*, %struct.class_String** %4, align 8
  %56 = bitcast %struct.class_String* %55 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* %56)
  %57 = load i8*, i8** %1, align 8
  %58 = icmp ne i8* %57, null
  br i1 %58, label %59, label %61

59:                                               ; preds = %49
  %60 = load i8*, i8** %1, align 8
  call void @free(i8* %60) #7
  br label %61

61:                                               ; preds = %59, %49
  %62 = load %struct.class_String*, %struct.class_String** %4, align 8
  ret %struct.class_String* %62
}

declare dso_local i32 @getchar() #2

; Function Attrs: nounwind
declare dso_local align 16 i8* @realloc(i8*, i64) #1

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @rct_Clear() #0 {
  %1 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.2.11, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @rct_SetCursor(i32 %0, i32 %1) #0 {
  %3 = alloca i32, align 4
  %4 = alloca i32, align 4
  store i32 %0, i32* %3, align 4
  store i32 %1, i32* %4, align 4
  %5 = load i32, i32* %4, align 4
  %6 = load i32, i32* %3, align 4
  %7 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([10 x i8], [10 x i8]* @.str.3.12, i64 0, i64 0), i32 27, i32 %5, i32 %6)
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @rct_SetCursorVisible(i1 zeroext %0) #0 {
  %2 = alloca i8, align 1
  %3 = zext i1 %0 to i8
  store i8 %3, i8* %2, align 1
  %4 = load i8, i8* %2, align 1
  %5 = trunc i8 %4 to i1
  %6 = zext i1 %5 to i8
  store i8 %6, i8* @isCursorVisible, align 1
  %7 = load i8, i8* %2, align 1
  %8 = trunc i8 %7 to i1
  br i1 %8, label %9, label %11

9:                                                ; preds = %1
  %10 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.4.13, i64 0, i64 0))
  br label %13

11:                                               ; preds = %1
  %12 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.4.13, i64 0, i64 0))
  br label %13

13:                                               ; preds = %11, %9
  ret void
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local zeroext i1 @rct_GetCursorVisible() #0 {
  %1 = load i8, i8* @isCursorVisible, align 1
  %2 = trunc i8 %1 to i1
  ret i1 %2
}

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i32 @rct_Random(i32 %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = call i32 @rand() #7
  %4 = load i32, i32* %2, align 4
  %5 = srem i32 %3, %4
  ret i32 %5
}

; Function Attrs: nounwind
declare dso_local i32 @rand() #1

; Function Attrs: noinline nounwind optnone uwtable
define dso_local void @rct_Sleep(i32 %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = sitofp i32 %3 to double
  %5 = fdiv double %4, 1.000000e+03
  %6 = fptoui double %5 to i32
  %7 = call i32 @sleep(i32 %6)
  ret void
}

declare dso_local i32 @sleep(i32) #2

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i32 @rct_Sqrt(i32 %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = sitofp i32 %3 to double
  %5 = call double @sqrt(double %4) #7
  %6 = call double @llvm.floor.f64(double %5)
  %7 = fptosi double %6 to i32
  ret i32 %7
}

; Function Attrs: nounwind
declare dso_local double @sqrt(double) #1

; Function Attrs: nofree nosync nounwind readnone speculatable willreturn
declare double @llvm.floor.f64(double) #6

; Function Attrs: noinline nounwind optnone uwtable
define dso_local i32 @rct_Now() #0 {
  %1 = call i64 @clock() #7
  %2 = trunc i64 %1 to i32
  ret i32 %2
}

; Function Attrs: nounwind
declare dso_local i64 @clock() #1

; Function Attrs: noinline nounwind optnone uwtable
define dso_local %struct.class_String* @rct_Char(i32 %0) #0 {
  %2 = alloca i32, align 4
  %3 = alloca i8*, align 8
  %4 = alloca %struct.class_String*, align 8
  store i32 %0, i32* %2, align 4
  %5 = call noalias align 16 i8* @malloc(i64 1) #7
  store i8* %5, i8** %3, align 8
  %6 = load i32, i32* %2, align 4
  %7 = trunc i32 %6 to i8
  %8 = load i8*, i8** %3, align 8
  %9 = getelementptr inbounds i8, i8* %8, i64 0
  store i8 %7, i8* %9, align 1
  %10 = call noalias align 16 i8* @malloc(i64 40) #7
  %11 = bitcast i8* %10 to %struct.class_String*
  store %struct.class_String* %11, %struct.class_String** %4, align 8
  %12 = load %struct.class_String*, %struct.class_String** %4, align 8
  call void @String_public_Constructor(%struct.class_String* %12)
  %13 = load %struct.class_String*, %struct.class_String** %4, align 8
  %14 = load i8*, i8** %3, align 8
  call void @String_public_Load(%struct.class_String* %13, i8* %14)
  %15 = load %struct.class_String*, %struct.class_String** %4, align 8
  %16 = bitcast %struct.class_String* %15 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* %16)
  %17 = load i8*, i8** %3, align 8
  call void @free(i8* %17) #7
  %18 = load %struct.class_String*, %struct.class_String** %4, align 8
  ret %struct.class_String* %18
}

attributes #0 = { noinline nounwind optnone uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { nounwind readonly willreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #4 = { argmemonly nofree nounwind willreturn }
attributes #5 = { noreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #6 = { nofree nosync nounwind readnone speculatable willreturn }
attributes #7 = { nounwind }
attributes #8 = { nounwind readonly willreturn }
attributes #9 = { noreturn }

!llvm.ident = !{!0, !0, !0}
!llvm.module.flags = !{!1, !2, !3}

!0 = !{!"Ubuntu clang version 13.0.0-2"}
!1 = !{i32 1, !"wchar_size", i32 4}
!2 = !{i32 7, !"uwtable", i32 1}
!3 = !{i32 7, !"frame-pointer", i32 2}
!4 = distinct !{!4, !5}
!5 = !{!"llvm.loop.mustprogress"}
!6 = distinct !{!6, !5}
!7 = distinct !{!7, !5}
