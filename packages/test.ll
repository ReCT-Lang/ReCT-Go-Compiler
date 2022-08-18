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
%struct.class_Array_String = type { %struct.Array_vTable*, i32, %struct.class_Any**, i32, i32, i32 }
%struct.Array_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_Array = type { %struct.Array_vTable*, i32, %struct.class_Any**, i32, i32, i32 }
%"struct.class_Array_T_array_$b$string$s$$e$" = type { %struct.Array_vTable*, i32, %struct.class_Any**, i32, i32, i32 }

@.str = private unnamed_addr constant [11 x i8] c"someString\00", align 1
@Thing_Fields_Const = dso_local global [1 x i8*] [i8* getelementptr inbounds ([11 x i8], [11 x i8]* @.str, i32 0, i32 0)], align 8
@.str.1 = private unnamed_addr constant [6 x i8] c"Thing\00", align 1
@Thing_vTable_Const = dso_local constant %struct.Thing_vTable { %struct.Any_vTable* null, i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.1, i32 0, i32 0), void (i8*)* @Thing_public_Die }, align 8
@.str.2 = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.3 = private unnamed_addr constant [21 x i8] c"cool string business\00", align 1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thing_public_Die(i8* noundef %0) #0 {
  %2 = alloca i8*, align 8
  %3 = alloca %struct.class_Thing*, align 8
  store i8* %0, i8** %2, align 8
  %4 = load i8*, i8** %2, align 8
  %5 = bitcast i8* %4 to %struct.class_Thing*
  store %struct.class_Thing* %5, %struct.class_Thing** %3, align 8
  %6 = load %struct.class_Thing*, %struct.class_Thing** %3, align 8
  %7 = getelementptr inbounds %struct.class_Thing, %struct.class_Thing* %6, i32 0, i32 2
  %8 = load %struct.class_String*, %struct.class_String** %7, align 8
  %9 = bitcast %struct.class_String* %8 to %struct.class_Any*
  call void @arc_UnregisterReference(%struct.class_Any* noundef %9)
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

declare void @arc_UnregisterReference(%struct.class_Any* noundef) #1

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
define dso_local void @test_PrintArr(%struct.class_Array_String* noundef %0) #0 {
  %2 = alloca %struct.class_Array_String*, align 8
  %3 = alloca i32, align 4
  store %struct.class_Array_String* %0, %struct.class_Array_String** %2, align 8
  store i32 0, i32* %3, align 4
  br label %4

4:                                                ; preds = %19, %1
  %5 = load i32, i32* %3, align 4
  %6 = load %struct.class_Array_String*, %struct.class_Array_String** %2, align 8
  %7 = getelementptr inbounds %struct.class_Array_String, %struct.class_Array_String* %6, i32 0, i32 3
  %8 = load i32, i32* %7, align 8
  %9 = icmp slt i32 %5, %8
  br i1 %9, label %10, label %22

10:                                               ; preds = %4
  %11 = load %struct.class_Array_String*, %struct.class_Array_String** %2, align 8
  %12 = bitcast %struct.class_Array_String* %11 to %struct.class_Array*
  %13 = load i32, i32* %3, align 4
  %14 = call %struct.class_Any* @Array_public_GetElement(%struct.class_Array* noundef %12, i32 noundef %13)
  %15 = bitcast %struct.class_Any* %14 to %struct.class_String*
  %16 = getelementptr inbounds %struct.class_String, %struct.class_String* %15, i32 0, i32 2
  %17 = load i8*, i8** %16, align 8
  %18 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([4 x i8], [4 x i8]* @.str.2, i64 0, i64 0), i8* noundef %17)
  br label %19

19:                                               ; preds = %10
  %20 = load i32, i32* %3, align 4
  %21 = add nsw i32 %20, 1
  store i32 %21, i32* %3, align 4
  br label %4, !llvm.loop !6

22:                                               ; preds = %4
  ret void
}

declare %struct.class_Any* @Array_public_GetElement(%struct.class_Array* noundef, i32 noundef) #1

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
  call void @String_public_Load(%struct.class_String* noundef %7, i8* noundef getelementptr inbounds ([21 x i8], [21 x i8]* @.str.3, i64 0, i64 0))
  %8 = load %struct.class_String*, %struct.class_String** %1, align 8
  ret %struct.class_String* %8
}

; Function Attrs: nounwind
declare noalias i8* @malloc(i64 noundef) #2

declare void @String_public_Constructor(%struct.class_String* noundef) #1

declare void @arc_RegisterReference(%struct.class_Any* noundef) #1

