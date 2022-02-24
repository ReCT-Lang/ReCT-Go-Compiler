; ModuleID = './objects.c'
source_filename = "./objects.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Int_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Float_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Bool_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Array_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }
%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.class_Int = type { %struct.Int_vTable*, i32, i32 }
%struct.class_Float = type { %struct.Float_vTable*, i32, float }
%struct.class_Bool = type { %struct.Bool_vTable*, i32, i8 }
%struct.class_Array = type { %struct.Array_vTable*, i32, %struct.class_Any**, i32 }

@.str = private unnamed_addr constant [4 x i8] c"Any\00", align 1
@Any_vTable_Const = dso_local constant %struct.Any_vTable { i8* null, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str, i32 0, i32 0), void (i8*)* @Any_public_Die }, align 8
@.str.1 = private unnamed_addr constant [7 x i8] c"String\00", align 1
@String_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.1, i32 0, i32 0), void (i8*)* @String_public_Die }, align 8
@.str.2 = private unnamed_addr constant [44 x i8] c"https://www.youtube.com/watch?v=dQw4w9WgXcQ\00", align 1
@.str.3 = private unnamed_addr constant [4 x i8] c"Int\00", align 1
@Int_vTable_Const = dso_local constant %struct.Int_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.3, i32 0, i32 0), void (i8*)* @Int_public_Die }, align 8
@.str.4 = private unnamed_addr constant [6 x i8] c"Float\00", align 1
@Float_vTable_Const = dso_local constant %struct.Float_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.4, i32 0, i32 0), void (i8*)* @Float_public_Die }, align 8
@.str.5 = private unnamed_addr constant [5 x i8] c"Bool\00", align 1
@Bool_vTable_Const = dso_local constant %struct.Bool_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.str.5, i32 0, i32 0), void (i8*)* @Bool_public_Die }, align 8
@.str.6 = private unnamed_addr constant [6 x i8] c"Array\00", align 1
@Array_vTable_Const = dso_local constant %struct.Array_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.6, i32 0, i32 0), void (i8*)* @Array_public_Die }, align 8

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Any_public_Die(i8* %0) #0 {
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
  call void @free(i8* %13) #5
  br label %14

14:                                               ; preds = %10, %1
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

; Function Attrs: nounwind
declare void @free(i8*) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @String_public_Load(%struct.class_String* %0, i8* %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i8*, align 8
  %5 = alloca i32, align 4
  %6 = alloca i8*, align 8
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i8* %1, i8** %4, align 8
  %7 = load i8*, i8** %4, align 8
  %8 = call i64 @strlen(i8* %7) #6
  %9 = trunc i64 %8 to i32
  store i32 %9, i32* %5, align 4
  %10 = load i32, i32* %5, align 4
  %11 = add nsw i32 %10, 1
  %12 = sext i32 %11 to i64
  %13 = call noalias align 16 i8* @malloc(i64 %12) #5
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
  call void @free(i8* %26) #5
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
declare i64 @strlen(i8*) #2

; Function Attrs: nounwind
declare noalias align 16 i8* @malloc(i64) #1

; Function Attrs: argmemonly nofree nounwind willreturn
declare void @llvm.memcpy.p0i8.p0i8.i64(i8* noalias nocapture writeonly, i8* noalias nocapture readonly, i64, i1 immarg) #3

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @String_public_Resize(%struct.class_String* %0, i32 %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i32, align 4
  %5 = alloca i8*, align 8
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i32 %1, i32* %4, align 4
  %6 = load i32, i32* %4, align 4
  %7 = sext i32 %6 to i64
  %8 = call noalias align 16 i8* @malloc(i64 %7) #5
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
  call void @free(i8* %19) #5
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
  %16 = call noalias align 16 i8* @malloc(i64 %15) #5
  store i8* %16, i8** %5, align 8
  %17 = load i8*, i8** %5, align 8
  %18 = load %struct.class_String*, %struct.class_String** %3, align 8
  %19 = getelementptr inbounds %struct.class_String, %struct.class_String* %18, i32 0, i32 2
  %20 = load i8*, i8** %19, align 8
  %21 = call i8* @strcpy(i8* %17, i8* %20) #5
  %22 = load i8*, i8** %5, align 8
  %23 = load %struct.class_String*, %struct.class_String** %4, align 8
  %24 = getelementptr inbounds %struct.class_String, %struct.class_String* %23, i32 0, i32 2
  %25 = load i8*, i8** %24, align 8
  %26 = call i8* @strcat(i8* %22, i8* %25) #5
  %27 = call noalias align 16 i8* @malloc(i64 40) #5
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
  call void @free(i8* %34) #5
  %35 = load %struct.class_String*, %struct.class_String** %6, align 8
  ret %struct.class_String* %35
}

; Function Attrs: nounwind
declare i8* @strcpy(i8*, i8*) #1

; Function Attrs: nounwind
declare i8* @strcat(i8*, i8*) #1

declare void @arc_RegisterReference(%struct.class_Any*) #4

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
  %12 = call i32 @strcmp(i8* %8, i8* %11) #6
  store i32 %12, i32* %5, align 4
  %13 = load i32, i32* %5, align 4
  %14 = icmp eq i32 %13, 0
  ret i1 %14
}

; Function Attrs: nounwind readonly willreturn
declare i32 @strcmp(i8*, i8*) #2

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
  store i8* getelementptr inbounds ([44 x i8], [44 x i8]* @.str.2, i64 0, i64 0), i8** %7, align 8
  br label %41

