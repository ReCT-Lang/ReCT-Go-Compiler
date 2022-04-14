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
%struct.class_Thread = type { %struct.Any_vTable*, i32, i8* (i8*)*, i8*, i64 }
%union.pthread_attr_t = type { i64, [48 x i8] }

@.str = private unnamed_addr constant [59 x i8] c"\1B[36mARC \1B[0m- \1B[32mRegistered %s reference [%d] - %s\1B[0m\0A\00", align 1
@.str.1 = private unnamed_addr constant [61 x i8] c"\1B[36mARC \1B[0m- \1B[33mUnregistered %s reference [%d] - %s\1B[0m\0A\00", align 1
@.str.2 = private unnamed_addr constant [53 x i8] c"\1B[36mARC \1B[0m- \1B[31mDestroying %s instance - %s\1B[0m\0A\00", align 1
@.str.3 = private unnamed_addr constant [44 x i8] c"\1B[36mARC \1B[0m- \1B[0;35mWhat?? [%d] - %s\1B[0m\0A\00", align 1
@.str.4 = private unnamed_addr constant [4 x i8] c"Any\00", align 1
@Any_vTable_Const = dso_local constant %struct.Any_vTable { i8* null, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.4, i32 0, i32 0), void (i8*)* @Any_public_Die }, align 8
@.str.1.5 = private unnamed_addr constant [7 x i8] c"String\00", align 1
@String_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.1.5, i32 0, i32 0), void (i8*)* @String_public_Die }, align 8
@.str.2.6 = private unnamed_addr constant [42 x i8] c"Substring start-index cannot be negative!\00", align 1
@.str.3.7 = private unnamed_addr constant [37 x i8] c"Substring length cannot be negative!\00", align 1
@.str.4.8 = private unnamed_addr constant [24 x i8] c"Substring out of range!\00", align 1
@.str.5 = private unnamed_addr constant [4 x i8] c"Int\00", align 1
@Int_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.5, i32 0, i32 0), void (i8*)* @Int_public_Die }, align 8
@.str.6 = private unnamed_addr constant [6 x i8] c"Float\00", align 1
@Float_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.6, i32 0, i32 0), void (i8*)* @Float_public_Die }, align 8
@.str.7 = private unnamed_addr constant [5 x i8] c"Bool\00", align 1
@Bool_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.str.7, i32 0, i32 0), void (i8*)* @Bool_public_Die }, align 8
@.str.8 = private unnamed_addr constant [6 x i8] c"Array\00", align 1
@Array_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.8, i32 0, i32 0), void (i8*)* @Array_public_Die }, align 8
@.str.9 = private unnamed_addr constant [26 x i8] c"Array index out of range!\00", align 1
@.str.10 = private unnamed_addr constant [7 x i8] c"pArray\00", align 1
@pArray_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.10, i32 0, i32 0), void (i8*)* @pArray_public_Die }, align 8
@.str.11 = private unnamed_addr constant [7 x i8] c"Thread\00", align 1
@Thread_vTable_Const = dso_local constant %struct.Any_vTable { i8* bitcast (%struct.Any_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.11, i32 0, i32 0), void (i8*)* @Thread_public_Die }, align 8
@.str.12 = private unnamed_addr constant [45 x i8] c"%s[RUNTIME] %sEncountered Exception! %s'%s'\0A\00", align 1
@.str.1.13 = private unnamed_addr constant [8 x i8] c"\1B[1;31m\00", align 1
@.str.2.14 = private unnamed_addr constant [8 x i8] c"\1B[0;31m\00", align 1
@.str.3.15 = private unnamed_addr constant [19 x i8] c"%s[STACKTRACE] %s\0A\00", align 1
@.str.4.16 = private unnamed_addr constant [8 x i8] c"\1B[1;33m\00", align 1
@.str.5.17 = private unnamed_addr constant [8 x i8] c"\1B[0;33m\00", align 1
@.str.6.18 = private unnamed_addr constant [4 x i8] c".so\00", align 1
@.str.7.19 = private unnamed_addr constant [5 x i8] c".dll\00", align 1
@.str.8.20 = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@isCursorVisible = dso_local global i8 1, align 1
@.str.21 = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.1.22 = private unnamed_addr constant [3 x i8] c"%s\00", align 1
@.str.2.23 = private unnamed_addr constant [8 x i8] c"\1B[2J\1B[H\00", align 1
@.str.3.24 = private unnamed_addr constant [10 x i8] c"%c[%d;%df\00", align 1
@.str.4.25 = private unnamed_addr constant [8 x i8] c"\1B[?251]\00", align 1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  call void @free(i8* %24) #8
  br label %25

25:                                               ; preds = %15, %6, %5
  ret void
}

; Function Attrs: nounwind
declare void @free(i8*) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

