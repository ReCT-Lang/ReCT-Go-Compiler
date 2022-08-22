; ModuleID = './konsole.c'
source_filename = "./konsole.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }

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
@.str.17 = private unnamed_addr constant [17 x i8] c"\1B[38;2;%d;%d;%dm\00", align 1
@.str.18 = private unnamed_addr constant [17 x i8] c"\1B[48;2;%d;%d;%dm\00", align 1
@.str.19 = private unnamed_addr constant [17 x i8] c"\1B[38;2;%u;%u;%um\00", align 1
@.str.20 = private unnamed_addr constant [17 x i8] c"\1B[48;2;%u;%u;%um\00", align 1
@.str.21 = private unnamed_addr constant [11 x i8] c"\1B[38;5;%dm\00", align 1
@.str.22 = private unnamed_addr constant [11 x i8] c"\1B[48;5;%dm\00", align 1
@.str.23 = private unnamed_addr constant [5 x i8] c"\1B[1m\00", align 1
@.str.24 = private unnamed_addr constant [5 x i8] c"\1B[3m\00", align 1
@.str.25 = private unnamed_addr constant [5 x i8] c"\1B[4m\00", align 1
@.str.26 = private unnamed_addr constant [5 x i8] c"\1B[9m\00", align 1
@.str.27 = private unnamed_addr constant [1 x i8] zeroinitializer, align 1
@.str.28 = private unnamed_addr constant [20 x i8] c"%s\1B[%d8;2;%d;%d;%dm\00", align 1
@.str.29 = private unnamed_addr constant [5 x i8] c"%s%c\00", align 1
@.str.30 = private unnamed_addr constant [2 x i8] c"\0A\00", align 1
@.str.31 = private unnamed_addr constant [5 x i8] c"%s%s\00", align 1
@.str.32 = private unnamed_addr constant [6 x i8] c"%s%s\0A\00", align 1
@.str.33 = private unnamed_addr constant [3 x i8] c"%s\00", align 1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_SetFgRGB(i32 noundef %0, i32 noundef %1, i32 noundef %2) #0 {
  %4 = alloca i32, align 4
  %5 = alloca i32, align 4
  %6 = alloca i32, align 4
  store i32 %0, i32* %4, align 4
  store i32 %1, i32* %5, align 4
  store i32 %2, i32* %6, align 4
  %7 = load i32, i32* %4, align 4
  %8 = load i32, i32* %5, align 4
  %9 = load i32, i32* %6, align 4
  %10 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([17 x i8], [17 x i8]* @.str.17, i64 0, i64 0), i32 noundef %7, i32 noundef %8, i32 noundef %9)
  ret void
}

