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
define dso_local void @sys_Print(%struct.class_String* noundef %0) #0  {
  ret void
}

declare void @exc_ThrowIfNull(i8* noundef) #1

declare i32 @printf(i8* noundef, ...) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_Write(%struct.class_String* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @sys_Input() #0  {
  ret void
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
define dso_local void @sys_Clear() #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_SetCursor(i32 noundef %0, i32 noundef %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_SetCursorVisible(i1 noundef zeroext %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local zeroext i1 @sys_GetCursorVisible() #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @sys_Random(i32 noundef %0) #0  {
  ret void
}

; Function Attrs: nounwind
declare i32 @rand() #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @sys_Sleep(i32 noundef %0) #0  {
  ret void
}

declare i32 @usleep(i32 noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @sys_Sqrt(i32 noundef %0) #0  {
  ret void
}

; Function Attrs: nounwind
declare double @sqrt(double noundef) #2

; Function Attrs: nofree nosync nounwind readnone speculatable willreturn
declare double @llvm.floor.f64(double) #3

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @sys_Now() #0  {
  ret void
}

; Function Attrs: nounwind
declare i64 @clock() #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @sys_Char(i32 noundef %0) #0  {
  ret void
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