23:                                               ; preds = %14
  %24 = load i32, i32* %6, align 4
  %25 = add nsw i32 %24, 1
  %26 = sext i32 %25 to i64
  %27 = call noalias align 16 i8* @malloc(i64 %26) #5
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
  %42 = call noalias align 16 i8* @malloc(i64 40) #5
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
  call void @free(i8* %49) #5
  %50 = load %struct.class_String*, %struct.class_String** %8, align 8
  ret %struct.class_String* %50
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Int_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Int_public_Constructor(%struct.class_Int* %0, i32 %1) #0 {
  %3 = alloca %struct.class_Int*, align 8
  %4 = alloca i32, align 4
  store %struct.class_Int* %0, %struct.class_Int** %3, align 8
  store i32 %1, i32* %4, align 4
  %5 = load %struct.class_Int*, %struct.class_Int** %3, align 8
  %6 = getelementptr inbounds %struct.class_Int, %struct.class_Int* %5, i32 0, i32 0
  store %struct.Int_vTable* @Int_vTable_Const, %struct.Int_vTable** %6, align 8
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
  %2 = alloca %struct.class_Int*, align 8
  store %struct.class_Int* %0, %struct.class_Int** %2, align 8
  %3 = load %struct.class_Int*, %struct.class_Int** %2, align 8
  %4 = getelementptr inbounds %struct.class_Int, %struct.class_Int* %3, i32 0, i32 2
  %5 = load i32, i32* %4, align 4
  ret i32 %5
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Float_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Float_public_Constructor(%struct.class_Float* %0, float %1) #0 {
  %3 = alloca %struct.class_Float*, align 8
  %4 = alloca float, align 4
  store %struct.class_Float* %0, %struct.class_Float** %3, align 8
  store float %1, float* %4, align 4
  %5 = load %struct.class_Float*, %struct.class_Float** %3, align 8
  %6 = getelementptr inbounds %struct.class_Float, %struct.class_Float* %5, i32 0, i32 0
  store %struct.Float_vTable* @Float_vTable_Const, %struct.Float_vTable** %6, align 8
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
  %2 = alloca %struct.class_Float*, align 8
  store %struct.class_Float* %0, %struct.class_Float** %2, align 8
  %3 = load %struct.class_Float*, %struct.class_Float** %2, align 8
  %4 = getelementptr inbounds %struct.class_Float, %struct.class_Float* %3, i32 0, i32 2
  %5 = load float, float* %4, align 4
  ret float %5
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Bool_public_Die(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  ret void
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
  store %struct.Bool_vTable* @Bool_vTable_Const, %struct.Bool_vTable** %7, align 8
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
  %2 = alloca %struct.class_Bool*, align 8
  store %struct.class_Bool* %0, %struct.class_Bool** %2, align 8
  %3 = load %struct.class_Bool*, %struct.class_Bool** %2, align 8
  %4 = getelementptr inbounds %struct.class_Bool, %struct.class_Bool* %3, i32 0, i32 2
  %5 = load i8, i8* %4, align 4
  %6 = trunc i8 %5 to i1
  ret i1 %6
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
  call void @free(i8* %28) #5
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Array_public_Constructor(%struct.class_Array* %0, i32 %1) #0 {
  %3 = alloca %struct.class_Array*, align 8
  %4 = alloca i32, align 4
  store %struct.class_Array* %0, %struct.class_Array** %3, align 8
  store i32 %1, i32* %4, align 4
  %5 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %6 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %5, i32 0, i32 0
  store %struct.Array_vTable* @Array_vTable_Const, %struct.Array_vTable** %6, align 8
  %7 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %8 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %7, i32 0, i32 1
  store i32 0, i32* %8, align 8
  %9 = load i32, i32* %4, align 4
  %10 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %11 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %10, i32 0, i32 3
  store i32 %9, i32* %11, align 8
  %12 = load i32, i32* %4, align 4
  %13 = sext i32 %12 to i64
  %14 = mul i64 8, %13
  %15 = call noalias align 16 i8* @malloc(i64 %14) #5
  %16 = bitcast i8* %15 to %struct.class_Any**
  %17 = load %struct.class_Array*, %struct.class_Array** %3, align 8
  %18 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %17, i32 0, i32 2
  store %struct.class_Any** %16, %struct.class_Any*** %18, align 8
  ret void
}

declare void @arc_UnregisterReference(%struct.class_Any*) #4

; Function Attrs: noinline nounwind optnone sspstrong uwtable
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
  br label %25

16:                                               ; preds = %9
  %17 = load %struct.class_Any*, %struct.class_Any** %6, align 8
  call void @arc_RegisterReference(%struct.class_Any* %17)
  %18 = load %struct.class_Any*, %struct.class_Any** %6, align 8
  %19 = load %struct.class_Array*, %struct.class_Array** %4, align 8
  %20 = getelementptr inbounds %struct.class_Array, %struct.class_Array* %19, i32 0, i32 2
  %21 = load %struct.class_Any**, %struct.class_Any*** %20, align 8
  %22 = load i32, i32* %5, align 4
  %23 = sext i32 %22 to i64
  %24 = getelementptr inbounds %struct.class_Any*, %struct.class_Any** %21, i64 %23
  store %struct.class_Any* %18, %struct.class_Any** %24, align 8
  br label %25

25:                                               ; preds = %16, %15
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

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { nounwind readonly willreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { argmemonly nofree nounwind willreturn }
attributes #4 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #5 = { nounwind }
attributes #6 = { nounwind readonly willreturn }

!llvm.module.flags = !{!0, !1, !2, !3, !4}
!llvm.ident = !{!5}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{i32 7, !"PIC Level", i32 2}
!2 = !{i32 7, !"PIE Level", i32 2}
!3 = !{i32 7, !"uwtable", i32 1}
!4 = !{i32 7, !"frame-pointer", i32 2}
!5 = !{!"clang version 13.0.0"}
!6 = distinct !{!6, !7}
!7 = !{!"llvm.loop.mustprogress"}