declare void @String_public_Load(%struct.class_String* noundef, i8* noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_Array_String* @test_GetStringArray(%struct.class_String* noundef %0, i32 noundef %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i32, align 4
  %5 = alloca %struct.class_Array*, align 8
  %6 = alloca i32, align 4
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i32 %1, i32* %4, align 4
  %7 = call noalias i8* @malloc(i64 noundef 40) #3
  %8 = bitcast i8* %7 to %struct.class_Array*
  store %struct.class_Array* %8, %struct.class_Array** %5, align 8
  %9 = load %struct.class_Array*, %struct.class_Array** %5, align 8
  %10 = load i32, i32* %4, align 4
  call void @Array_public_Constructor(%struct.class_Array* noundef %9, i32 noundef %10)
  %11 = load %struct.class_Array*, %struct.class_Array** %5, align 8
  %12 = bitcast %struct.class_Array* %11 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* noundef %12)
  store i32 0, i32* %6, align 4
  br label %13

13:                                               ; preds = %22, %2
  %14 = load i32, i32* %6, align 4
  %15 = load i32, i32* %4, align 4
  %16 = icmp slt i32 %14, %15
  br i1 %16, label %17, label %25

17:                                               ; preds = %13
  %18 = load %struct.class_Array*, %struct.class_Array** %5, align 8
  %19 = load i32, i32* %6, align 4
  %20 = load %struct.class_String*, %struct.class_String** %3, align 8
  %21 = bitcast %struct.class_String* %20 to %struct.class_Any*
  call void @Array_public_SetElement(%struct.class_Array* noundef %18, i32 noundef %19, %struct.class_Any* noundef %21)
  br label %22

22:                                               ; preds = %17
  %23 = load i32, i32* %6, align 4
  %24 = add nsw i32 %23, 1
  store i32 %24, i32* %6, align 4
  br label %13, !llvm.loop !8

25:                                               ; preds = %13
  %26 = load %struct.class_Array*, %struct.class_Array** %5, align 8
  %27 = bitcast %struct.class_Array* %26 to %struct.class_Array_String*
  ret %struct.class_Array_String* %27
}

declare void @Array_public_Constructor(%struct.class_Array* noundef, i32 noundef) #1

declare void @Array_public_SetElement(%struct.class_Array* noundef, i32 noundef, %struct.class_Any* noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %"struct.class_Array_T_array_$b$string$s$$e$"* @test_Get2DStringArray(%struct.class_String* noundef %0, i32 noundef %1, i32 noundef %2) #0 {
  %4 = alloca %struct.class_String*, align 8
  %5 = alloca i32, align 4
  %6 = alloca i32, align 4
  %7 = alloca %struct.class_Array*, align 8
  %8 = alloca %struct.class_Array_String*, align 8
  %9 = alloca i32, align 4
  store %struct.class_String* %0, %struct.class_String** %4, align 8
  store i32 %1, i32* %5, align 4
  store i32 %2, i32* %6, align 4
  %10 = call noalias i8* @malloc(i64 noundef 40) #3
  %11 = bitcast i8* %10 to %struct.class_Array*
  store %struct.class_Array* %11, %struct.class_Array** %7, align 8
  %12 = load %struct.class_Array*, %struct.class_Array** %7, align 8
  %13 = load i32, i32* %5, align 4
  call void @Array_public_Constructor(%struct.class_Array* noundef %12, i32 noundef %13)
  %14 = load %struct.class_Array*, %struct.class_Array** %7, align 8
  %15 = bitcast %struct.class_Array* %14 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* noundef %15)
  %16 = load %struct.class_String*, %struct.class_String** %4, align 8
  %17 = load i32, i32* %6, align 4
  %18 = call %struct.class_Array_String* @test_GetStringArray(%struct.class_String* noundef %16, i32 noundef %17)
  store %struct.class_Array_String* %18, %struct.class_Array_String** %8, align 8
  store i32 0, i32* %9, align 4
  br label %19

19:                                               ; preds = %28, %3
  %20 = load i32, i32* %9, align 4
  %21 = load i32, i32* %5, align 4
  %22 = icmp slt i32 %20, %21
  br i1 %22, label %23, label %31

23:                                               ; preds = %19
  %24 = load %struct.class_Array*, %struct.class_Array** %7, align 8
  %25 = load i32, i32* %9, align 4
  %26 = load %struct.class_Array_String*, %struct.class_Array_String** %8, align 8
  %27 = bitcast %struct.class_Array_String* %26 to %struct.class_Any*
  call void @Array_public_SetElement(%struct.class_Array* noundef %24, i32 noundef %25, %struct.class_Any* noundef %27)
  br label %28

28:                                               ; preds = %23
  %29 = load i32, i32* %9, align 4
  %30 = add nsw i32 %29, 1
  store i32 %30, i32* %9, align 4
  br label %19, !llvm.loop !9

31:                                               ; preds = %19
  %32 = load %struct.class_Array_String*, %struct.class_Array_String** %8, align 8
  %33 = bitcast %struct.class_Array_String* %32 to %struct.class_Any*
  call void @arc_UnregisterReference(%struct.class_Any* noundef %33)
  %34 = load %struct.class_Array*, %struct.class_Array** %7, align 8
  %35 = bitcast %struct.class_Array* %34 to %"struct.class_Array_T_array_$b$string$s$$e$"*
  ret %"struct.class_Array_T_array_$b$string$s$$e$"* %35
}

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
!6 = distinct !{!6, !7}
!7 = !{!"llvm.loop.mustprogress"}
!8 = distinct !{!8, !7}
!9 = distinct !{!9, !7}
