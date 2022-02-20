; ModuleID = './arc.c'
source_filename = "./arc.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.class_Any = type { %struct.Any_vTable*, i32 }
%struct.Any_vTable = type { i8*, i8*, void (i8*)* }

@.str = private unnamed_addr constant [54 x i8] c"\1B[36mARC \1B[0m- \1B[32mRegistered %s reference [%d]\1B[0m\0A\00", align 1
@.str.1 = private unnamed_addr constant [56 x i8] c"\1B[36mARC \1B[0m- \1B[33mUnregistered %s reference [%d]\1B[0m\0A\00", align 1
@.str.2 = private unnamed_addr constant [48 x i8] c"\1B[36mARC \1B[0m- \1B[31mDestroying %s instance\1B[0m\0A\00", align 1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @arc_RegisterReference(%struct.class_Any* %0) #0 {
  %2 = alloca %struct.class_Any*, align 8
  store %struct.class_Any* %0, %struct.class_Any** %2, align 8
  %3 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %4 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %3, i32 0, i32 1
  %5 = load i32, i32* %4, align 8
  %6 = add nsw i32 %5, 1
  store i32 %6, i32* %4, align 8
  %7 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %8 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %7, i32 0, i32 0
  %9 = load %struct.Any_vTable*, %struct.Any_vTable** %8, align 8
  %10 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %9, i32 0, i32 1
  %11 = load i8*, i8** %10, align 8
  %12 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %13 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %12, i32 0, i32 1
  %14 = load i32, i32* %13, align 8
  %15 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([54 x i8], [54 x i8]* @.str, i64 0, i64 0), i8* %11, i32 %14)
  ret void
}

declare i32 @printf(i8*, ...) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @arc_UnregisterReference(%struct.class_Any* %0) #0 {
  %2 = alloca %struct.class_Any*, align 8
  store %struct.class_Any* %0, %struct.class_Any** %2, align 8
  %3 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %4 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %3, i32 0, i32 1
  %5 = load i32, i32* %4, align 8
  %6 = add nsw i32 %5, -1
  store i32 %6, i32* %4, align 8
  %7 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %8 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %7, i32 0, i32 0
  %9 = load %struct.Any_vTable*, %struct.Any_vTable** %8, align 8
  %10 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %9, i32 0, i32 1
  %11 = load i8*, i8** %10, align 8
  %12 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %13 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %12, i32 0, i32 1
  %14 = load i32, i32* %13, align 8
  %15 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([56 x i8], [56 x i8]* @.str.1, i64 0, i64 0), i8* %11, i32 %14)
  %16 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %17 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %16, i32 0, i32 1
  %18 = load i32, i32* %17, align 8
  %19 = icmp sle i32 %18, 0
  br i1 %19, label %20, label %34

20:                                               ; preds = %1
  %21 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %22 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %21, i32 0, i32 0
  %23 = load %struct.Any_vTable*, %struct.Any_vTable** %22, align 8
  %24 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %23, i32 0, i32 1
  %25 = load i8*, i8** %24, align 8
  %26 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([48 x i8], [48 x i8]* @.str.2, i64 0, i64 0), i8* %25)
  %27 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %28 = getelementptr inbounds %struct.class_Any, %struct.class_Any* %27, i32 0, i32 0
  %29 = load %struct.Any_vTable*, %struct.Any_vTable** %28, align 8
  %30 = getelementptr inbounds %struct.Any_vTable, %struct.Any_vTable* %29, i32 0, i32 2
  %31 = load void (i8*)*, void (i8*)** %30, align 8
  %32 = load %struct.class_Any*, %struct.class_Any** %2, align 8
  %33 = bitcast %struct.class_Any* %32 to i8*
  call void %31(i8* %33)
  br label %34

34:                                               ; preds = %20, %1
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
