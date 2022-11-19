; ModuleID = './systemlib_lin.bc'
source_filename = "llvm-link"
target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-pc-linux-gnu"

%struct.Standard_vTable = type { i8*, i8*, void (i8*)*, i8* }
%struct.class_Any = type { %struct.Standard_vTable, i32 }
%struct.class_String = type { %struct.Standard_vTable, i32, i8*, i32, i32, i32 }
%struct.class_Array = type { %struct.Standard_vTable, i32, %struct.class_Any**, i32, i32, i32 }
%struct.class_pArray = type { %struct.Standard_vTable, i32, i8*, i32, i32, i32, i32 }
%struct.class_Int = type { %struct.Standard_vTable, i32, i32 }
%struct.class_Byte = type { %struct.Standard_vTable, i32, i8 }
%struct.class_Long = type { %struct.Standard_vTable, i32, i64 }
%struct.class_Float = type { %struct.Standard_vTable, i32, float }
%struct.class_Thread = type { %struct.Standard_vTable, i32, i8* (i8*)*, i8*, i64 }
%union.pthread_attr_t = type { i64, [48 x i8] }

@.str = private unnamed_addr constant [59 x i8] c"\1B[36mARC \1B[0m- \1B[32mRegistered %s reference [%d] - %s\1B[0m\0A\00", align 1
@.str.1 = private unnamed_addr constant [61 x i8] c"\1B[36mARC \1B[0m- \1B[33mUnregistered %s reference [%d] - %s\1B[0m\0A\00", align 1
@.str.2 = private unnamed_addr constant [53 x i8] c"\1B[36mARC \1B[0m- \1B[31mDestroying %s instance - %s\1B[0m\0A\00", align 1
@.str.3 = private unnamed_addr constant [44 x i8] c"\1B[36mARC \1B[0m- \1B[0;35mWhat?? [%d] - %s\1B[0m\0A\00", align 1
@.str.4 = private unnamed_addr constant [4 x i8] c"Any\00", align 1
@Any_vTable_Const = dso_local constant %struct.Standard_vTable { i8* null, i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.4, i32 0, i32 0), void (i8*)* @Any_public_Die, i8* null }, align 8
@.str.1.5 = private unnamed_addr constant [7 x i8] c"String\00", align 1
@String_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.1.5, i32 0, i32 0), void (i8*)* @String_public_Die, i8* null }, align 8
@.str.2.6 = private unnamed_addr constant [42 x i8] c"Substring start-index cannot be negative!\00", align 1
@.str.3.7 = private unnamed_addr constant [37 x i8] c"Substring length cannot be negative!\00", align 1
@.str.4.8 = private unnamed_addr constant [24 x i8] c"Substring out of range!\00", align 1
@.str.5 = private unnamed_addr constant [4 x i8] c"Int\00", align 1
@Int_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.5, i32 0, i32 0), void (i8*)* @Int_public_Die, i8* null }, align 8
@.str.6 = private unnamed_addr constant [5 x i8] c"Byte\00", align 1
@Byte_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.str.6, i32 0, i32 0), void (i8*)* @Byte_public_Die, i8* null }, align 8
@.str.7 = private unnamed_addr constant [5 x i8] c"Long\00", align 1
@Long_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.str.7, i32 0, i32 0), void (i8*)* @Long_public_Die, i8* null }, align 8
@.str.8 = private unnamed_addr constant [6 x i8] c"Float\00", align 1
@Float_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.8, i32 0, i32 0), void (i8*)* @Float_public_Die, i8* null }, align 8
@.str.9 = private unnamed_addr constant [7 x i8] c"Double\00", align 1
@Double_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.9, i32 0, i32 0), void (i8*)* @Double_public_Die, i8* null }, align 8
@.str.10 = private unnamed_addr constant [5 x i8] c"Bool\00", align 1
@Bool_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.str.10, i32 0, i32 0), void (i8*)* @Bool_public_Die, i8* null }, align 8
@.str.11 = private unnamed_addr constant [6 x i8] c"Array\00", align 1
@Array_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([6 x i8], [6 x i8]* @.str.11, i32 0, i32 0), void (i8*)* @Array_public_Die, i8* null }, align 8
@.str.12 = private unnamed_addr constant [26 x i8] c"Array index out of range!\00", align 1
@.str.13 = private unnamed_addr constant [7 x i8] c"pArray\00", align 1
@pArray_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.13, i32 0, i32 0), void (i8*)* @pArray_public_Die, i8* null }, align 8
@.str.14 = private unnamed_addr constant [7 x i8] c"Thread\00", align 1
@Thread_vTable_Const = dso_local constant %struct.Standard_vTable { i8* bitcast (%struct.Standard_vTable* @Any_vTable_Const to i8*), i8* getelementptr inbounds ([7 x i8], [7 x i8]* @.str.14, i32 0, i32 0), void (i8*)* @Thread_public_Die, i8* null }, align 8
@.str.15 = private unnamed_addr constant [45 x i8] c"%s[RUNTIME] %sEncountered Exception! %s'%s'\0A\00", align 1
@.str.1.16 = private unnamed_addr constant [8 x i8] c"\1B[1;31m\00", align 1
@.str.2.17 = private unnamed_addr constant [8 x i8] c"\1B[0;31m\00", align 1
@.str.3.18 = private unnamed_addr constant [19 x i8] c"%s[STACKTRACE] %s\0A\00", align 1
@.str.4.19 = private unnamed_addr constant [8 x i8] c"\1B[1;33m\00", align 1
@.str.5.20 = private unnamed_addr constant [8 x i8] c"\1B[0;33m\00", align 1
@.str.6.21 = private unnamed_addr constant [4 x i8] c".so\00", align 1
@.str.7.22 = private unnamed_addr constant [5 x i8] c".dll\00", align 1
@.str.8.23 = private unnamed_addr constant [4 x i8] c"%s\0A\00", align 1
@.str.9.24 = private unnamed_addr constant [54 x i8] c"Null-Pointer exception! The given reference was null.\00", align 1
@.str.10.25 = private unnamed_addr constant [90 x i8] c"Conversion vTable for output type could not be found! This indicates a broken executable.\00", align 1
@.str.11.26 = private unnamed_addr constant [4 x i8] c"Any\00", align 1
@.str.12.27 = private unnamed_addr constant [50 x i8] c"Object of type %s could not be casted to type %s!\00", align 1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @arc_RegisterReference(%struct.class_Any* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @arc_UnregisterReference(%struct.class_Any* noundef %0) #0  {
  ret void
}