declare i32 @printf(i8*, ...) #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  call void @free(i8* %43) #8
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Any_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  call void @free(i8* %13) #8
  br label %14

14:                                               ; preds = %10, %1
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Int_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Float_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Bool_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  br label %7, !llvm.loop !6

24:                                               ; preds = %7
  %25 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %26 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %25, i32 0, i32 2
  %27 = load %struct.class_Any**, %struct.class_Any*** %26, align 8
  %28 = bitcast %struct.class_Any** %27 to i8*
  call void @free(i8* %28) #8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  call void @free(i8* %8) #8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @String_public_Load(%struct.class_String* %0, i8* %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i8*, align 8
  %5 = alloca i32, align 4
  %6 = alloca i8*, align 8
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i8* %1, i8** %4, align 8
  %7 = load i8*, i8** %4, align 8
  %8 = call i64 @strlen(i8* %7) #9
  %9 = trunc i64 %8 to i32
  store i32 %9, i32* %5, align 4
  %10 = load i32, i32* %5, align 4
  %11 = add nsw i32 %10, 1
  %12 = sext i32 %11 to i64
  %13 = call noalias align 16 i8* @malloc(i64 %12) #8
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
  call void @free(i8* %26) #8
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
declare i64 @strlen(i8*) #3

; Function Attrs: nounwind
declare noalias align 16 i8* @malloc(i64) #1

; Function Attrs: argmemonly nofree nounwind willreturn
declare void @llvm.memcpy.p0i8.p0i8.i64(i8* noalias nocapture writeonly, i8* noalias nocapture readonly, i64, i1 immarg) #4

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @String_public_Resize(%struct.class_String* %0, i32 %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i32, align 4
  %5 = alloca i8*, align 8
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i32 %1, i32* %4, align 4
  %6 = load i32, i32* %4, align 4
  %7 = sext i32 %6 to i64
  %8 = call noalias align 16 i8* @malloc(i64 %7) #8
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
  call void @free(i8* %19) #8
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  %16 = call noalias align 16 i8* @malloc(i64 %15) #8
  store i8* %16, i8** %5, align 8
  %17 = load i8*, i8** %5, align 8
  %18 = load %struct.class_String*, %struct.class_String** %3, align 8
  %19 = getelementptr inbounds %struct.class_String, %struct.class_String* %18, i32 0, i32 2
  %20 = load i8*, i8** %19, align 8
  %21 = call i8* @strcpy(i8* %17, i8* %20) #8
  %22 = load i8*, i8** %5, align 8
  %23 = load %struct.class_String*, %struct.class_String** %4, align 8
  %24 = getelementptr inbounds %struct.class_String, %struct.class_String* %23, i32 0, i32 2
  %25 = load i8*, i8** %24, align 8
  %26 = call i8* @strcat(i8* %22, i8* %25) #8
  %27 = call noalias align 16 i8* @malloc(i64 40) #8
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
  call void @free(i8* %34) #8
  %35 = load %struct.class_String*, %struct.class_String** %6, align 8
  ret %struct.class_String* %35
}

; Function Attrs: nounwind
declare i8* @strcpy(i8*, i8*) #1

; Function Attrs: nounwind
declare i8* @strcat(i8*, i8*) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  %12 = call i32 @strcmp(i8* %8, i8* %11) #9
  store i32 %12, i32* %5, align 4
  %13 = load i32, i32* %5, align 4
  %14 = icmp eq i32 %13, 0
  ret i1 %14
}