declare i32 @printf(i8* noundef, ...) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_SetBgRGB(i32 noundef %0, i32 noundef %1, i32 noundef %2) #0 {
  %4 = alloca i32, align 4
  %5 = alloca i32, align 4
  %6 = alloca i32, align 4
  store i32 %0, i32* %4, align 4
  store i32 %1, i32* %5, align 4
  store i32 %2, i32* %6, align 4
  %7 = load i32, i32* %4, align 4
  %8 = load i32, i32* %5, align 4
  %9 = load i32, i32* %6, align 4
  %10 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([17 x i8], [17 x i8]* @.str.18, i64 0, i64 0), i32 noundef %7, i32 noundef %8, i32 noundef %9)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_SetFgCol(i32 noundef %0) #0 {
  %2 = alloca i32, align 4
  %3 = alloca i32, align 4
  %4 = alloca i32, align 4
  %5 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %6 = load i32, i32* %2, align 4
  %7 = and i32 %6, 16711680
  %8 = ashr i32 %7, 16
  store i32 %8, i32* %3, align 4
  %9 = load i32, i32* %2, align 4
  %10 = and i32 %9, 65280
  %11 = ashr i32 %10, 8
  store i32 %11, i32* %4, align 4
  %12 = load i32, i32* %2, align 4
  %13 = and i32 %12, 255
  store i32 %13, i32* %5, align 4
  %14 = load i32, i32* %3, align 4
  %15 = load i32, i32* %4, align 4
  %16 = load i32, i32* %5, align 4
  %17 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([17 x i8], [17 x i8]* @.str.19, i64 0, i64 0), i32 noundef %14, i32 noundef %15, i32 noundef %16)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_SetBgCol(i32 noundef %0) #0 {
  %2 = alloca i32, align 4
  %3 = alloca i32, align 4
  %4 = alloca i32, align 4
  %5 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %6 = load i32, i32* %2, align 4
  %7 = and i32 %6, 16711680
  %8 = ashr i32 %7, 16
  store i32 %8, i32* %3, align 4
  %9 = load i32, i32* %2, align 4
  %10 = and i32 %9, 65280
  %11 = ashr i32 %10, 8
  store i32 %11, i32* %4, align 4
  %12 = load i32, i32* %2, align 4
  %13 = and i32 %12, 255
  store i32 %13, i32* %5, align 4
  %14 = load i32, i32* %3, align 4
  %15 = load i32, i32* %4, align 4
  %16 = load i32, i32* %5, align 4
  %17 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([17 x i8], [17 x i8]* @.str.20, i64 0, i64 0), i32 noundef %14, i32 noundef %15, i32 noundef %16)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_SetFgInt(i32 noundef %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([11 x i8], [11 x i8]* @.str.21, i64 0, i64 0), i32 noundef %3)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_SetBgInt(i32 noundef %0) #0 {
  %2 = alloca i32, align 4
  store i32 %0, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([11 x i8], [11 x i8]* @.str.22, i64 0, i64 0), i32 noundef %3)
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_Bold() #0 {
  %1 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([5 x i8], [5 x i8]* @.str.23, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_Italic() #0 {
  %1 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([5 x i8], [5 x i8]* @.str.24, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_Underline() #0 {
  %1 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([5 x i8], [5 x i8]* @.str.25, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_CrossedOut() #0 {
  %1 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([5 x i8], [5 x i8]* @.str.26, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i8* @internal_konsole_GetGradient(i8* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4, i32 noundef %5, i32 noundef %6, i1 noundef zeroext %7) #0 {
  %9 = alloca i8*, align 8
  %10 = alloca i32, align 4
  %11 = alloca i32, align 4
  %12 = alloca i32, align 4
  %13 = alloca i32, align 4
  %14 = alloca i32, align 4
  %15 = alloca i32, align 4
  %16 = alloca i8, align 1
  %17 = alloca i8*, align 8
  %18 = alloca i32, align 4
  %19 = alloca i32, align 4
  %20 = alloca i32, align 4
  %21 = alloca i32, align 4
  %22 = alloca float, align 4
  %23 = alloca float, align 4
  %24 = alloca float, align 4
  %25 = alloca i32, align 4
  %26 = alloca i32, align 4
  %27 = alloca i32, align 4
  %28 = alloca i32, align 4
  store i8* %0, i8** %9, align 8
  store i32 %1, i32* %10, align 4
  store i32 %2, i32* %11, align 4
  store i32 %3, i32* %12, align 4
  store i32 %4, i32* %13, align 4
  store i32 %5, i32* %14, align 4
  store i32 %6, i32* %15, align 4
  %29 = zext i1 %7 to i8
  store i8 %29, i8* %16, align 1
  store i8* getelementptr inbounds ([1 x i8], [1 x i8]* @.str.27, i64 0, i64 0), i8** %17, align 8
  %30 = load i8*, i8** %9, align 8
  %31 = call i64 @strlen(i8* noundef %30) #4
  %32 = trunc i64 %31 to i32
  store i32 %32, i32* %18, align 4
  %33 = load i32, i32* %13, align 4
  %34 = load i32, i32* %10, align 4
  %35 = sub nsw i32 %33, %34
  store i32 %35, i32* %19, align 4
  %36 = load i32, i32* %14, align 4
  %37 = load i32, i32* %11, align 4
  %38 = sub nsw i32 %36, %37
  store i32 %38, i32* %20, align 4
  %39 = load i32, i32* %15, align 4
  %40 = load i32, i32* %12, align 4
  %41 = sub nsw i32 %39, %40
  store i32 %41, i32* %21, align 4
  %42 = load i32, i32* %19, align 4
  %43 = sitofp i32 %42 to float
  %44 = load i32, i32* %18, align 4
  %45 = sitofp i32 %44 to float
  %46 = fdiv float %43, %45
  store float %46, float* %22, align 4
  %47 = load i32, i32* %20, align 4
  %48 = sitofp i32 %47 to float
  %49 = load i32, i32* %18, align 4
  %50 = sitofp i32 %49 to float
  %51 = fdiv float %48, %50
  store float %51, float* %23, align 4
  %52 = load i32, i32* %21, align 4
  %53 = sitofp i32 %52 to float
  %54 = load i32, i32* %18, align 4
  %55 = sitofp i32 %54 to float
  %56 = fdiv float %53, %55
  store float %56, float* %24, align 4
  store i32 0, i32* %25, align 4
  br label %57

57:                                               ; preds = %100, %8
  %58 = load i32, i32* %25, align 4
  %59 = load i32, i32* %18, align 4
  %60 = icmp slt i32 %58, %59
  br i1 %60, label %61, label %103

61:                                               ; preds = %57
  %62 = load float, float* %22, align 4
  %63 = load i32, i32* %25, align 4
  %64 = sitofp i32 %63 to float
  %65 = fmul float %62, %64
  %66 = fptosi float %65 to i32
  store i32 %66, i32* %26, align 4
  %67 = load float, float* %23, align 4
  %68 = load i32, i32* %25, align 4
  %69 = sitofp i32 %68 to float
  %70 = fmul float %67, %69
  %71 = fptosi float %70 to i32
  store i32 %71, i32* %27, align 4
  %72 = load float, float* %24, align 4
  %73 = load i32, i32* %25, align 4
  %74 = sitofp i32 %73 to float
  %75 = fmul float %72, %74
  %76 = fptosi float %75 to i32
  store i32 %76, i32* %28, align 4
  %77 = load i8*, i8** %17, align 8
  %78 = load i8, i8* %16, align 1
  %79 = trunc i8 %78 to i1
  %80 = zext i1 %79 to i64
  %81 = select i1 %79, i32 4, i32 3
  %82 = load i32, i32* %10, align 4
  %83 = load i32, i32* %26, align 4
  %84 = add nsw i32 %82, %83
  %85 = load i32, i32* %11, align 4
  %86 = load i32, i32* %27, align 4
  %87 = add nsw i32 %85, %86
  %88 = load i32, i32* %12, align 4
  %89 = load i32, i32* %28, align 4
  %90 = add nsw i32 %88, %89
  %91 = call i32 (i8**, i8*, ...) @asprintf(i8** noundef %17, i8* noundef getelementptr inbounds ([20 x i8], [20 x i8]* @.str.28, i64 0, i64 0), i8* noundef %77, i32 noundef %81, i32 noundef %84, i32 noundef %87, i32 noundef %90) #5
  %92 = load i8*, i8** %17, align 8
  %93 = load i8*, i8** %9, align 8
  %94 = load i32, i32* %25, align 4
  %95 = sext i32 %94 to i64
  %96 = getelementptr inbounds i8, i8* %93, i64 %95
  %97 = load i8, i8* %96, align 1
  %98 = sext i8 %97 to i32
  %99 = call i32 (i8**, i8*, ...) @asprintf(i8** noundef %17, i8* noundef getelementptr inbounds ([5 x i8], [5 x i8]* @.str.29, i64 0, i64 0), i8* noundef %92, i32 noundef %98) #5
  br label %100

100:                                              ; preds = %61
  %101 = load i32, i32* %25, align 4
  %102 = add nsw i32 %101, 1
  store i32 %102, i32* %25, align 4
  br label %57, !llvm.loop !6

103:                                              ; preds = %57
  %104 = load i8*, i8** %17, align 8
  ret i8* %104
}

; Function Attrs: nounwind readonly willreturn
declare i64 @strlen(i8* noundef) #2

; Function Attrs: nounwind
declare i32 @asprintf(i8** noundef, i8* noundef, ...) #3

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @konsole_GetGradient(%struct.class_String* noundef %0, i1 noundef zeroext %1, i32 noundef %2, i32 noundef %3, i32 noundef %4, i32 noundef %5, i32 noundef %6, i32 noundef %7) #0 {
  %9 = alloca %struct.class_String*, align 8
  %10 = alloca %struct.class_String*, align 8
  %11 = alloca i8, align 1
  %12 = alloca i32, align 4
  %13 = alloca i32, align 4
  %14 = alloca i32, align 4
  %15 = alloca i32, align 4
  %16 = alloca i32, align 4
  %17 = alloca i32, align 4
  %18 = alloca i8*, align 8
  %19 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %10, align 8
  %20 = zext i1 %1 to i8
  store i8 %20, i8* %11, align 1
  store i32 %2, i32* %12, align 4
  store i32 %3, i32* %13, align 4
  store i32 %4, i32* %14, align 4
  store i32 %5, i32* %15, align 4
  store i32 %6, i32* %16, align 4
  store i32 %7, i32* %17, align 4
  %21 = load %struct.class_String*, %struct.class_String** %10, align 8
  %22 = icmp eq %struct.class_String* %21, null
  br i1 %22, label %23, label %24

23:                                               ; preds = %8
  store %struct.class_String* null, %struct.class_String** %9, align 8
  br label %46

24:                                               ; preds = %8
  %25 = load %struct.class_String*, %struct.class_String** %10, align 8
  %26 = getelementptr inbounds %struct.class_String, %struct.class_String* %25, i32 0, i32 2
  %27 = load i8*, i8** %26, align 8
  %28 = load i32, i32* %12, align 4
  %29 = load i32, i32* %13, align 4
  %30 = load i32, i32* %14, align 4
  %31 = load i32, i32* %15, align 4
  %32 = load i32, i32* %16, align 4
  %33 = load i32, i32* %17, align 4
  %34 = load i8, i8* %11, align 1
  %35 = trunc i8 %34 to i1
  %36 = call i8* @internal_konsole_GetGradient(i8* noundef %27, i32 noundef %28, i32 noundef %29, i32 noundef %30, i32 noundef %31, i32 noundef %32, i32 noundef %33, i1 noundef zeroext %35)
  store i8* %36, i8** %18, align 8
  %37 = call noalias i8* @malloc(i64 noundef 40) #5
  %38 = bitcast i8* %37 to %struct.class_String*
  store %struct.class_String* %38, %struct.class_String** %19, align 8
  %39 = load %struct.class_String*, %struct.class_String** %19, align 8
  call void @String_public_Constructor(%struct.class_String* noundef %39)
  %40 = load %struct.class_String*, %struct.class_String** %19, align 8
  %41 = bitcast %struct.class_String* %40 to %struct.class_Any*
  call void @arc_RegisterReference(%struct.class_Any* noundef %41)
  %42 = load %struct.class_String*, %struct.class_String** %19, align 8
  %43 = load i8*, i8** %18, align 8
  call void @String_public_Load(%struct.class_String* noundef %42, i8* noundef %43)
  %44 = load i8*, i8** %18, align 8
  call void @free(i8* noundef %44) #5
  %45 = load %struct.class_String*, %struct.class_String** %19, align 8
  store %struct.class_String* %45, %struct.class_String** %9, align 8
  br label %46

46:                                               ; preds = %24, %23
  %47 = load %struct.class_String*, %struct.class_String** %9, align 8
  ret %struct.class_String* %47
}

; Function Attrs: nounwind
declare noalias i8* @malloc(i64 noundef) #3

declare void @String_public_Constructor(%struct.class_String* noundef) #1

declare void @arc_RegisterReference(%struct.class_Any* noundef) #1

declare void @String_public_Load(%struct.class_String* noundef, i8* noundef) #1

; Function Attrs: nounwind
declare void @free(i8* noundef) #3

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_PrintGradient(%struct.class_String* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4, i32 noundef %5, i32 noundef %6) #0 {
  %8 = alloca %struct.class_String*, align 8
  %9 = alloca i32, align 4
  %10 = alloca i32, align 4
  %11 = alloca i32, align 4
  %12 = alloca i32, align 4
  %13 = alloca i32, align 4
  %14 = alloca i32, align 4
  store %struct.class_String* %0, %struct.class_String** %8, align 8
  store i32 %1, i32* %9, align 4
  store i32 %2, i32* %10, align 4
  store i32 %3, i32* %11, align 4
  store i32 %4, i32* %12, align 4
  store i32 %5, i32* %13, align 4
  store i32 %6, i32* %14, align 4
  %15 = load %struct.class_String*, %struct.class_String** %8, align 8
  %16 = load i32, i32* %9, align 4
  %17 = load i32, i32* %10, align 4
  %18 = load i32, i32* %11, align 4
  %19 = load i32, i32* %12, align 4
  %20 = load i32, i32* %13, align 4
  %21 = load i32, i32* %14, align 4
  call void @konsole_WriteGradient(%struct.class_String* noundef %15, i32 noundef %16, i32 noundef %17, i32 noundef %18, i32 noundef %19, i32 noundef %20, i32 noundef %21)
  %22 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([2 x i8], [2 x i8]* @.str.30, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_WriteGradient(%struct.class_String* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4, i32 noundef %5, i32 noundef %6) #0 {
  %8 = alloca %struct.class_String*, align 8
  %9 = alloca i32, align 4
  %10 = alloca i32, align 4
  %11 = alloca i32, align 4
  %12 = alloca i32, align 4
  %13 = alloca i32, align 4
  %14 = alloca i32, align 4
  %15 = alloca i8*, align 8
  store %struct.class_String* %0, %struct.class_String** %8, align 8
  store i32 %1, i32* %9, align 4
  store i32 %2, i32* %10, align 4
  store i32 %3, i32* %11, align 4
  store i32 %4, i32* %12, align 4
  store i32 %5, i32* %13, align 4
  store i32 %6, i32* %14, align 4
  %16 = load %struct.class_String*, %struct.class_String** %8, align 8
  %17 = icmp eq %struct.class_String* %16, null
  br i1 %17, label %18, label %19

18:                                               ; preds = %7
  call void @konsole_Reset()
  br label %33

19:                                               ; preds = %7
  %20 = load %struct.class_String*, %struct.class_String** %8, align 8
  %21 = getelementptr inbounds %struct.class_String, %struct.class_String* %20, i32 0, i32 2
  %22 = load i8*, i8** %21, align 8
  %23 = load i32, i32* %9, align 4
  %24 = load i32, i32* %10, align 4
  %25 = load i32, i32* %11, align 4
  %26 = load i32, i32* %12, align 4
  %27 = load i32, i32* %13, align 4
  %28 = load i32, i32* %14, align 4
  %29 = call i8* @internal_konsole_GetGradient(i8* noundef %22, i32 noundef %23, i32 noundef %24, i32 noundef %25, i32 noundef %26, i32 noundef %27, i32 noundef %28, i1 noundef zeroext false)
  store i8* %29, i8** %15, align 8
  %30 = load i8*, i8** %15, align 8
  %31 = load i8*, i8** @reset, align 8
  %32 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([5 x i8], [5 x i8]* @.str.31, i64 0, i64 0), i8* noundef %30, i8* noundef %31)
  br label %33

33:                                               ; preds = %19, %18
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_PrintBgGradient(%struct.class_String* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4, i32 noundef %5, i32 noundef %6) #0 {
  %8 = alloca %struct.class_String*, align 8
  %9 = alloca i32, align 4
  %10 = alloca i32, align 4
  %11 = alloca i32, align 4
  %12 = alloca i32, align 4
  %13 = alloca i32, align 4
  %14 = alloca i32, align 4
  store %struct.class_String* %0, %struct.class_String** %8, align 8
  store i32 %1, i32* %9, align 4
  store i32 %2, i32* %10, align 4
  store i32 %3, i32* %11, align 4
  store i32 %4, i32* %12, align 4
  store i32 %5, i32* %13, align 4
  store i32 %6, i32* %14, align 4
  %15 = load %struct.class_String*, %struct.class_String** %8, align 8
  %16 = load i32, i32* %9, align 4
  %17 = load i32, i32* %10, align 4
  %18 = load i32, i32* %11, align 4
  %19 = load i32, i32* %12, align 4
  %20 = load i32, i32* %13, align 4
  %21 = load i32, i32* %14, align 4
  call void @konsole_WriteBgGradient(%struct.class_String* noundef %15, i32 noundef %16, i32 noundef %17, i32 noundef %18, i32 noundef %19, i32 noundef %20, i32 noundef %21)
  %22 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([2 x i8], [2 x i8]* @.str.30, i64 0, i64 0))
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_WriteBgGradient(%struct.class_String* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4, i32 noundef %5, i32 noundef %6) #0 {
  %8 = alloca %struct.class_String*, align 8
  %9 = alloca i32, align 4
  %10 = alloca i32, align 4
  %11 = alloca i32, align 4
  %12 = alloca i32, align 4
  %13 = alloca i32, align 4
  %14 = alloca i32, align 4
  %15 = alloca i8*, align 8
  store %struct.class_String* %0, %struct.class_String** %8, align 8
  store i32 %1, i32* %9, align 4
  store i32 %2, i32* %10, align 4
  store i32 %3, i32* %11, align 4
  store i32 %4, i32* %12, align 4
  store i32 %5, i32* %13, align 4
  store i32 %6, i32* %14, align 4
  %16 = load %struct.class_String*, %struct.class_String** %8, align 8
  %17 = icmp eq %struct.class_String* %16, null
  br i1 %17, label %18, label %19

18:                                               ; preds = %7
  call void @konsole_Reset()
  br label %33

19:                                               ; preds = %7
  %20 = load %struct.class_String*, %struct.class_String** %8, align 8
  %21 = getelementptr inbounds %struct.class_String, %struct.class_String* %20, i32 0, i32 2
  %22 = load i8*, i8** %21, align 8
  %23 = load i32, i32* %9, align 4
  %24 = load i32, i32* %10, align 4
  %25 = load i32, i32* %11, align 4
  %26 = load i32, i32* %12, align 4
  %27 = load i32, i32* %13, align 4
  %28 = load i32, i32* %14, align 4
  %29 = call i8* @internal_konsole_GetGradient(i8* noundef %22, i32 noundef %23, i32 noundef %24, i32 noundef %25, i32 noundef %26, i32 noundef %27, i32 noundef %28, i1 noundef zeroext true)
  store i8* %29, i8** %15, align 8
  %30 = load i8*, i8** %15, align 8
  %31 = load i8*, i8** @reset, align 8
  %32 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([5 x i8], [5 x i8]* @.str.31, i64 0, i64 0), i8* noundef %30, i8* noundef %31)
  br label %33