; Function Attrs: nounwind
declare void @free(i8* noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @arc_DestroyObject(%struct.class_Any* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @arc_RegisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1) #0  {
  ret void
}

declare i32 @printf(i8* noundef, ...) #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @arc_UnregisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Any_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @String_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Int_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Byte_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Long_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Float_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Double_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Bool_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Array_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @pArray_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Die(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Any_public_Constructor(%struct.class_Any* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @String_public_Constructor(%struct.class_String* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @String_public_Load(%struct.class_String* noundef %0, i8* noundef %1) #0  {
  ret void
}

; Function Attrs: nounwind readonly willreturn
declare i64 @strlen(i8* noundef) #3

; Function Attrs: nounwind
declare noalias i8* @malloc(i64 noundef) #1

; Function Attrs: argmemonly nofree nounwind willreturn
declare void @llvm.memcpy.p0i8.p0i8.i64(i8* noalias nocapture writeonly, i8* noalias nocapture readonly, i64, i1 immarg) #4

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @String_public_Resize(%struct.class_String* noundef %0, i32 noundef %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @String_public_AddChar(%struct.class_String* noundef %0, i8 noundef signext %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @String_public_Concat(%struct.class_String* noundef %0, %struct.class_String* noundef %1) #0  {
  ret void
}

; Function Attrs: nounwind
declare i8* @strcpy(i8* noundef, i8* noundef) #1

; Function Attrs: nounwind
declare i8* @strcat(i8* noundef, i8* noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local zeroext i1 @String_public_Equal(%struct.class_String* noundef %0, %struct.class_String* noundef %1) #0  {
  ret void
}

; Function Attrs: nounwind readonly willreturn
declare i32 @strcmp(i8* noundef, i8* noundef) #3

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i8* @String_public_GetBuffer(%struct.class_String* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @String_public_GetLength(%struct.class_String* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_String* @String_public_Substring(%struct.class_String* noundef %0, i32 noundef %1, i32 noundef %2) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Int_public_Constructor(%struct.class_Int* noundef %0, i32 noundef %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @Int_public_GetValue(%struct.class_Int* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Byte_public_Constructor(%struct.class_Byte* noundef %0, i8 noundef signext %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local signext i8 @Byte_public_GetValue(%struct.class_Byte* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Long_public_Constructor(%struct.class_Long* noundef %0, i64 noundef %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i64 @Long_public_GetValue(%struct.class_Long* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Float_public_Constructor(%struct.class_Float* noundef %0, float noundef %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local float @Float_public_GetValue(%struct.class_Float* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Double_public_Constructor(%struct.class_Float* noundef %0, double noundef %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local double @Double_public_GetValue(%struct.class_Float* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Bool_public_Constructor(%struct.class_Byte* noundef %0, i1 noundef zeroext %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local zeroext i1 @Bool_public_GetValue(%struct.class_Byte* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Array_public_Constructor(%struct.class_Array* noundef %0, i32 noundef %1) #0  {
  ret void
}

; Function Attrs: nounwind
declare noalias i8* @calloc(i64 noundef, i64 noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local %struct.class_Any* @Array_public_GetElement(%struct.class_Array* noundef %0, i32 noundef %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Array_public_SetElement(%struct.class_Array* noundef %0, i32 noundef %1, %struct.class_Any* noundef %2) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @Array_public_GetLength(%struct.class_Array* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Array_public_Push(%struct.class_Array* noundef %0, %struct.class_Any* noundef %1) #0  {
  ret void
}

; Function Attrs: nounwind
declare i8* @realloc(i8* noundef, i64 noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @pArray_public_Constructor(%struct.class_pArray* noundef %0, i32 noundef %1, i32 noundef %2) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i32 @pArray_public_GetLength(%struct.class_pArray* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i8* @pArray_public_Grow(%struct.class_pArray* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local i8* @pArray_public_GetElementPtr(%struct.class_pArray* noundef %0, i32 noundef %1) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Constructor(%struct.class_Thread* noundef %0, i8* (i8*)* noundef %1, i8* noundef %2) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Start(%struct.class_Thread* noundef %0) #0  {
  ret void
}

; Function Attrs: nounwind
declare i32 @pthread_create(i64* noundef, %union.pthread_attr_t* noundef, i8* (i8*)* noundef, i8* noundef) #1

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Join(%struct.class_Thread* noundef %0) #0  {
  ret void
}

declare i32 @pthread_join(i64 noundef, i8** noundef) #2

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @Thread_public_Kill(%struct.class_Thread* noundef %0) #0  {
  ret void
}

; Function Attrs: noreturn
declare void @pthread_exit(i8* noundef) #5

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @exc_Throw(i8* noundef %0) #0  {
  ret void
}

declare i32 @backtrace(i8** noundef, i32 noundef) #2

; Function Attrs: nounwind
declare i8** @backtrace_symbols(i8** noundef, i32 noundef) #1

; Function Attrs: nounwind readonly willreturn
declare i8* @strstr(i8* noundef, i8* noundef) #3

; Function Attrs: noreturn nounwind
declare void @exit(i32 noundef) #6

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @exc_ThrowIfNull(i8* noundef %0) #0  {
  ret void
}

; Function Attrs: noinline nounwind optnone sspstrong uwtable
define dso_local void @exc_ThrowIfInvalidCast(%struct.class_Any* noundef %0, %struct.Standard_vTable* noundef %1, i8* noundef %2) #0  {
  ret void
}

; Function Attrs: nounwind
declare i32 @snprintf(i8* noundef, i64 noundef, i8* noundef, ...) #1

; Function Attrs: nounwind
declare i32 @sprintf(i8* noundef, i8* noundef, ...) #1

attributes #0 = { noinline nounwind optnone sspstrong uwtable "frame-pointer"="all" "min-legal-vector-width"="0" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #1 = { nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #2 = { "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #3 = { nounwind readonly willreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #4 = { argmemonly nofree nounwind willreturn }
attributes #5 = { noreturn "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #6 = { noreturn nounwind "frame-pointer"="all" "no-trapping-math"="true" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+cx8,+fxsr,+mmx,+sse,+sse2,+x87" "tune-cpu"="generic" }
attributes #7 = { nounwind }
attributes #8 = { nounwind readonly willreturn }
attributes #9 = { noreturn }
attributes #10 = { noreturn nounwind }

!llvm.ident = !{!0, !0, !0}
!llvm.module.flags = !{!1, !2, !3, !4, !5}

!0 = !{!"clang version 14.0.6"}
!1 = !{i32 1, !"wchar_size", i32 4}
!2 = !{i32 7, !"PIC Level", i32 2}
!3 = !{i32 7, !"PIE Level", i32 2}
!4 = !{i32 7, !"uwtable", i32 1}
!5 = !{i32 7, !"frame-pointer", i32 2}
!6 = distinct !{!6, !7}
!7 = !{!"llvm.loop.mustprogress"}
!8 = distinct !{!8, !7}
!9 = distinct !{!9, !7}
!10 = distinct !{!10, !7}
!11 = distinct !{!11, !7}