; Function Attrs: nounwind readonly willreturn
declare i32 @strcmp(i8*, i8*) #3

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i8* @String_public_GetBuffer(%struct.class_String* %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = getelementptr inbounds %struct.class_String, %struct.class_String* %3, i32 0, i32 2
  %5 = load i8*, i8** %4, align 8
  ret i8* %5
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @String_public_GetLength(%struct.class_String* %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = getelementptr inbounds %struct.class_String, %struct.class_String* %3, i32 0, i32 3
  %5 = load i32, i32* %4, align 8
  ret i32 %5
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  br i1 %10, label %11, label %12

11:                                               ; preds = %3
  call void @exc_Throw(i8* getelementptr inbounds ([42 x i8], [42 x i8]* @.str.2.6, i64 0, i64 0))
  br label %45

12:                                               ; preds = %3
  %13 = load i32, i32* %6, align 4
  %14 = icmp slt i32 %13, 0
  br i1 %14, label %15, label %16

15:                                               ; preds = %12
  call void @exc_Throw(i8* getelementptr inbounds ([37 x i8], [37 x i8]* @.str.3.7, i64 0, i64 0))
  br label %44

16:                                               ; preds = %12
  %17 = load i32, i32* %5, align 4
  %18 = load i32, i32* %6, align 4
  %19 = add nsw i32 %17, %18
  %20 = load %struct.class_String*, %struct.class_String** %4, align 8
  %21 = getelementptr inbounds %struct.class_String, %struct.class_String* %20, i32 0, i32 3
  %22 = load i32, i32* %21, align 8
  %23 = icmp sgt i32 %19, %22
  br i1 %23, label %24, label %25

24:                                               ; preds = %16
  call void @exc_Throw(i8* getelementptr inbounds ([24 x i8], [24 x i8]* @.str.4.8, i64 0, i64 0))
  br label %43

25:                                               ; preds = %16
  %26 = load i32, i32* %6, align 4
  %27 = add nsw i32 %26, 1
  %28 = sext i32 %27 to i64
  %29 = call noalias align 16 i8* @malloc(i64 %28) #8
  store i8* %29, i8** %7, align 8
  %30 = load i8*, i8** %7, align 8
  %31 = load %struct.class_String*, %struct.class_String** %4, align 8
  %32 = getelementptr inbounds %struct.class_String, %struct.class_String* %31, i32 0, i32 2
  %33 = load i8*, i8** %32, align 8
  %34 = load i32, i32* %5, align 4
  %35 = sext i32 %34 to i64
  %36 = getelementptr inbounds i8, i8* %33, i64 %35
  %37 = load i32, i32* %6, align 4
  %38 = sext i32 %37 to i64
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 1 %30, i8* align 1 %36, i64 %38, i1 false)
  %39 = load i8*, i8** %7, align 8
  %40 = load i32, i32* %6, align 4
  %41 = sext i32 %40 to i64
  %42 = getelementptr inbounds i8, i8* %39, i64 %41
  store i8 0, i8* %42, align 1
  br label %43

43:                                               ; preds = %25, %24
  br label %44

44:                                               ; preds = %43, %15
  br label %45

45:                                               ; preds = %44, %11
  %46 = call noalias align 16 i8* @malloc(i64 40) #8
  %47 = bitcast i8* %46 to %struct.class_String*
  store %struct.class_String* %47, %struct.class_String** %8, align 8
  %48 = load %struct.class_String*, %struct.class_String** %8, align 8
  call void @String_public_Constructor(%struct.class_String* %48)
  %49 = load %struct.class_String*, %struct.class_String** %8, align 8
  %50 = load i8*, i8** %7, align 8
  call void @String_public_Load(%struct.class_String* %49, i8* %50)
  %51 = load %struct.class_String*, %struct.class_String** %8, align 8
  %52 = bitcast %struct.class_String* %51 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* %52)
  %53 = load i8*, i8** %7, align 8
  call void @free(i8* %53) #8
  %54 = load %struct.class_String*, %struct.class_String** %8, align 8
  ret %struct.class_String* %54
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  %19 = call noalias align 16 i8* @calloc(i64 %18, i64 8) #8
  %20 = bitcast i8* %19 to %struct.class_Any**
  %21 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %22 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %21, i32 0, i32 2
  store %struct.class_Any** %20, %struct.class_Any*** %22, align 8
  ret void
}

; Function Attrs: nounwind
declare noalias align 16 i8* @calloc(i64, i64) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_Any* @Array_public_GetElement(%struct.class_Array* %0, i32 %1) #0 {
  %3 = alloca %struct.class_Array*, align 8
  %4 = alloca i32, align 4
  store %struct.class_Array* %0, %struct.class_Array** %3, align 8
  store i32 %1, i32* %4, align 4
  %5 = load i32, i32* %4, align 4
  %6 = icmp slt i32 %5, 0
  br i1 %6, label %13, label %7

7:                                                ; preds = %2
  %8 = load i32, i32* %4, align 4
  %9 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %10 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %9, i32 0, i32 3
  %11 = load i32, i32* %10, align 8
  %12 = icmp sge i32 %8, %11
  br i1 %12, label %13, label %14

13:                                               ; preds = %7, %2
  call void @exc_Throw(i8* getelementptr inbounds ([26 x i8], [26 x i8]* @.str.9, i64 0, i64 0))
  br label %14

14:                                               ; preds = %13, %7
  %15 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %16 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %15, i32 0, i32 2
  %17 = load %struct.class_Any**, %struct.class_Any*** %16, align 8
  %18 = load i32, i32* %4, align 4
  %19 = sext i32 %18 to i64
  %20 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %17, i64 %19
  %21 = load %struct.class_Any*, %struct.class_Any** %20, align 8
  ret %struct.class_Any* %21
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  call void @exc_Throw(i8* getelementptr inbounds ([26 x i8], [26 x i8]* @.str.9, i64 0, i64 0))
  br label %16

