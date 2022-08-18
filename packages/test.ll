; ModuleID = './test.c'
source_filename = "./test.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.Thing_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.class_Thing = type { %struct.Thing_vTable*, i32, %struct.class_String* }
%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }

@.str = private unnamed_addr constant [11 x i8] c"someString\00", align 1
@Thing_Fields_Const = dso_local global [1 x i8*] [i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.str, i32 0, i32 0)], align 8
@.str.1 = private unnamed_addr constant [6 x i8] c"Thing\00", align 1
@Thing_vTable_Const = dso_local constant %struct.Thing_vTable { %struct.Any_vTable* null, i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.1, i32 0, i32 0), void (i8*)* @Thing_public_Die }, align 8
@.str.2 = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.3 = private unnamed_addr constant [12 x i8] c"cool string\00", align 1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thing_public_Die(i8* noundef %0) #0 {
  %2 = alloca i8*, align 8
  %3 = alloca %struct.class_Thing*, align 8
  store i8* %0, i8** %2, align 8
  %4 = load i8*, i8** %2, align 8
  %5 = bitcast i8* %4 to %struct.class_Thing*
  store %struct.class_Thing* %5, %struct.class_Thing** %3, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thing_public_Constructor(%struct.class_Thing* noundef %0, %struct.class_String* noundef %1) #0 {
  %3 = alloca %struct.class_Thing*, align 8
  %4 = alloca %struct.class_String*, align 8
  store %struct.class_Thing* %0, %struct.class_Thing** %3, align 8
  store %struct.class_String* %1, %struct.class_String** %4, align 8
  %5 = load %struct.class_Thing*, %struct.class_Thing** %3, align 8
  %6 = getelementptr inbounds %struct.class_Thing, %struct.class_Thing* %5, i32 0, i32 0
  store %struct.Thing_vTable* @Thing_vTable_Const, %struct.Thing_vTable** %6, align 8
  %7 = load %struct.class_Thing*, %struct.class_Thing** %3, align 8
  %8 = getelementptr inbounds %struct.class_Thing, %struct.class_Thing* %7, i32 0, i32 1
  store i32 0, i32* %8, align 8
  %9 = load %struct.class_String*, %struct.class_String** %4, align 8
  %10 = load %struct.class_Thing*, %struct.class_Thing** %3, align 8
  %11 = getelementptr inbounds %struct.class_Thing, %struct.class_Thing* %10, i32 0, i32 2
  store %struct.class_String* %9, %struct.class_String** %11, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thing_public_Output(%struct.class_Thing* noundef %0) #0 {
  %2 = alloca %struct.class_Thing*, align 8
  store %struct.class_Thing* %0, %struct.class_Thing** %2, align 8
  %3 = load %struct.class_Thing*, %struct.class_Thing** %2, align 8
  %4 = getelementptr inbounds %struct.class_Thing, %struct.class_Thing* %3, i32 0, i32 2
  %5 = load %struct.class_String*, %struct.class_String** %4, align 8
  %6 = getelementptr inbounds %struct.class_String, %struct.class_String* %5, i32 0, i32 2
  %7 = load i8*, i8** %6, align 8
  %8 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([4 x i8], [4 x i8]* @.str.2, i64 0, i64 0), i8* noundef %7)
  ret void
}

declare i32 @printf(i8* noundef, ...) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thing_public_ChangeString(%struct.class_Thing* noundef %0, %struct.class_String* noundef %1) #0 {
  %3 = alloca %struct.class_Thing*, align 8
  %4 = alloca %struct.class_String*, align 8
  store %struct.class_Thing* %0, %struct.class_Thing** %3, align 8
  store %struct.class_String* %1, %struct.class_String** %4, align 8
  %5 = load %struct.class_String*, %struct.class_String** %4, align 8
  %6 = load %struct.class_Thing*, %struct.class_Thing** %3, align 8
  %7 = getelementptr inbounds %struct.class_Thing, %struct.class_Thing* %6, i32 0, i32 2
  store %struct.class_String* %5, %struct.class_String** %7, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @Thing_public_GetString(%struct.class_Thing* noundef %0) #0 {
  %2 = alloca %struct.class_Thing*, align 8
  store %struct.class_Thing* %0, %struct.class_Thing** %2, align 8
  %3 = load %struct.class_Thing*, %struct.class_Thing** %2, align 8
  %4 = getelementptr inbounds %struct.class_Thing, %struct.class_Thing* %3, i32 0, i32 2
  %5 = load %struct.class_String*, %struct.class_String** %4, align 8
  ret %struct.class_String* %5
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @test_PrintArr(%struct.class_Thing* noundef %0) #0 {
  %2 = alloca %struct.class_Thing*, align 8
  store %struct.class_Thing* %0, %struct.class_Thing** %2, align 8
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @test_GetString() #0 {
  %1 = alloca %struct.class_String*, align 8
  %2 = call noalias i8* @malloc(i64 noundef 40) #3
  %3 = bitcast i8* %2 to %struct.class_String*
  store %struct.class_String* %3, %struct.class_String** %1, align 8
  %4 = load %struct.class_String*, %struct.class_String** %1, align 8
  call void @String_public_Constructor(%struct.class_String* noundef %4)
  %5 = load %struct.class_String*, %struct.class_String** %1, align 8
  %6 = bitcast %struct.class_String* %5 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* noundef %6)
  %7 = load %struct.class_String*, %struct.class_String** %1, align 8
  call void @String_public_Load(%struct.class_String* noundef %7, i8* noundef getelementptr inbounds ([12 x i8], [12 x i8]* @.str.3, i64 0, i64 0))
  %8 = load %struct.class_String*, %struct.class_String** %1, align 8
  ret %struct.class_String* %8
}

; Function Attrs: nounwind
declare noalias i8* @malloc(i64 noundef) #2

declare void @String_public_Constructor(%struct.class_String* noundef) #1

declare void @arc_RegisterReference(%struct.class_Any* noundef) #1

declare void @String_public_Load(%struct.class_String* noundef, i8* noundef) #1

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { nounwind }

!llvm.module.flags = !{!0, !1, !2, !3, !4}
!llvm.ident = !{!5}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{i32 7, !"PIC Level", i32 2}
!2 = !{i32 7, !"PIE Level", i32 2}
!3 = !{i32 7, !"uwtable", i32 1}
!4 = !{i32 7, !"frame-pointer", i32 2}
!5 = !{!"clang version 14.0.6"}
