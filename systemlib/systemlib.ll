; ModuleID = './systemlib.c'
source_filename = "./systemlib.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

@isCursorVisible = dso_local global i8 1, align 1
@.str = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.1 = private unnamed_addr constant [3 x i8] c"%s\00", align 1
@.str.2 = private unnamed_addr constant [8 x i8] c"\1B[2J\1B[H\00", align 1
@.str.3 = private unnamed_addr constant [10 x i8] c"%c[%d;%df\00", align 1
@.str.4 = private unnamed_addr constant [8 x i8] c"\1B[?251]\00", align 1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @rct_Print(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  %3 = load i8*, i8** %2, align 8
  %4 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str, i64 0, i64 0), i8* %3)
  ret void
}

declare i32 @printf(i8*, ...) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @rct_Write(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  %3 = load i8*, i8** %2, align 8
  %4 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.str.1, i64 0, i64 0), i8* %3)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i8* @rct_Input() #0 {
  %1 = alloca i8*, align 8
  %2 = alloca i8*, align 8
  %3 = alloca i32, align 4
  %4 = call noalias align 16 i8* @malloc(i64 1042) #4
  store i8* %4, i8** %1, align 8
  store i32 0, i32* %3, align 4
  br label %5

5:                                                ; preds = %37, %0
  %6 = load i8*, i8** %1, align 8
  %7 = icmp ne i8* %6, null
  br i1 %7, label %8, label %17

8:                                                ; preds = %5
  %9 = call i32 @getchar()
  %10 = trunc i32 %9 to i8
  %11 = load i8*, i8** %1, align 8
  %12 = load i32, i32* %3, align 4
  %13 = sext i32 %12 to i64
  %14 = getelementptr inbounds i8, i8* %11, i64 %13
  store i8 %10, i8* %14, align 1
  %15 = sext i8 %10 to i32
  %16 = icmp ne i32 %15, 10
  br label %17

17:                                               ; preds = %8, %5
  %18 = phi i1 [ false, %5 ], [ %16, %8 ]
  br i1 %18, label %19, label %40

19:                                               ; preds = %17
  %20 = load i32, i32* %3, align 4
  %21 = srem i32 %20, 1042
  %22 = icmp eq i32 %21, 1041
  br i1 %22, label %23, label %36

23:                                               ; preds = %19
  %24 = load i8*, i8** %1, align 8
  %25 = load i32, i32* %3, align 4
  %26 = add nsw i32 1042, %25
  %27 = add nsw i32 %26, 1
  %28 = sext i32 %27 to i64
  %29 = mul i64 1, %28
  %30 = call align 16 i8* @realloc(i8* %24, i64 %29) #4
  store i8* %30, i8** %2, align 8
  %31 = icmp eq i8* %30, null
  br i1 %31, label %32, label %34

32:                                               ; preds = %23
  %33 = load i8*, i8** %1, align 8
  call void @free(i8* %33) #4
  br label %34

34:                                               ; preds = %32, %23
  %35 = load i8*, i8** %2, align 8
  store i8* %35, i8** %1, align 8
  br label %36

36:                                               ; preds = %34, %19
  br label %37

37:                                               ; preds = %36
  %38 = load i32, i32* %3, align 4
  %39 = add nsw i32 %38, 1
  store i32 %39, i32* %3, align 4
  br label %5, !llvm.loop !6

40:                                               ; preds = %17
  %41 = load i8*, i8** %1, align 8
  %42 = icmp ne i8* %41, null
  br i1 %42, label %43, label %48

43:                                               ; preds = %40
  %44 = load i8*, i8** %1, align 8
  %45 = load i32, i32* %3, align 4
  %46 = sext i32 %45 to i64
  %47 = getelementptr inbounds i8, i8* %44, i64 %46
  store i8 0, i8* %47, align 1
  br label %48

48:                                               ; preds = %43, %40
  %49 = load i8*, i8** %1, align 8
  ret i8* %49
}

; Function Attrs: nounwind
declare noalias align 16 i8* @malloc(i64) #2

declare i32 @getchar() #1

; Function Attrs: nounwind
declare align 16 i8* @realloc(i8*, i64) #2

; Function Attrs: nounwind
declare void @free(i8*) #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @rct_Clear() #0 {
  %1 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.2, i64 0, i64 0))
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
  %7 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([10 x i8], [10 x i8]* @.str.3, i64 0, i64 0), i32 27, i32 %5, i32 %6)
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
  %10 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.4, i64 0, i64 0))
  br label %13

11:                                               ; preds = %1
  %12 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([8 x i8], [8 x i8]* @.str.4, i64 0, i64 0))
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
  %3 = call i32 @rand() #4
  %4 = load i32, i32* %2, align 4
  %5 = srem i32 %3, %4
  ret i32 %5
}

; Function Attrs: nounwind
declare i32 @rand() #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @rct_Sleep(i32 %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = sdiv i32 %3, 1000
  %5 = call i32 @sleep(i32 %4)
  ret void
}

declare i32 @sleep(i32) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i8* @util_copy_string(i8* %0) #0 {
  %2 = alloca i8*, align 8
  %3 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  %4 = load i8*, i8** %2, align 8
  %5 = call i64 @strlen(i8* %4) #5
  %6 = add i64 %5, 1
  %7 = call noalias align 16 i8* @malloc(i64 %6) #4
  store i8* %7, i8** %3, align 8
  %8 = load i8*, i8** %3, align 8
  %9 = load i8*, i8** %2, align 8
  %10 = call i8* @strcpy(i8* %8, i8* %9) #4
  %11 = load i8*, i8** %3, align 8
  ret i8* %11
}

; Function Attrs: nounwind readonly willreturn
declare i64 @strlen(i8*) #3

; Function Attrs: nounwind
declare i8* @strcpy(i8*, i8*) #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @util_free_string_if_not_null(i8* %0) #0 {
  %2 = alloca i8*, align 8
  store i8* %0, i8** %2, align 8
  %3 = load i8*, i8** %2, align 8
  %4 = icmp ne i8* %3, null
  br i1 %4, label %5, label %7

5:                                                ; preds = %1
  %6 = load i8*, i8** %2, align 8
  call void @free(i8* %6) #4
  br label %7

7:                                                ; preds = %5, %1
  ret void
}

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { nounwind readonly willreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
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
!6 = distinct !{!6, !7}
!7 = !{!"llvm.loop.mustprogress"}