16:                                               ; preds = %15, %9
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
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @Array_public_GetLength(%struct.class_Array* %0) #0 {
  %2 = alloca %struct.class_Array*, align 8
  store %struct.class_Array* %0, %struct.class_Array** %2, align 8
  %3 = load %struct.class_Array*, %struct.class_Array** %2, align 8
  %4 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %3, i32 0, i32 3
  %5 = load i32, i32* %4, align 8
  ret i32 %5
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  br i1 %14, label %15, label %84

15:                                               ; preds = %2
  %16 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %17 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %16, i32 0, i32 3
  %18 = load i32, i32* %17, align 8
  %19 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %20 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %19, i32 0, i32 5
  %21 = load i32, i32* %20, align 8
  %22 = add nsw i32 %18, %21
  store i32 %22, i32* %5, align 4
  %23 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %24 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %23, i32 0, i32 2
  %25 = load %struct.class_Any**, %struct.class_Any*** %24, align 8
  %26 = bitcast %struct.class_Any** %25 to i8*
  %27 = load i32, i32* %5, align 4
  %28 = sext i32 %27 to i64
  %29 = mul i64 8, %28
  %30 = call align 16 i8* @realloc(i8* %26, i64 %29) #8
  %31 = bitcast i8* %30 to %struct.class_Any**
  store %struct.class_Any** %31, %struct.class_Any*** %6, align 8
  %32 = load %struct.class_Any**, %struct.class_Any*** %6, align 8
  %33 = icmp eq %struct.class_Any** %32, null
  br i1 %33, label %34, label %55

34:                                               ; preds = %15
  %35 = load i32, i32* %5, align 4
  %36 = sext i32 %35 to i64
  %37 = mul i64 8, %36
  %38 = call noalias align 16 i8* @malloc(i64 %37) #8
  %39 = bitcast i8* %38 to %struct.class_Any**
  store %struct.class_Any** %39, %struct.class_Any*** %6, align 8
  %40 = load %struct.class_Any**, %struct.class_Any*** %6, align 8
  %41 = bitcast %struct.class_Any** %40 to i8*
  %42 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %43 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %42, i32 0, i32 2
  %44 = load %struct.class_Any**, %struct.class_Any*** %43, align 8
  %45 = bitcast %struct.class_Any** %44 to i8*
  %46 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %47 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %46, i32 0, i32 3
  %48 = load i32, i32* %47, align 8
  %49 = sext i32 %48 to i64
  %50 = mul i64 8, %49
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 8 %41, i8* align 8 %45, i64 %50, i1 false)
  %51 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %52 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %51, i32 0, i32 2
  %53 = load %struct.class_Any**, %struct.class_Any*** %52, align 8
  %54 = bitcast %struct.class_Any** %53 to i8*
  call void @free(i8* %54) #8
  br label %55

55:                                               ; preds = %34, %15
  %56 = load %struct.class_Any**, %struct.class_Any*** %6, align 8
  %57 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %58 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %57, i32 0, i32 2
  store %struct.class_Any** %56, %struct.class_Any*** %58, align 8
  store i32 0, i32* %7, align 4
  br label %59

59:                                               ; preds = %77, %55
  %60 = load i32, i32* %7, align 4
  %61 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %62 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %61, i32 0, i32 5
  %63 = load i32, i32* %62, align 8
  %64 = icmp slt i32 %60, %63
  br i1 %64, label %65, label %80

65:                                               ; preds = %59
  %66 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %67 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %66, i32 0, i32 2
  %68 = load %struct.class_Any**, %struct.class_Any*** %67, align 8
  %69 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %70 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %69, i32 0, i32 3
  %71 = load i32, i32* %70, align 8
  %72 = sext i32 %71 to i64
  %73 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %68, i64 %72
  %74 = load i32, i32* %7, align 4
  %75 = sext i32 %74 to i64
  %76 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %73, i64 %75
  store %struct.class_Any* null, %struct.class_Any** %76, align 8
  br label %77

77:                                               ; preds = %65
  %78 = load i32, i32* %7, align 4
  %79 = add nsw i32 %78, 1
  store i32 %79, i32* %7, align 4
  br label %59, !llvm.loop !8

80:                                               ; preds = %59
  %81 = load i32, i32* %5, align 4
  %82 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %83 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %82, i32 0, i32 4
  store i32 %81, i32* %83, align 4
  br label %84

84:                                               ; preds = %80, %2
  %85 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %86 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %85, i32 0, i32 3
  %87 = load i32, i32* %86, align 8
  %88 = add nsw i32 %87, 1
  store i32 %88, i32* %86, align 8
  %89 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %90 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %91 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %90, i32 0, i32 3
  %92 = load i32, i32* %91, align 8
  %93 = sub nsw i32 %92, 1
  %94 = load %struct.class_Any*, %struct.class_Any** %4, align 8
  call void @Array_public_SetElement(%struct.class_Array* %89, i32 %93, %struct.class_Any* %94)
  ret void
}

