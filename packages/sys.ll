; ModuleID = './sys.c'
source_filename = "./sys.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }

@isCursorVisible = dso_local global i8 1, align 1
@.str = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.1 = private unnamed_addr constant [3 x i8] c"%s\00", align 1
@.str.2 = private unnamed_addr constant [8 x i8] c"\1B[2J\1B[H\00", align 1
@.str.3 = private unnamed_addr constant [10 x i8] c"%c[%d;%df\00", align 1
@.str.4 = private unnamed_addr constant [8 x i8] c"\1B[?251]\00", align 1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_Print(%struct.class_String* noundef %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = bitcast %struct.class_String* %3 to i8*
  call void @exc_ThrowIfNull(i8* noundef %4)
  %5 = load %struct.class_String*, %struct.class_String** %2, align 8
  %6 = getelementptr inbounds %struct.class_String, %struct.class_String* %5, i32 0, i32 2
  %7 = load i8*, i8** %6, align 8
  %8 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([4 x i8], [4 x i8]* @.str, i64 0, i64 0), i8* noundef %7)
  ret void
}

declare void @exc_ThrowIfNull(i8* noundef) #1

declare i32 @printf(i8* noundef, ...) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_Write(%struct.class_String* noundef %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = bitcast %struct.class_String* %3 to i8*
  call void @exc_ThrowIfNull(i8* noundef %4)
  %5 = load %struct.class_String*, %struct.class_String** %2, align 8
  %6 = getelementptr inbounds %struct.class_String, %struct.class_String* %5, i32 0, i32 2
  %7 = load i8*, i8** %6, align 8
  %8 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([3 x i8], [3 x i8]* @.str.1, i64 0, i64 0), i8* noundef %7)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @sys_Input() #0 {
  %1 = alloca i8*, align 8
  %2 = alloca i8*, align 8
  %3 = alloca i32, align 4
  %4 = alloca %struct.class_String*, align 8
  %5 = call noalias i8* @malloc(i64 noundef 1042) #4
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
  %31 = call i8* @realloc(i8* noundef %25, i64 noundef %30) #4
  store i8* %31, i8** %2, align 8
  %32 = icmp eq i8* %31, null
  br i1 %32, label %33, label %35

33:                                               ; preds = %24
  %34 = load i8*, i8** %1, align 8
  call void @free(i8* noundef %34) #4
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
  br label %6, !llvm.loop !6

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
  %50 = call noalias i8* @malloc(i64 noundef 40) #4
  %51 = bitcast i8* %50 to %struct.class_String*
  store %struct.class_String* %51, %struct.class_String** %4, align 8
  %52 = load %struct.class_String*, %struct.class_String** %4, align 8
  call void @String_public_Constructor(%struct.class_String* noundef %52)
  %53 = load %struct.class_String*, %struct.class_String** %4, align 8
  %54 = load i8*, i8** %1, align 8
  call void @String_public_Load(%struct.class_String* noundef %53, i8* noundef %54)
  %55 = load %struct.class_String*, %struct.class_String** %4, align 8
  %56 = bitcast %struct.class_String* %55 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* noundef %56)
  %57 = load i8*, i8** %1, align 8
  %58 = icmp ne i8* %57, null
  br i1 %58, label %59, label %61

59:                                               ; preds = %49
  %60 = load i8*, i8** %1, align 8
  call void @free(i8* noundef %60) #4
  br label %61

61:                                               ; preds = %59, %49
  %62 = load %struct.class_String*, %struct.class_String** %4, align 8
  ret %struct.class_String* %62
}

; Function Attrs: nounwind
declare noalias i8* @malloc(i64 noundef) #2

declare i32 @getchar() #1

; Function Attrs: nounwind
declare i8* @realloc(i8* noundef, i64 noundef) #2

; Function Attrs: nounwind
declare void @free(i8* noundef) #2

declare void @String_public_Constructor(%struct.class_String* noundef) #1

declare void @String_public_Load(%struct.class_String* noundef, i8* noundef) #1

