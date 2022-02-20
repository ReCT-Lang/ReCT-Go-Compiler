; ModuleID = './arc.c'
source_filename = "./arc.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.class_Any = type { %struct.Any_vTable*, i32 }
%struct.Any_vTable = type { i8*, i8*, void (i8*)* }

@.str = private unnamed_addr constant [59 x i8] c"\1B[36mARC \1B[0m- \1B[32mRegistered %s reference [%d] - %s\1B[0m\0A\00", align 1
@.str.1 = private unnamed_addr constant [61 x i8] c"\1B[36mARC \1B[0m- \1B[33mUnregistered %s reference [%d] - %s\1B[0m\0A\00", align 1
@.str.2 = private unnamed_addr constant [53 x i8] c"\1B[36mARC \1B[0m- \1B[31mDestroying %s instance - %s\1B[0m\0A\00", align 1
@.str.3 = private unnamed_addr constant [44 x i8] c"\1B[36mARC \1B[0m- \1B[0;35mWhat?? [%d] - %s\1B[0m\0A\00", align 1

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
  br label %23

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
  br i1 %14, label %15, label %23

15:                                               ; preds = %6
  %16 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %17 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %16, i32 0, i32 0
  %18 = load %struct.Any_vTable*, %struct.Any_vTable** %17, align 8
  %19 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %18, i32 0, i32 2
  %20 = load void (i8*)*, void (i8*)** %19, align 8
  %21 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %22 = bitcast %struct.class_Any* %21 to i8*
  call void %20(i8* %22)
  br label %23

23:                                               ; preds = %5, %15, %6
  ret void
}

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

declare i32 @printf(i8*, ...) #1

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
  br label %54

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
  br i1 %26, label %27, label %42

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
  br label %54

42:                                               ; preds = %8
  %43 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %44 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %43, i32 0, i32 1
  %45 = load i32, i32* %44, align 8
  %46 = icmp slt i32 %45, 0
  br i1 %46, label %47, label %53

47:                                               ; preds = %42
  %48 = load %struct.class_Any*, %struct.class_Any** %3, align 8
  %49 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %48, i32 0, i32 1
  %50 = load i32, i32* %49, align 8
  %51 = load i8*, i8** %4, align 8
  %52 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([44 x i8], [44 x i8]* @.str.3, i64 0, i64 0), i32 %50, i8* %51)
  br label %53

53:                                               ; preds = %47, %42
  br label %54

54:                                               ; preds = %7, %53, %27
  ret void
}

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }

!llvm.module.flags = !{!0, !1, !2, !3, !4}
!llvm.ident = !{!5}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{i32 7, !"PIC Level", i32 2}
!2 = !{i32 7, !"PIE Level", i32 2}
!3 = !{i32 7, !"uwtable", i32 1}
!4 = !{i32 7, !"frame-pointer", i32 2}
!5 = !{!"clang version 13.0.0"}