; Function Attrs: nounwind
declare align 16 i8* @realloc(i8*, i64) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  store i32 4, i32* %18, align 8
  %19 = load i32, i32* %6, align 4
  %20 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %21 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %20, i32 0, i32 6
  store i32 %19, i32* %21, align 4
  %22 = load i32, i32* %5, align 4
  %23 = sext i32 %22 to i64
  %24 = load i32, i32* %6, align 4
  %25 = sext i32 %24 to i64
  %26 = call noalias align 16 i8* @calloc(i64 %23, i64 %25) #8
  %27 = load %struct.class_pArray*, %struct.class_pArray** %4, align 8
  %28 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %27, i32 0, i32 2
  store i8* %26, i8** %28, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @pArray_public_GetLength(%struct.class_pArray* %0) #0 {
  %2 = alloca %struct.class_pArray*, align 8
  store %struct.class_pArray* %0, %struct.class_pArray** %2, align 8
  %3 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %4 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %3, i32 0, i32 3
  %5 = load i32, i32* %4, align 8
  ret i32 %5
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  br i1 %11, label %12, label %64

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
  %24 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %25 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %24, i32 0, i32 2
  %26 = load i8*, i8** %25, align 8
  %27 = load i32, i32* %3, align 4
  %28 = sext i32 %27 to i64
  %29 = call align 16 i8* @realloc(i8* %26, i64 %28) #8
  store i8* %29, i8** %4, align 8
  %30 = load i8*, i8** %4, align 8
  %31 = icmp eq i8* %30, null
  br i1 %31, label %32, label %51

32:                                               ; preds = %12
  %33 = load i32, i32* %3, align 4
  %34 = sext i32 %33 to i64
  %35 = call noalias align 16 i8* @malloc(i64 %34) #8
  store i8* %35, i8** %4, align 8
  %36 = load i8*, i8** %4, align 8
  %37 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %38 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %37, i32 0, i32 2
  %39 = load i8*, i8** %38, align 8
  %40 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %41 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %40, i32 0, i32 3
  %42 = load i32, i32* %41, align 8
  %43 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %44 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %43, i32 0, i32 6
  %45 = load i32, i32* %44, align 4
  %46 = mul nsw i32 %42, %45
  %47 = sext i32 %46 to i64
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* align 1 %36, i8* align 1 %39, i64 %47, i1 false)
  %48 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %49 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %48, i32 0, i32 2
  %50 = load i8*, i8** %49, align 8
  call void @free(i8* %50) #8
  br label %51

51:                                               ; preds = %32, %12
  %52 = load i8*, i8** %4, align 8
  %53 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %54 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %53, i32 0, i32 2
  store i8* %52, i8** %54, align 8
  %55 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %56 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %55, i32 0, i32 3
  %57 = load i32, i32* %56, align 8
  %58 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %59 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %58, i32 0, i32 5
  %60 = load i32, i32* %59, align 8
  %61 = add nsw i32 %57, %60
  %62 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %63 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %62, i32 0, i32 4
  store i32 %61, i32* %63, align 4
  br label %64

64:                                               ; preds = %51, %1
  %65 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %66 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %65, i32 0, i32 3
  %67 = load i32, i32* %66, align 8
  %68 = add nsw i32 %67, 1
  store i32 %68, i32* %66, align 8
  %69 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %70 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %69, i32 0, i32 2
  %71 = load i8*, i8** %70, align 8
  %72 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %73 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %72, i32 0, i32 3
  %74 = load i32, i32* %73, align 8
  %75 = sub nsw i32 %74, 1
  %76 = load %struct.class_pArray*, %struct.class_pArray** %2, align 8
  %77 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %76, i32 0, i32 6
  %78 = load i32, i32* %77, align 4
  %79 = mul nsw i32 %75, %78
  %80 = sext i32 %79 to i64
  %81 = getelementptr i8, i8* %71, i64 %80
  ret i8* %81
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i8* @pArray_public_GetElementPtr(%struct.class_pArray* %0, i32 %1) #0 {
  %3 = alloca %struct.class_pArray*, align 8
  %4 = alloca i32, align 4
  store %struct.class_pArray* %0, %struct.class_pArray** %3, align 8
  store i32 %1, i32* %4, align 4
  %5 = load i32, i32* %4, align 4
  %6 = icmp slt i32 %5, 0
  br i1 %6, label %13, label %7

7:                                                ; preds = %2
  %8 = load i32, i32* %4, align 4
  %9 = load %struct.class_pArray*, %struct.class_pArray** %3, align 8
  %10 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %9, i32 0, i32 3
  %11 = load i32, i32* %10, align 8
  %12 = icmp sge i32 %8, %11
  br i1 %12, label %13, label %14

