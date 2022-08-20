; ModuleID = './konsole.c'
source_filename = "./konsole.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Any_vTable = type { i8*, i8*, void (i8*)* }

@.str = private unnamed_addr constant [6 x i8] c"\1B[30m\00", align 1
@.str.1 = private unnamed_addr constant [6 x i8] c"\1B[34m\00", align 1
@.str.2 = private unnamed_addr constant [6 x i8] c"\1B[32m\00", align 1
@.str.3 = private unnamed_addr constant [6 x i8] c"\1B[36m\00", align 1
@.str.4 = private unnamed_addr constant [6 x i8] c"\1B[31m\00", align 1
@.str.5 = private unnamed_addr constant [6 x i8] c"\1B[35m\00", align 1
@.str.6 = private unnamed_addr constant [6 x i8] c"\1B[33m\00", align 1
@.str.7 = private unnamed_addr constant [6 x i8] c"\1B[37m\00", align 1
@.str.8 = private unnamed_addr constant [6 x i8] c"\1B[90m\00", align 1
@.str.9 = private unnamed_addr constant [6 x i8] c"\1B[94m\00", align 1
@.str.10 = private unnamed_addr constant [6 x i8] c"\1B[92m\00", align 1
@.str.11 = private unnamed_addr constant [6 x i8] c"\1B[96m\00", align 1
@.str.12 = private unnamed_addr constant [6 x i8] c"\1B[91m\00", align 1
@.str.13 = private unnamed_addr constant [6 x i8] c"\1B[95m\00", align 1
@.str.14 = private unnamed_addr constant [6 x i8] c"\1B[93m\00", align 1
@.str.15 = private unnamed_addr constant [6 x i8] c"\1B[97m\00", align 1
@colorStrings = dso_local global [16 x i8*] [i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.1, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.2, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.3, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.4, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.5, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.6, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.7, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.8, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.9, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.10, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.11, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.12, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.13, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.14, i32 0, i32 0), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.15, i32 0, i32 0)], align 16
@.str.16 = private unnamed_addr constant [5 x i8] c"\1B[0m\00", align 1
@reset = dso_local global i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.str.16, i32 0, i32 0), align 8
@.str.17 = private unnamed_addr constant [2 x i8] c"\0A\00", align 1
@.str.18 = private unnamed_addr constant [11 x i8] c"\1B[38;5;%dm\00", align 1
@.str.19 = private unnamed_addr constant [3 x i8] c"%s\00", align 1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_PrintInt(%struct.class_String* noundef %0, i32 noundef %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i32, align 4
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i32 %1, i32* %4, align 4
  %5 = load %struct.class_String*, %struct.class_String** %3, align 8
  %6 = load i32, i32* %4, align 4
  call void @konsole_WriteInt(%struct.class_String* noundef %5, i32 noundef %6)
  %7 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([2 x i8], [2 x i8]* @.str.17, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_WriteInt(%struct.class_String* noundef %0, i32 noundef %1) #0 {
  %3 = alloca %struct.class_String*, align 8
  %4 = alloca i32, align 4
  store %struct.class_String* %0, %struct.class_String** %3, align 8
  store i32 %1, i32* %4, align 4
  %5 = load %struct.class_String*, %struct.class_String** %3, align 8
  %6 = icmp eq %struct.class_String* %5, null
  br i1 %6, label %7, label %8

7:                                                ; preds = %2
  br label %17

8:                                                ; preds = %2
  %9 = load i32, i32* %4, align 4
  %10 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([11 x i8], [11 x i8]* @.str.18, i64 0, i64 0), i32 noundef %9)
  %11 = load %struct.class_String*, %struct.class_String** %3, align 8
  %12 = getelementptr inbounds %struct.class_String, %struct.class_String* %11, i32 0, i32 2
  %13 = load i8*, i8** %12, align 8
  %14 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([3 x i8], [3 x i8]* @.str.19, i64 0, i64 0), i8* noundef %13)
  %15 = load i8*, i8** @reset, align 8
  %16 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([3 x i8], [3 x i8]* @.str.19, i64 0, i64 0), i8* noundef %15)
  br label %17

17:                                               ; preds = %8, %7
  ret void
}

declare i32 @printf(i8* noundef, ...) #1

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }

!llvm.module.flags = !{!0, !1, !2, !3, !4}
!llvm.ident = !{!5}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{i32 7, !"PIC Level", i32 2}
!2 = !{i32 7, !"PIE Level", i32 2}
!3 = !{i32 7, !"uwtable", i32 1}
!4 = !{i32 7, !"frame-pointer", i32 2}
!5 = !{!"clang version 14.0.6"}