33:                                               ; preds = %19, %18
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_PrintAndReset(%struct.class_String* noundef %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = icmp eq %struct.class_String* %3, null
  br i1 %4, label %5, label %6

5:                                                ; preds = %1
  br label %12

6:                                                ; preds = %1
  %7 = load %struct.class_String*, %struct.class_String** %2, align 8
  %8 = getelementptr inbounds %struct.class_String, %struct.class_String* %7, i32 0, i32 2
  %9 = load i8*, i8** %8, align 8
  %10 = load i8*, i8** @reset, align 8
  %11 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([6 x i8], [6 x i8]* @.str.32, i64 0, i64 0), i8* noundef %9, i8* noundef %10)
  br label %12

12:                                               ; preds = %6, %5
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_WriteAndReset(%struct.class_String* noundef %0) #0 {
  %2 = alloca %struct.class_String*, align 8
  store %struct.class_String* %0, %struct.class_String** %2, align 8
  %3 = load %struct.class_String*, %struct.class_String** %2, align 8
  %4 = icmp eq %struct.class_String* %3, null
  br i1 %4, label %5, label %6

5:                                                ; preds = %1
  br label %12

6:                                                ; preds = %1
  %7 = load %struct.class_String*, %struct.class_String** %2, align 8
  %8 = getelementptr inbounds %struct.class_String, %struct.class_String* %7, i32 0, i32 2
  %9 = load i8*, i8** %8, align 8
  %10 = load i8*, i8** @reset, align 8
  %11 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([5 x i8], [5 x i8]* @.str.31, i64 0, i64 0), i8* noundef %9, i8* noundef %10)
  br label %12

12:                                               ; preds = %6, %5
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @konsole_Reset() #0 {
  %1 = load i8*, i8** @reset, align 8
  %2 = call i32 (i8*, ...) @printf(i8* noundef getelementptr inbounds ([3 x i8], [3 x i8]* @.str.33, i64 0, i64 0), i8* noundef %1)
  ret void
}

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { nounwind readonly willreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #4 = { nounwind readonly willreturn }
attributes #5 = { nounwind }

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