13:                                               ; preds = %7, %2
  call void @exc_Throw(i8* getelementptr inbounds ([26 x i8], [26 x i8]* @.str.9, i64 0, i64 0))
  br label %14

14:                                               ; preds = %13, %7
  %15 = load %struct.class_pArray*, %struct.class_pArray** %3, align 8
  %16 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %15, i32 0, i32 2
  %17 = load i8*, i8** %16, align 8
  %18 = load i32, i32* %4, align 4
  %19 = load %struct.class_pArray*, %struct.class_pArray** %3, align 8
  %20 = getelementptr inbounds %struct.class_pArray, %struct.class_pArray* %19, i32 0, i32 6
  %21 = load i32, i32* %20, align 4
  %22 = mul nsw i32 %18, %21
  %23 = sext i32 %22 to i64
  %24 = getelementptr i8, i8* %17, i64 %23
  ret i8* %24
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Constructor(%struct.class_Thread* %0, i8* (i8*)* %1, i8* %2) #0 {
  %4 = alloca %struct.class_Thread*, align 8
  %5 = alloca i8* (i8*)*, align 8
  %6 = alloca i8*, align 8
  store %struct.class_Thread* %0, %struct.class_Thread** %4, align 8
  store i8* (i8*)* %1, i8* (i8*)** %5, align 8
  store i8* %2, i8** %6, align 8
  %7 = load %struct.class_Thread*, %struct.class_Thread** %4, align 8
  %8 = getelementptr inbounds %struct.class_Thread, %struct.class_Thread* %7, i32 0, i32 0
  store %struct.Any_vTable* @Thread_vTable_Const, %struct.Any_vTable** %8, align 8
  %9 = load %struct.class_Thread*, %struct.class_Thread** %4, align 8
  %10 = getelementptr inbounds %struct.class_Thread, %struct.class_Thread* %9, i32 0, i32 1
  store i32 0, i32* %10, align 8
  %11 = load i8* (i8*)*, i8* (i8*)** %5, align 8
  %12 = load %struct.class_Thread*, %struct.class_Thread** %4, align 8
  %13 = getelementptr inbounds %struct.class_Thread, %struct.class_Thread* %12, i32 0, i32 2
  store i8* (i8*)* %11, i8* (i8*)** %13, align 8
  %14 = load i8*, i8** %6, align 8
  %15 = load %struct.class_Thread*, %struct.class_Thread** %4, align 8
  %16 = getelementptr inbounds %struct.class_Thread, %struct.class_Thread* %15, i32 0, i32 3
  store i8* %14, i8** %16, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Start(%struct.class_Thread* %0) #0 {
  %2 = alloca %struct.class_Thread*, align 8
  store %struct.class_Thread* %0, %struct.class_Thread** %2, align 8
  %3 = load %struct.class_Thread*, %struct.class_Thread** %2, align 8
  %4 = getelementptr inbounds %struct.class_Thread, %struct.class_Thread* %3, i32 0, i32 4
  %5 = load %struct.class_Thread*, %struct.class_Thread** %2, align 8
  %6 = getelementptr inbounds %struct.class_Thread, %struct.class_Thread* %5, i32 0, i32 2
  %7 = load i8* (i8*)*, i8* (i8*)** %6, align 8
  %8 = load %struct.class_Thread*, %struct.class_Thread** %2, align 8
  %9 = getelementptr inbounds %struct.class_Thread, %struct.class_Thread* %8, i32 0, i32 3
  %10 = load i8*, i8** %9, align 8
  %11 = call i32 @pthread_create(i64* %4, %union.pthread_attr_t* null, i8* (i8*)* %7, i8* %10) #8
  ret void
}

; Function Attrs: nounwind
declare i32 @pthread_create(i64*, %union.pthread_attr_t*, i8* (i8*)*, i8*) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Join(%struct.class_Thread* %0) #0 {
  %2 = alloca %struct.class_Thread*, align 8
  store %struct.class_Thread* %0, %struct.class_Thread** %2, align 8
  %3 = load %struct.class_Thread*, %struct.class_Thread** %2, align 8
  %4 = getelementptr inbounds %struct.class_Thread, %struct.class_Thread* %3, i32 0, i32 4
  %5 = load i64, i64* %4, align 8
  %6 = call i32 @pthread_join(i64 %5, i8** null)
  ret void
}

declare i32 @pthread_join(i64, i8**) #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Kill(%struct.class_Thread* %0) #0 {
  %2 = alloca %struct.class_Thread*, align 8
  store %struct.class_Thread* %0, %struct.class_Thread** %2, align 8
  call void @pthread_exit(i8* null) #10
  unreachable
}

