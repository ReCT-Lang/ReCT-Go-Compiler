; ModuleID = './smallTest.c'
source_filename = "./smallTest.c"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }

@.str = private unnamed_addr constant [12 x i8] c"cool string\00", align 1

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
  call void @String_public_Load(%struct.class_String* noundef %7, i8* noundef getelementptr inbounds ([12 x i8], [12 x i8]* @.str, i64 0, i64 0))
  %8 = load %struct.class_String*, %struct.class_String** %1, align 8
  ret %struct.class_String* %8
}

; Function Attrs: nounwind
declare noalias i8* @malloc(i64 noundef) #1

declare void @String_public_Constructor(%struct.class_String* noundef) #2

declare void @arc_RegisterReference(%struct.class_Any* noundef) #2

declare void @String_public_Load(%struct.class_String* noundef, i8* noundef) #2

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { nounwind }

!llvm.module.flags = !{!0, !1, !2, !3, !4}
!llvm.ident = !{!5}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{i32 7, !"PIC Level", i32 2}
!2 = !{i32 7, !"PIE Level", i32 2}
!3 = !{i32 7, !"uwtable", i32 1}
!4 = !{i32 7, !"frame-pointer", i32 2}
!5 = !{!"clang version 14.0.6"}
