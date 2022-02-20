; ModuleID = './objects.c'
source_filename = "./objects.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Int_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Float_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Bool_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }
%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.class_Int = type { %struct.Int_vTable*, i32, i32 }
%struct.class_Float = type { %struct.Float_vTable*, i32, float }
%struct.class_Bool = type { %struct.Bool_vTable*, i32, i8 }

@.str = private unnamed_addr constant [4 x i8] c"Any\00", align 1
@Any_vTable_Const = dso_local constant %struct.Any_vTable { i8* null, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str, i32 0, i32 0), void (i8*)* @Any_public_Die }, align 8
@.str.1 = private unnamed_addr constant [7 x i8] c"String\00", align 1
@String_vTable_Const = dso_local constant %struct.String_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.1, i32 0, i32 0), void (i8*)* @String_public_Die }, align 8
@.str.2 = private unnamed_addr constant [4 x i8] c"Int\00", align 1
@Int_vTable_Const = dso_local constant %struct.Int_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.2, i32 0, i32 0), void (i8*)* @Int_public_Die }, align 8
@.str.3 = private unnamed_addr constant [6 x i8] c"Float\00", align 1
@Float_vTable_Const = dso_local constant %struct.Float_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.3, i32 0, i32 0), void (i8*)* @Float_public_Die }, align 8
@.str.4 = private unnamed_addr constant [5 x i8] c"Bool\00", align 1
@Bool_vTable_Const = dso_local constant %struct.Bool_vTable { %struct.Any_vTable* @Any_vTable_Const, i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.str.4, i32 0, i32 0), void (i8*)* @Bool_public_Die }, align 8

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
  call void @free(i8* %13) #4
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
  %8 = call i64 @strlen(i8* %7) #5
  %9 = trunc i64 %8 to i32
  store i32 %9, i32* %5, align 4
  %10 = load i32, i32* %5, align 4
  %11 = add nsw i32 %10, 1
  %12 = sext i32 %11 to i64
  %13 = call noalias align 16 i8* @malloc(i64 %12) #4
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
  call void @free(i8* %21) #4
  %22 = load i8*, i8** %6, align 8
  %23 = load %struct.class_String*, %struct.class_String** %3, align 8
  %24 = getelementptr inbounds %struct.class_String, %struct.class_String* %23, i32 0, i32 2
  store i8* %22, i8** %24, align 8
  %25 = load i32, i32* %5, align 4
  %26 = load %struct.class_String*, %struct.class_String** %3, align 8
  %27 = getelementptr inbounds %struct.class_String, %struct.class_String* %26, i32 0, i32 4
  store i32 %25, i32* %27, align 4
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
  %8 = call noalias align 16 i8* @malloc(i64 %7) #4
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
  call void @free(i8* %19) #4
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

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { nounwind readonly willreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { argmemonly nofree nounwind willreturn }
attributes #4 = { nounwind }
attributes #5 = { nounwind readonly willreturn }

!llvm.module.flags = !{!0, !1, !2, !3, !4}
!llvm.ident = !{!5}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{i32 7, !"PIC Level", i32 2}
!2 = !{i32 7, !"PIE Level", i32 2}
!3 = !{i32 7, !"uwtable", i32 1}
!4 = !{i32 7, !"frame-pointer", i32 2}
!5 = !{!"clang version 13.0.0"}