; Function Attrs: noreturn
declare void @pthread_exit(i8*) #5

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @exc_Throw(i8* %0) #0 {
  %2 = alloca i8*, align 8
  %3 = alloca [128 x i8*], align 16
  %4 = alloca i32, align 4
  %5 = alloca i8**, align 8
  %6 = alloca i32, align 4
  %7 = alloca i8*, align 8
  %8 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  %9 = load i8*, i8** %2, align 8
  %10 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([45 x i8], [45 x i8]* @.str.12, i64 0, i64 0), i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.1.13, i64 0, i64 0), i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.2.14, i64 0, i64 0), i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.1.13, i64 0, i64 0), i8* %9)
  %11 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([19 x i8], [19 x i8]* @.str.3.15, i64 0, i64 0), i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.4.16, i64 0, i64 0), i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.5.17, i64 0, i64 0))
  %12 = getelementptr inbounds [128 x i8*], [128 x i8*]* %3, i64 0, i64 0
  %13 = call i32 @backtrace(i8** %12, i32 128)
  store i32 %13, i32* %4, align 4
  %14 = getelementptr inbounds [128 x i8*], [128 x i8*]* %3, i64 0, i64 0
  %15 = load i32, i32* %4, align 4
  %16 = call i8** @backtrace_symbols(i8** %14, i32 %15) #8
  store i8** %16, i8*** %5, align 8
  store i32 1, i32* %6, align 4
  br label %17

17:                                               ; preds = %48, %1
  %18 = load i32, i32* %6, align 4
  %19 = load i32, i32* %4, align 4
  %20 = icmp slt i32 %18, %19
  br i1 %20, label %21, label %51

21:                                               ; preds = %17
  %22 = load i8**, i8*** %5, align 8
  %23 = load i32, i32* %6, align 4
  %24 = sext i32 %23 to i64
  %25 = getelementptr inbounds i8*, i8** %22, i64 %24
  %26 = load i8*, i8** %25, align 8
  %27 = call i8* @strstr(i8* %26, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.6.18, i64 0, i64 0)) #9
  store i8* %27, i8** %7, align 8
  %28 = load i8**, i8*** %5, align 8
  %29 = load i32, i32* %6, align 4
  %30 = sext i32 %29 to i64
  %31 = getelementptr inbounds i8*, i8** %28, i64 %30
  %32 = load i8*, i8** %31, align 8
  %33 = call i8* @strstr(i8* %32, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.str.7.19, i64 0, i64 0)) #9
  store i8* %33, i8** %8, align 8
  %34 = load i8*, i8** %7, align 8
  %35 = icmp ne i8* %34, null
  br i1 %35, label %36, label %37

36:                                               ; preds = %21
  br label %51

37:                                               ; preds = %21
  %38 = load i8*, i8** %8, align 8
  %39 = icmp ne i8* %38, null
  br i1 %39, label %40, label %41

40:                                               ; preds = %37
  br label %51

41:                                               ; preds = %37
  %42 = load i8**, i8*** %5, align 8
  %43 = load i32, i32* %6, align 4
  %44 = sext i32 %43 to i64
  %45 = getelementptr inbounds i8*, i8** %42, i64 %44
  %46 = load i8*, i8** %45, align 8
  %47 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.8.20, i64 0, i64 0), i8* %46)
  br label %48

48:                                               ; preds = %41
  %49 = load i32, i32* %6, align 4
  %50 = add nsw i32 %49, 1
  store i32 %50, i32* %6, align 4
  br label %17, !llvm.loop !9

51:                                               ; preds = %40, %36, %17
  %52 = load i8**, i8*** %5, align 8
  %53 = bitcast i8** %52 to i8*
  call void @free(i8* %53) #8
  call void @exit(i32 -1) #11
  unreachable
}

declare i32 @backtrace(i8**, i32) #2

; Function Attrs: nounwind
declare i8** @backtrace_symbols(i8**, i32) #1

; Function Attrs: nounwind readonly willreturn
declare i8* @strstr(i8*, i8*) #3

; Function Attrs: noreturn nounwind
declare void @exit(i32) #6

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @rct_Print(%struct.class_String* %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = getelementptr inbounds %struct.class_String, %struct.class_String* %3, i32 0, i32 2
  %5 = load i8*, i8** %4, align 8
  %6 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.21, i64 0, i64 0), i8* %5)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @rct_Write(%struct.class_String* %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = getelementptr inbounds %struct.class_String, %struct.class_String* %3, i32 0, i32 2
  %5 = load i8*, i8** %4, align 8
  %6 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.str.1.22, i64 0, i64 0), i8* %5)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @rct_Input() #0 {
  %1 = alloca i8*, align 8
  %2 = alloca i8*, align 8
  %3 = alloca i32, align 4
  %4 = alloca %struct.class_String*, align 8
  %5 = call noalias align 16 i8* @malloc(i64 1042) #8
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
  %31 = call align 16 i8* @realloc(i8* %25, i64 %30) #8
  store i8* %31, i8** %2, align 8
  %32 = icmp eq i8* %31, null
  br i1 %32, label %33, label %35