declare void @arc_RegisterReference(%struct.class_Any* noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_Clear() #0 {
  %1 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([8 x i8], [8 x i8]* @.str.2, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_SetCursor(i32 noundef %0, i32 noundef %1) #0 {
  %3 = alloca i32, align 4
  %4 = alloca i32, align 4
  store i32 %0, i32* %3, align 4
  store i32 %1, i32* %4, align 4
  %5 = load i32, i32* %4, align 4
  %6 = load i32, i32* %3, align 4
  %7 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([10 x i8], [10 x i8]* @.str.3, i64 0, i64 0), i32 noundef 27, i32 noundef %5, i32 noundef %6)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_SetCursorVisible(i1 noundef zeroext %0) #0 {
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
  %10 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([8 x i8], [8 x i8]* @.str.4, i64 0, i64 0))
  br label %13

11:                                               ; preds = %1
  %12 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([8 x i8], [8 x i8]* @.str.4, i64 0, i64 0))
  br label %13

13:                                               ; preds = %11, %9
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local zeroext i1 @sys_GetCursorVisible() #0 {
  %1 = load i8, i8* @isCursorVisible, align 1
  %2 = trunc i8 %1 to i1
  ret i1 %2
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @sys_Random(i32 noundef %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = call i32 @rand() #4
  %4 = load i32, i32* %2, align 4
  %5 = srem i32 %3, %4
  ret i32 %5
}

; Function Attrs: nounwind
declare i32 @rand() #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_Sleep(i32 noundef %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = mul nsw i32 %3, 1000
  %5 = call i32 @usleep(i32 noundef %4)
  ret void
}

declare i32 @usleep(i32 noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @sys_Sqrt(i32 noundef %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = sitofp i32 %3 to double
  %5 = call double @sqrt(double noundef %4) #4
  %6 = call double @llvm.floor.f64(double %5)
  %7 = fptosi double %6 to i32
  ret i32 %7
}

; Function Attrs: nounwind
declare double @sqrt(double noundef) #2

; Function Attrs: nofree nosync nounwind readnone speculatable willreturn
declare double @llvm.floor.f64(double) #3

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @sys_Now() #0 {
  %1 = call i64 @clock() #4
  %2 = trunc i64 %1 to i32
  ret i32 %2
}

; Function Attrs: nounwind
declare i64 @clock() #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @sys_Char(i32 noundef %0) #0 {
  %2 = alloca i32, align 4
  %3 = alloca i8*, align 8
  %4 = alloca %struct.class_String*, align 8
  store i32 %0, i32* %2, align 4
  %5 = call noalias i8* @malloc(i64 noundef 1) #4
  store i8* %5, i8** %3, align 8
  %6 = load i32, i32* %2, align 4
  %7 = trunc i32 %6 to i8
  %8 = load i8*, i8** %3, align 8
  %9 = getelementptr inbounds i8, i8* %8, i64 0
  store i8 %7, i8* %9, align 1
  %10 = call noalias i8* @malloc(i64 noundef 40) #4
  %11 = bitcast i8* %10 to %struct.class_String*
  store %struct.class_String* %11, %struct.class_String** %4, align 8
  %12 = load %struct.class_String*, %struct.class_String** %4, align 8
  call void @String_public_Constructor(%struct.class_String* noundef %12)
  %13 = load %struct.class_String*, %struct.class_String** %4, align 8
  %14 = load i8*, i8** %3, align 8
  call void @String_public_Load(%struct.class_String* noundef %13, i8* noundef %14)
  %15 = load %struct.class_String*, %struct.class_String** %4, align 8
  %16 = bitcast %struct.class_String* %15 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* noundef %16)
  %17 = load i8*, i8** %3, align 8
  call void @free(i8* noundef %17) #4
  %18 = load %struct.class_String*, %struct.class_String** %4, align 8
  ret %struct.class_String* %18
}

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { nofree nosync nounwind readnone speculatable willreturn }
attributes #4 = { nounwind }

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