33:                                               ; preds = %24
  %34 = load i8*, i8** %1, align 8
  call void @free(i8* %34) #8
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
  br label %6, !llvm.loop !10

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
  %50 = call noalias align 16 i8* @malloc(i64 40) #8
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
  call void @free(i8* %60) #8
  br label %61

61:                                               ; preds = %59, %49
  %62 = load %struct.class_String*, %struct.class_String** %4, align 8
  ret %struct.class_String* %62
}

declare i32 @getchar() #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @rct_Clear() #0 {
  %1 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.2.23, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @rct_SetCursor(i32 %0, i32 %1) #0 {
  %3 = alloca i32, align 4
  %4 = alloca i32, align 4
  store i32 %0, i32* %3, align 4
  store i32 %1, i32* %4, align 4
  %5 = load i32, i32* %4, align 4
  %6 = load i32, i32* %3, align 4
  %7 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([10 x i8], [10 x i8]* @.str.3.24, i64 0, i64 0), i32 27, i32 %5, i32 %6)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  %10 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.4.25, i64 0, i64 0))
  br label %13

11:                                               ; preds = %1
  %12 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.4.25, i64 0, i64 0))
  br label %13

13:                                               ; preds = %11, %9
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local zeroext i1 @rct_GetCursorVisible() #0 {
  %1 = load i8, i8* @isCursorVisible, align 1
  %2 = trunc i8 %1 to i1
  ret i1 %2
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @rct_Random(i32 %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = call i32 @rand() #8
  %4 = load i32, i32* %2, align 4
  %5 = srem i32 %3, %4
  ret i32 %5
}

; Function Attrs: nounwind
declare i32 @rand() #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @rct_Sleep(i32 %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = mul nsw i32 %3, 1000
  %5 = call i32 @usleep(i32 %4)
  ret void
}

declare i32 @usleep(i32) #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @rct_Sqrt(i32 %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = sitofp i32 %3 to double
  %5 = call double @sqrt(double %4) #8
  %6 = call double @llvm.floor.f64(double %5)
  %7 = fptosi double %6 to i32
  ret i32 %7
}

; Function Attrs: nounwind
declare double @sqrt(double) #1

; Function Attrs: nofree nosync nounwind readnone speculatable willreturn
declare double @llvm.floor.f64(double) #7

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @rct_Now() #0 {
  %1 = call i64 @clock() #8
  %2 = trunc i64 %1 to i32
  ret i32 %2
}

; Function Attrs: nounwind
declare i64 @clock() #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @rct_Char(i32 %0) #0 {
  %2 = alloca i32, align 4
  %3 = alloca i8*, align 8
  %4 = alloca %struct.class_String*, align 8
  store i32 %0, i32* %2, align 4
  %5 = call noalias align 16 i8* @malloc(i64 1) #8
  store i8* %5, i8** %3, align 8
  %6 = load i32, i32* %2, align 4
  %7 = trunc i32 %6 to i8
  %8 = load i8*, i8** %3, align 8
  %9 = getelementptr inbounds i8, i8* %8, i64 0
  store i8 %7, i8* %9, align 1
  %10 = call noalias align 16 i8* @malloc(i64 40) #8
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
  call void @free(i8* %17) #8
  %18 = load %struct.class_String*, %struct.class_String** %4, align 8
  ret %struct.class_String* %18
}

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { nounwind readonly willreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #4 = { argmemonly nofree nounwind willreturn }
attributes #5 = { noreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #6 = { noreturn nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #7 = { nofree nosync nounwind readnone speculatable willreturn }
attributes #8 = { nounwind }
attributes #9 = { nounwind readonly willreturn }
attributes #10 = { noreturn }
attributes #11 = { noreturn nounwind }

!llvm.ident = !{!0, !0, !0, !0}
!llvm.module.flags = !{!1, !2, !3, !4, !5}

!0 = !{!"clang version 13.0.0"}
!1 = !{i32 1, !"wchar_size", i32 4}
!2 = !{i32 7, !"PIC Level", i32 2}
!3 = !{i32 7, !"PIE Level", i32 2}
!4 = !{i32 7, !"uwtable", i32 1}
!5 = !{i32 7, !"frame-pointer", i32 2}
!6 = distinct !{!6, !7}
!7 = !{!"llvm.loop.mustprogress"}
!8 = distinct !{!8, !7}
!9 = distinct !{!9, !7}
!10 = distinct !{!10, !7}
