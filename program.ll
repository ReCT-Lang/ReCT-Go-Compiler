%struct.Any_vTable = type { i8*, i8*, void (i8*)* }
%struct.class_Any = type { %struct.Any_vTable*, i32 }
%struct.String_vTable = type { %struct.Any_vTable*, i8*, void (i8*)* }
%struct.class_Array = type { %struct.String_vTable*, i32, %struct.class_Any**, i32, i32, i32 }
%struct.class_Byte = type { %struct.String_vTable*, i32, i8 }
%struct.class_Float = type { %struct.String_vTable*, i32, float }
%struct.class_Int = type { %struct.String_vTable*, i32, i32 }
%struct.class_Long = type { %struct.String_vTable*, i32, i64 }
%struct.class_String = type { %struct.String_vTable*, i32, i8*, i32, i32, i32 }
%struct.class_Thread = type { %struct.Any_vTable*, i32, i8* (i8*)*, i8*, i64 }
%struct.class_pArray = type { %struct.String_vTable*, i32, i8*, i32, i32, i32, i32 }
%struct.class_AudioStream = type { %struct.String_vTable*, i32, %struct.AudioStream }
%struct.class_BoneInfo = type { %struct.String_vTable*, i32, [32 x i8], i32 }
%struct.class_Camera2D = type { %struct.String_vTable*, i32, %struct.class_Vector2*, %struct.class_Vector2*, float, float }
%struct.class_Camera3D = type { %struct.String_vTable*, i32, %struct.class_Vector3*, %struct.class_Vector3*, %struct.class_Vector3*, float, i32 }
%struct.class_Color = type { %struct.String_vTable*, i32, i32, i32, i32, i32 }
%struct.class_FilePathList = type { %struct.String_vTable*, i32, %struct.FilePathList }
%struct.class_Font = type { %struct.String_vTable*, i32, i32, i32, i32, %struct.class_Texture*, %struct.Vector4*, %struct.GlyphInfo* }
%struct.class_GlyphInfo = type { %struct.String_vTable*, i32, i32, i32, i32, i32, %struct.class_Image* }
%struct.class_Image = type { %struct.String_vTable*, i32, i8*, i32, i32, i32, i32 }
%struct.class_Material = type { %struct.String_vTable*, i32, %struct.class_Shader*, %struct.MaterialMap*, [4 x float] }
%struct.class_MaterialMap = type { %struct.String_vTable*, i32, %struct.class_Texture*, %struct.class_Color*, float }
%struct.class_Matrix = type { %struct.String_vTable*, i32, %struct.Matrix }
%struct.class_Mesh = type { %struct.String_vTable*, i32, %struct.Mesh }
%struct.class_Model = type { %struct.String_vTable*, i32, %struct.Model }
%struct.class_ModelAnimation = type { %struct.String_vTable*, i32, i32, i32, %struct.BoneInfo*, %struct.Transform** }
%struct.class_Music = type { %struct.String_vTable*, i32, %struct.class_AudioStream*, i32, i8, i32, i8* }
%struct.class_NPatchInfo = type { %struct.String_vTable*, i32, %struct.class_Vector4*, i32, i32, i32, i32, i32 }
%struct.class_Ray = type { %struct.String_vTable*, i32, %struct.class_Vector3*, %struct.class_Vector3* }
%struct.class_RayCollision = type { %struct.String_vTable*, i32, i8, float, %struct.class_Vector3*, %struct.class_Vector3* }
%struct.class_RenderTexture = type { %struct.String_vTable*, i32, i32, %struct.class_Texture*, %struct.class_Texture* }
%struct.class_Shader = type { %struct.String_vTable*, i32, i32, i32* }
%struct.class_Sound = type { %struct.String_vTable*, i32, %struct.class_AudioStream*, i32 }
%struct.class_Texture = type { %struct.String_vTable*, i32, i32, i32, i32, i32, i32 }
%struct.class_Transform = type { %struct.String_vTable*, i32, %struct.class_Vector3*, %struct.class_Vector4*, %struct.class_Vector3* }
%struct.class_Vector2 = type { %struct.String_vTable*, i32, float, float }
%struct.class_Vector3 = type { %struct.String_vTable*, i32, float, float, float }
%struct.class_Vector4 = type { %struct.String_vTable*, i32, float, float, float, float }
%struct.class_VrDeviceInfo = type { %struct.String_vTable*, i32, i32, i32, float, float, float, float, float, float, [4 x float], [4 x float] }
%struct.class_VrStereoConfig = type { %struct.String_vTable*, i32, %struct.VrStereoConfig }
%struct.class_Wave = type { %struct.String_vTable*, i32, %struct.Wave }
%struct.AudioStream = type { %struct.rAudioBuffer*, %struct.rAudioProcessor*, i32, i32, i32 }
%struct.BoneInfo = type { [32 x i8], i32 }
%struct.Camera2D = type { %struct.Vector2, %struct.Vector2, float, float }
%struct.Camera3D = type { %struct.Vector3, %struct.Vector3, %struct.Vector3, float, i32 }
%struct.Color = type { i8, i8, i8, i8 }
%struct.FilePathList = type { i32, i32, i8** }
%struct.Font = type { i32, i32, i32, %struct.Texture, %struct.Vector4*, %struct.GlyphInfo* }
%struct.GlyphInfo = type { i32, i32, i32, i32, %struct.Image }
%struct.Image = type { i8*, i32, i32, i32, i32 }
%struct.Material = type { %struct.Shader, %struct.MaterialMap*, [4 x float] }
%struct.MaterialMap = type { %struct.Texture, %struct.Color, float }
%struct.Matrix = type { float, float, float, float, float, float, float, float, float, float, float, float, float, float, float, float }
%struct.Mesh = type { i32, i32, float*, float*, float*, float*, float*, i8*, i16*, float*, float*, i8*, float*, i32, i32* }
%struct.Model = type { %struct.Matrix, i32, i32, %struct.Mesh*, %struct.Material*, i32*, i32, %struct.BoneInfo*, %struct.Transform* }
%struct.ModelAnimation = type { i32, i32, %struct.BoneInfo*, %struct.Transform** }
%struct.Music = type { %struct.AudioStream, i32, i8, i32, i8* }
%struct.NPatchInfo = type { %struct.Vector4, i32, i32, i32, i32, i32 }
%struct.Ray = type { %struct.Vector3, %struct.Vector3 }
%struct.RayCollision = type { i8, float, %struct.Vector3, %struct.Vector3 }
%struct.RenderTexture = type { i32, %struct.Texture, %struct.Texture }
%struct.Shader = type { i32, i32* }
%struct.Sound = type { %struct.AudioStream, i32 }
%struct.Texture = type { i32, i32, i32, i32, i32 }
%struct.Transform = type { %struct.Vector3, %struct.Vector4, %struct.Vector3 }
%struct.Vector2 = type { float, float }
%struct.Vector3 = type { float, float, float }
%struct.Vector4 = type { float, float, float, float }
%struct.VrDeviceInfo = type { i32, i32, float, float, float, float, float, float, [4 x float], [4 x float] }
%struct.VrStereoConfig = type { [2 x %struct.Matrix], [2 x %struct.Matrix], [2 x float], [2 x float], [2 x float], [2 x float], [2 x float], [2 x float] }
%struct.Wave = type { i32, i32, i32, i32, i8* }
%struct.rAudioBuffer = type opaque
%struct.rAudioProcessor = type opaque

@Any_vTable_Const = external global %struct.Any_vTable
@Array_vTable_Const = external global %struct.String_vTable
@Byte_vTable_Const = external global %struct.String_vTable
@Float_vTable_Const = external global %struct.String_vTable
@Int_vTable_Const = external global %struct.String_vTable
@Long_vTable_Const = external global %struct.String_vTable
@String_vTable_Const = external global %struct.String_vTable
@Thread_vTable_Const = external global %struct.Any_vTable
@pArray_vTable_Const = external global %struct.String_vTable
@AudioStream_vTable_Const = external global %struct.String_vTable
@BoneInfo_vTable_Const = external global %struct.String_vTable
@Camera2D_vTable_Const = external global %struct.String_vTable
@Camera3D_vTable_Const = external global %struct.String_vTable
@Color_vTable_Const = external global %struct.String_vTable
@FilePathList_vTable_Const = external global %struct.String_vTable
@Font_vTable_Const = external global %struct.String_vTable
@GlyphInfo_vTable_Const = external global %struct.String_vTable
@Image_vTable_Const = external global %struct.String_vTable
@Material_vTable_Const = external global %struct.String_vTable
@MaterialMap_vTable_Const = external global %struct.String_vTable
@Matrix_vTable_Const = external global %struct.String_vTable
@Mesh_vTable_Const = external global %struct.String_vTable
@Model_vTable_Const = external global %struct.String_vTable
@ModelAnimation_vTable_Const = external global %struct.String_vTable
@Music_vTable_Const = external global %struct.String_vTable
@NPatchInfo_vTable_Const = external global %struct.String_vTable
@Ray_vTable_Const = external global %struct.String_vTable
@RayCollision_vTable_Const = external global %struct.String_vTable
@RenderTexture_vTable_Const = external global %struct.String_vTable
@Shader_vTable_Const = external global %struct.String_vTable
@Sound_vTable_Const = external global %struct.String_vTable
@Texture_vTable_Const = external global %struct.String_vTable
@Transform_vTable_Const = external global %struct.String_vTable
@Vector2_vTable_Const = external global %struct.String_vTable
@Vector3_vTable_Const = external global %struct.String_vTable
@Vector4_vTable_Const = external global %struct.String_vTable
@VrDeviceInfo_vTable_Const = external global %struct.String_vTable
@VrStereoConfig_vTable_Const = external global %struct.String_vTable
@Wave_vTable_Const = external global %struct.String_vTable
@color = global %struct.class_Color* null
@a = global i32 0
@.str.0 = constant [11 x i8] c"TestWindow\00"

declare i32 @printf(i8* %format, ...)

declare i32 @scanf(i8* %format, i8* %dest, ...)

declare void @strcpy(i8* %dest, i8* %src)

declare void @strcat(i8* %dest, i8* %src)

declare i32 @strlen(i8* %str)

declare i32 @strcmp(i8* %left, i8* %right)

declare i8* @malloc(i32 %len)

declare void @free(i8* %dest)

declare i32 @snprintf(i8* %dest, i32 %len, i8* %format, ...)

declare i32 @atoi(i8* %str)

declare double @atof(i8* %str)

declare void @Any_public_Constructor(%struct.class_Any* noundef %0)

declare void @Any_public_Die(i8* noundef %0)

declare void @Array_public_Constructor(%struct.class_Array* noundef %0, i32 noundef %1)

declare void @Array_public_Die(i8* noundef %0)

declare %struct.class_Any* @Array_public_GetElement(%struct.class_Array* noundef %0, i32 noundef %1)

declare void @Array_public_SetElement(%struct.class_Array* noundef %0, i32 noundef %1, %struct.class_Any* noundef %2)

declare i32 @Array_public_GetLength(%struct.class_Array* noundef %0)

declare void @Array_public_Push(%struct.class_Array* noundef %0, %struct.class_Any* noundef %1)

declare void @Byte_public_Constructor(%struct.class_Byte* noundef %0, i8 noundef signext %1)

declare void @Byte_public_Die(i8* noundef %0)

declare i8 @Byte_public_GetValue(%struct.class_Byte* noundef %0)

declare void @Float_public_Constructor(%struct.class_Float* noundef %0, float noundef %1)

declare void @Float_public_Die(i8* noundef %0)

declare float @Float_public_GetValue(%struct.class_Float* noundef %0)

declare void @Int_public_Constructor(%struct.class_Int* noundef %0, i32 noundef %1)

declare void @Int_public_Die(i8* noundef %0)

declare i32 @Int_public_GetValue(%struct.class_Int* noundef %0)

declare void @Long_public_Constructor(%struct.class_Long* noundef %0, i64 noundef %1)

declare void @Long_public_Die(i8* noundef %0)

declare i64 @Long_public_GetValue(%struct.class_Long* noundef %0)

declare void @String_public_Constructor(%struct.class_String* noundef %0)

declare void @String_public_Die(i8* noundef %0)

declare void @String_public_Load(%struct.class_String* noundef %0, i8* noundef %1)

declare void @String_public_Resize(%struct.class_String* noundef %0, i32 noundef %1)

declare void @String_public_AddChar(%struct.class_String* noundef %0, i8 noundef signext %1)

declare %struct.class_String* @String_public_Concat(%struct.class_String* noundef %0, %struct.class_String* noundef %1)

declare i1 @String_public_Equal(%struct.class_String* noundef %0, %struct.class_String* noundef %1)

declare i8* @String_public_GetBuffer(%struct.class_String* noundef %0)

declare i32 @String_public_GetLength(%struct.class_String* noundef %0)

declare %struct.class_String* @String_public_Substring(%struct.class_String* noundef %0, i32 noundef %1, i32 noundef %2)

declare void @Thread_public_Constructor(%struct.class_Thread* noundef %0, i8* (i8*)* noundef %1, i8* noundef %2)

declare void @Thread_public_Die(i8* noundef %0)

declare void @Thread_public_Start(%struct.class_Thread* noundef %0)

declare void @Thread_public_Join(%struct.class_Thread* noundef %0)

declare void @Thread_public_Kill(%struct.class_Thread* noundef %0)

declare void @pArray_public_Constructor(%struct.class_pArray* noundef %0, i32 noundef %1, i32 noundef %2)

declare void @pArray_public_Die(i8* noundef %0)

declare i32 @pArray_public_GetLength(%struct.class_pArray* noundef %0)

declare i8* @pArray_public_Grow(%struct.class_pArray* noundef %0)

declare i8* @pArray_public_GetElementPtr(%struct.class_pArray* noundef %0, i32 noundef %1)

declare void @arc_RegisterReference(%struct.class_Any* noundef %0)

declare void @arc_UnregisterReference(%struct.class_Any* noundef %0)

declare void @arc_DestroyObject(%struct.class_Any* noundef %0)

declare void @arc_RegisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1)

declare void @arc_UnregisterReferenceVerbose(%struct.class_Any* noundef %0, i8* noundef %1)

declare void @exc_Throw(i8* noundef %0)

declare void @exc_ThrowIfNull(i8* noundef %0)

declare void @exc_ThrowIfInvalidCast(%struct.class_Any* noundef %0, %struct.Any_vTable* noundef %1)

declare void @sys_Print(%struct.class_String* noundef %0)

declare void @sys_Write(%struct.class_String* noundef %0)

declare %struct.class_String* @sys_Input()

declare void @sys_Clear()

declare void @sys_SetCursor(i32 noundef %0, i32 noundef %1)

declare void @sys_SetCursorVisible(i1 noundef zeroext %0)

declare i1 @sys_GetCursorVisible()

declare i32 @sys_Random(i32 noundef %0)

declare void @sys_Sleep(i32 noundef %0)

declare i32 @sys_Sqrt(i32 noundef %0)

declare i32 @sys_Now()

declare %struct.class_String* @sys_Char(i32 noundef %0)

declare void @Model_public_Constructor(%struct.class_Model* noundef %0)

declare void @Model_public_Die(i8* noundef %0)

declare %struct.class_Matrix* @Model_public_GetTransform(%struct.class_Model* noundef %0)

declare i32 @Model_public_GetMeshCount(%struct.class_Model* noundef %0)

declare i32 @Model_public_GetMaterialCount(%struct.class_Model* noundef %0)

declare %struct.class_Mesh* @Model_public_GetMesh(%struct.class_Model* noundef %0, i32 noundef %1)

declare %struct.class_Material* @Model_public_GetMaterial(%struct.class_Model* noundef %0, i32 noundef %1)

declare i32 @Model_public_GetMeshMaterial(%struct.class_Model* noundef %0, i32 noundef %1)

declare i32 @Model_public_GetBoneCount(%struct.class_Model* noundef %0)

declare %struct.class_BoneInfo* @Model_public_GetBone(%struct.class_Model* noundef %0, i32 noundef %1)

declare %struct.class_Transform* @Model_public_GetBindPose(%struct.class_Model* noundef %0, i32 noundef %1)

declare void @ModelAnimation_public_Constructor(%struct.class_ModelAnimation* noundef %0, i32 noundef %1, i32 noundef %2)

declare void @ModelAnimation_public_Die(i8* noundef %0)

declare %struct.class_BoneInfo* @ModelAnimation_public_GetBone(%struct.class_ModelAnimation* noundef %0, i32 noundef %1)

declare %struct.class_Transform* @ModelAnimation_public_GetPose(%struct.class_ModelAnimation* noundef %0, i32 noundef %1, i32 noundef %2)

declare void @Music_public_Constructor(%struct.class_Music* noundef %0, %struct.class_AudioStream* noundef %1, i1 noundef zeroext %2, i32 noundef %3)

declare void @Music_public_Die(i8* noundef %0)

declare void @NPatchInfo_public_Constructor(%struct.class_NPatchInfo* noundef %0, %struct.class_Vector4* noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4, i32 noundef %5, i32 noundef %6)

declare void @NPatchInfo_public_Die(i8* noundef %0)

declare void @Material_public_Constructor(%struct.class_Material* noundef %0, %struct.class_Shader* noundef %1)

declare void @Material_public_Die(i8* noundef %0)

declare void @Material_public_SetParam(%struct.class_Material* noundef %0, i32 noundef %1, float noundef %2)

declare float @Material_public_GetParam(%struct.class_Material* noundef %0, i32 noundef %1)

declare %struct.class_MaterialMap* @Material_public_GetMap(%struct.class_Material* noundef %0, i32 noundef %1)

declare void @MaterialMap_public_Constructor(%struct.class_MaterialMap* noundef %0, %struct.class_Texture* noundef %1, %struct.class_Color* noundef %2, float noundef %3)

declare void @MaterialMap_public_Die(i8* noundef %0)

declare void @Mesh_public_Constructor(%struct.class_Mesh* noundef %0)

declare void @Mesh_public_Die(i8* noundef %0)

declare void @FilePathList_public_Constructor(%struct.class_FilePathList* noundef %0)

declare void @FilePathList_public_Die(i8* noundef %0)

declare i32 @FilePathList_public_GetCapacity(%struct.class_FilePathList* noundef %0)

declare i32 @FilePathList_public_GetCount(%struct.class_FilePathList* noundef %0)

declare void @GlyphInfo_public_Constructor(%struct.class_GlyphInfo* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4, %struct.class_Image* noundef %5)

declare void @GlyphInfo_public_Die(i8* noundef %0)

declare void @Texture_public_Constructor(%struct.class_Texture* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4)

declare void @Texture_public_Die(i8* noundef %0)

declare void @BoneInfo_public_Constructor(%struct.class_BoneInfo* noundef %0, i32 noundef %1)

declare void @BoneInfo_public_Die(i8* noundef %0)

declare void @Matrix_public_Constructor(%struct.class_Matrix* noundef %0)

declare void @Matrix_public_Die(i8* noundef %0)

declare void @Matrix_public_SetIndex(%struct.class_Matrix* noundef %0, i32 noundef %1, float noundef %2)

declare float @Matrix_public_GetIndex(%struct.class_Matrix* noundef %0, i32 noundef %1)

declare void @Ray_public_Constructor(%struct.class_Ray* noundef %0, %struct.class_Vector3* noundef %1, %struct.class_Vector3* noundef %2)

declare void @Ray_public_Die(i8* noundef %0)

declare void @RenderTexture_public_Constructor(%struct.class_RenderTexture* noundef %0, %struct.class_Texture* noundef %1, %struct.class_Texture* noundef %2)

declare void @RenderTexture_public_Die(i8* noundef %0)

declare void @Shader_public_Constructor(%struct.class_Shader* noundef %0)

declare void @Shader_public_Die(i8* noundef %0)

declare i32 @Shader_public_GetLoc(%struct.class_Shader* noundef %0, i32 noundef %1)

declare void @Camera2D_public_Constructor(%struct.class_Camera2D* noundef %0, %struct.class_Vector2* noundef %1, %struct.class_Vector2* noundef %2, float noundef %3, float noundef %4)

declare void @Camera2D_public_Die(i8* noundef %0)

declare void @Camera3D_public_Constructor(%struct.class_Camera3D* noundef %0, %struct.class_Vector3* noundef %1, %struct.class_Vector3* noundef %2, %struct.class_Vector3* noundef %3, float noundef %4, i32 noundef %5)

declare void @Camera3D_public_Die(i8* noundef %0)

declare void @Font_public_Constructor(%struct.class_Font* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, %struct.class_Texture* noundef %4)

declare void @Font_public_Die(i8* noundef %0)

declare %struct.class_Vector4* @Font_public_GetRec(%struct.class_Font* noundef %0, i32 noundef %1)

declare %struct.class_GlyphInfo* @Font_public_GetGlyph(%struct.class_Font* noundef %0, i32 noundef %1)

declare void @Image_public_Constructor(%struct.class_Image* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4)

declare void @Image_public_Die(i8* noundef %0)

declare %struct.class_Image* @Image_public_ReadRawData(%struct.class_Image* noundef %0, i32 noundef %1, i32 noundef %2)

declare void @RayCollision_public_Constructor(%struct.class_RayCollision* noundef %0, i1 noundef zeroext %1, float noundef %2, %struct.class_Vector3* noundef %3, %struct.class_Vector3* noundef %4)

declare void @RayCollision_public_Die(i8* noundef %0)

declare void @VrDeviceInfo_public_Constructor(%struct.class_VrDeviceInfo* noundef %0, i32 noundef %1, i32 noundef %2, float noundef %3, float noundef %4, float noundef %5, float noundef %6, float noundef %7, float noundef %8)

declare void @VrDeviceInfo_public_Die(i8* noundef %0)

declare void @Sound_public_Constructor(%struct.class_Sound* noundef %0, %struct.class_AudioStream* noundef %1)

declare void @Sound_public_Die(i8* noundef %0)

declare i32 @Sound_public_GetFrameCount(%struct.class_Sound* noundef %0)

declare void @Transform_public_Constructor(%struct.class_Transform* noundef %0, %struct.class_Vector3* noundef %1, %struct.class_Vector4* noundef %2, %struct.class_Vector3* noundef %3)

declare void @Transform_public_Die(i8* noundef %0)

declare void @Vector4_public_Constructor(%struct.class_Vector4* noundef %0, float noundef %1, float noundef %2, float noundef %3, float noundef %4)

declare void @Vector4_public_Die(i8* noundef %0)

declare void @Vector3_public_Constructor(%struct.class_Vector3* noundef %0, float noundef %1, float noundef %2, float noundef %3)

declare void @Vector3_public_Die(i8* noundef %0)

declare void @VrStereoConfig_public_Constructor(%struct.class_VrStereoConfig* noundef %0)

declare void @VrStereoConfig_public_Die(i8* noundef %0)

declare void @Wave_public_Constructor(%struct.class_Wave* noundef %0)

declare void @Wave_public_Die(i8* noundef %0)

declare void @AudioStream_public_Constructor(%struct.class_AudioStream* noundef %0)

declare void @AudioStream_public_Die(i8* noundef %0)

declare i32 @AudioStream_public_GetSampleRate(%struct.class_AudioStream* noundef %0)

declare i32 @AudioStream_public_GetSampleSize(%struct.class_AudioStream* noundef %0)

declare i32 @AudioStream_public_GetChannels(%struct.class_AudioStream* noundef %0)

declare void @Color_public_Constructor(%struct.class_Color* noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4)

declare void @Color_public_Die(i8* noundef %0)

declare void @Vector2_public_Constructor(%struct.class_Vector2* noundef %0, float noundef %1, float noundef %2)

declare void @Vector2_public_Die(i8* noundef %0)

declare void @rlb_InitWindow(i32 noundef %0, i32 noundef %1, %struct.class_String* noundef %2)

declare i1 @rlb_WindowShouldClose()

declare void @rlb_CloseWindow()

declare i1 @rlb_IsWindowReady()

declare i1 @rlb_IsWindowFullscreen()

declare i1 @rlb_IsWindowHidden()

declare i1 @rlb_IsWindowMinimized()

declare i1 @rlb_IsWindowMaximized()

declare i1 @rlb_IsWindowFocused()

declare i1 @rlb_IsWindowResized()

declare i1 @rlb_IsWindowState(i32 noundef %0)

declare void @rlb_ClearWindowState(i32 noundef %0)

declare void @rlb_ToggleFullscreen()

declare void @rlb_MaximizeWindow()

declare void @rlb_MinimizeWindow()

declare void @rlb_RestoreWindow()

declare void @rlb_SetWindowIcon(%struct.class_Image* noundef %0)

declare void @rlb_SetWindowTitle(%struct.class_String* noundef %0)

declare void @rlb_SetWindowPosition(i32 noundef %0, i32 noundef %1)

declare void @rlb_SetWindowMonitor(i32 noundef %0)

declare void @rlb_SetWindowMinSize(i32 noundef %0, i32 noundef %1)

declare void @rlb_SetWindowSize(i32 noundef %0, i32 noundef %1)

declare void @rlb_SetWindowOpacity(float noundef %0)

declare i32 @rlb_GetScreenWidth()

declare i32 @rlb_GetScreenHeight()

declare i32 @rlb_GetRenderWidth()

declare i32 @rlb_GetRenderHeight()

declare i32 @rlb_GetMonitorCount()

declare i32 @rlb_GetCurrentMonitor()

declare %struct.class_Vector2* @rlb_GetMonitorPosition(i32 noundef %0)

declare i32 @rlb_GetMonitorWidth(i32 noundef %0)

declare i32 @rlb_GetMonitorHeight(i32 noundef %0)

declare i32 @rlb_GetMonitorPhysicalWidth(i32 noundef %0)

declare i32 @rlb_GetMonitorPhysicalHeight(i32 noundef %0)

declare i32 @rlb_GetMonitorRefreshRate(i32 noundef %0)

declare %struct.class_Vector2* @rlb_GetWindowPosition()

declare %struct.class_Vector2* @rlb_GetWindowScaleDPI()

declare %struct.class_String* @rlb_GetMonitorName(i32 noundef %0)

declare void @rlb_SetClipboardText(%struct.class_String* noundef %0)

declare %struct.class_String* @rlb_GetClipboardText()

declare void @rlb_EnableEventWaiting()

declare void @rlb_DisableEventWaiting()

declare void @rlb_SwapScreenBuffer()

declare void @rlb_PollInputEvents()

declare void @rlb_WaitTime(float noundef %0)

declare void @rlb_ShowCursor()

declare void @rlb_HideCursor()

declare i1 @rlb_IsCursorHidden()

declare void @rlb_EnableCursor()

declare void @rlb_DisableCursor()

declare i1 @rlb_IsCursorOnScreen()

declare void @rlb_ClearBackground(%struct.class_Color* noundef %0)

declare void @rlb_BeginDrawing()

declare void @rlb_EndDrawing()

declare void @rlb_BeginMode2D(%struct.class_Camera2D* noundef %0)

declare void @rlb_EndMode2D()

declare void @rlb_BeginMode3D(%struct.class_Camera3D* noundef %0)

declare void @rlb_EndMode3D()

declare void @rlb_BeginTextureMode(%struct.class_RenderTexture* noundef %0)

declare void @rlb_EndTextureMode()

declare void @rlb_BeginShaderMode(%struct.class_Shader* noundef %0)

declare void @rlb_EndShaderMode()

declare void @rlb_BeginBlendMode(i32 noundef %0)

declare void @rlb_EndBlendMode()

declare void @rlb_BeginScissorMode(i32 noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3)

declare void @rlb_EndScissorMode()

declare void @rlb_BeginVrStereoMode(%struct.class_VrStereoConfig* noundef %0)

declare void @rlb_EndVrStereoMode()

declare %struct.class_VrStereoConfig* @rlb_LoadVrStereoConfig(%struct.class_VrDeviceInfo* noundef %0)

declare void @rlb_UnloadVrStereoConfig(%struct.class_VrStereoConfig* noundef %0)

declare %struct.class_Shader* @rlb_LoadShader(%struct.class_String* noundef %0, %struct.class_String* noundef %1)

declare %struct.class_Shader* @rlb_LoadShaderFromMemory(%struct.class_String* noundef %0, %struct.class_String* noundef %1)

declare i32 @rlb_GetShaderLocation(%struct.class_Shader* noundef %0, %struct.class_String* noundef %1)

declare i32 @rlb_GetShaderLocationAttrib(%struct.class_Shader* noundef %0, %struct.class_String* noundef %1)

declare void @rlb_SetShaderValue(%struct.class_Shader* noundef %0, i32 noundef %1, %struct.class_Any* noundef %2, i32 noundef %3)

declare void @rlb_SetShaderValueMatrix(%struct.class_Shader* noundef %0, i32 noundef %1, %struct.class_Matrix* noundef %2)

declare void @rlb_SetShaderValueTexture(%struct.class_Shader* noundef %0, i32 noundef %1, %struct.class_Texture* noundef %2)

declare void @rlb_UnloadShader(%struct.class_Shader* noundef %0)

declare %struct.class_Ray* @rlb_GetMouseRay(%struct.class_Vector2* noundef %0, %struct.class_Camera3D* noundef %1)

declare %struct.class_Matrix* @rlb_GetCameraMatrix(%struct.class_Camera3D* noundef %0)

declare %struct.class_Matrix* @rlb_GetCameraMatrix2D(%struct.class_Camera2D* noundef %0)

declare %struct.class_Vector2* @rlb_GetWorldToScreen(%struct.class_Vector3* noundef %0, %struct.class_Camera3D* noundef %1)

declare %struct.class_Vector2* @rlb_GetScreenToWorld2D(%struct.class_Vector2* noundef %0, %struct.class_Camera2D* noundef %1)

declare %struct.class_Vector2* @rlb_GetWorldToScreenEx(%struct.class_Vector3* noundef %0, %struct.class_Camera3D* noundef %1, i32 noundef %2, i32 noundef %3)

declare %struct.class_Vector2* @rlb_GetWorldToScreen2D(%struct.class_Vector2* noundef %0, %struct.class_Camera2D* noundef %1)

declare void @rlb_SetTargetFPS(i32 noundef %0)

declare i32 @rlb_GetFPS()

declare float @rlb_GetFrameTime()

declare float @rlb_GetTime()

declare i32 @rlb_GetRandomValue(i32 noundef %0, i32 noundef %1)

declare void @rlb_SetRandomSeed(i32 noundef %0)

declare void @rlb_TakeScreenshot(%struct.class_String* noundef %0)

declare void @rlb_SetConfigFlags(i32 noundef %0)

declare void @rlb_SetTraceLogLevel(i32 noundef %0)

declare void @rlb_OpenURL(%struct.class_String* noundef %0)

declare i1 @rlb_IsKeyPressed(i32 noundef %0)

declare i1 @rlb_IsKeyDown(i32 noundef %0)

declare i1 @rlb_IsKeyReleased(i32 noundef %0)

declare i1 @rlb_IsKeyUp(i32 noundef %0)

declare void @rlb_SetExitKey(i32 noundef %0)

declare i32 @rlb_GetKeyPressed()

declare i32 @rlb_GetCharPressed()

declare i1 @rlb_IsGamepadAvailable(i32 noundef %0)

declare %struct.class_String* @rlb_GetGamepadName(i32 noundef %0)

declare i1 @rlb_IsGamepadButtonPressed(i32 noundef %0, i32 noundef %1)

declare i1 @rlb_IsGamepadButtonDown(i32 noundef %0, i32 noundef %1)

declare i1 @rlb_IsGamepadButtonReleased(i32 noundef %0, i32 noundef %1)

declare i1 @rlb_IsGamepadButtonUp(i32 noundef %0, i32 noundef %1)

declare i32 @rlb_GetGamepadButtonPressed()

declare i32 @rlb_GetGamepadAxisCount(i32 noundef %0)

declare float @rlb_GetGamepadAxisMovement(i32 noundef %0, i32 noundef %1)

declare i32 @rlb_SetGamepadMappings(%struct.class_String* noundef %0)

declare i1 @rlb_IsMouseButtonPressed(i32 noundef %0)

declare i1 @rlb_IsMouseButtonDown(i32 noundef %0)

declare i1 @rlb_IsMouseButtonReleased(i32 noundef %0)

declare i1 @rlb_IsMouseButtonUp(i32 noundef %0)

declare i32 @rlb_GetMouseX()

declare i32 @rlb_GetMouseY()

declare %struct.class_Vector2* @rlb_GetMousePosition()

declare %struct.class_Vector2* @rlb_GetMouseDelta()

declare void @rlb_SetMousePosition(i32 noundef %0, i32 noundef %1)

declare void @rlb_SetMouseOffset(i32 noundef %0, i32 noundef %1)

declare void @rlb_SetMouseScale(float noundef %0, float noundef %1)

declare float @rlb_GetMouseWheelMove()

declare %struct.class_Vector2* @rlb_GetMouseWheelMoveV()

declare void @rlb_SetMouseCursor(i32 noundef %0)

declare void @rlb_SetCameraMode(%struct.class_Camera3D* noundef %0, i32 noundef %1)

declare %struct.class_Camera3D* @rlb_UpdateCamera(%struct.class_Camera3D* noundef %0)

declare void @rlb_SetCameraPanControl(i32 noundef %0)

declare void @rlb_SetCameraAltControl(i32 noundef %0)

declare void @rlb_SetCameraSmoothZoomControl(i32 noundef %0)

declare void @rlb_SetCameraMoveControls(i32 noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, i32 noundef %4, i32 noundef %5)

declare void @rlb_SetShapesTexture(%struct.class_Texture* noundef %0, %struct.class_Vector4* noundef %1)

declare void @rlb_DrawPixel(i32 noundef %0, i32 noundef %1, %struct.class_Color* noundef %2)

declare void @rlb_DrawPixelV(%struct.class_Vector2* noundef %0, %struct.class_Color* noundef %1)

declare void @rlb_DrawLine(i32 noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawLineV(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, %struct.class_Color* noundef %2)

declare void @rlb_DrawLineEx(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, float noundef %2, %struct.class_Color* noundef %3)

declare void @rlb_DrawLineBezier(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, float noundef %2, %struct.class_Color* noundef %3)

declare void @rlb_DrawLineBezierQuad(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, %struct.class_Vector2* noundef %2, float noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawLineBezierCubic(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, %struct.class_Vector2* noundef %2, %struct.class_Vector2* noundef %3, float noundef %4, %struct.class_Color* noundef %5)

declare void @rlb_DrawCircle(i32 noundef %0, i32 noundef %1, float noundef %2, %struct.class_Color* noundef %3)

declare void @rlb_DrawCircleSector(%struct.class_Vector2* noundef %0, float noundef %1, float noundef %2, float noundef %3, i32 noundef %4, %struct.class_Color* noundef %5)

declare void @rlb_DrawCircleSectorLines(%struct.class_Vector2* noundef %0, float noundef %1, float noundef %2, float noundef %3, i32 noundef %4, %struct.class_Color* noundef %5)

declare void @rlb_DrawCircleGradient(i32 noundef %0, i32 noundef %1, float noundef %2, %struct.class_Color* noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawCircleV(%struct.class_Vector2* noundef %0, float noundef %1, %struct.class_Color* noundef %2)

declare void @rlb_DrawCircleLines(i32 noundef %0, i32 noundef %1, float noundef %2, %struct.class_Color* noundef %3)

declare void @rlb_DrawEllipse(i32 noundef %0, i32 noundef %1, float noundef %2, float noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawEllipseLines(i32 noundef %0, i32 noundef %1, float noundef %2, float noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawRing(%struct.class_Vector2* noundef %0, float noundef %1, float noundef %2, float noundef %3, float noundef %4, i32 noundef %5, %struct.class_Color* noundef %6)

declare void @rlb_DrawRingLines(%struct.class_Vector2* noundef %0, float noundef %1, float noundef %2, float noundef %3, float noundef %4, i32 noundef %5, %struct.class_Color* noundef %6)

declare void @rlb_DrawRectangle(i32 noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawRectangleV(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, %struct.class_Color* noundef %2)

declare void @rlb_DrawRectangleRec(%struct.class_Vector4* noundef %0, %struct.class_Color* noundef %1)

declare void @rlb_DrawRectanglePro(%struct.class_Vector4* noundef %0, %struct.class_Vector2* noundef %1, float noundef %2, %struct.class_Color* noundef %3)

declare void @rlb_DrawRectangleGradientV(i32 noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, %struct.class_Color* noundef %4, %struct.class_Color* noundef %5)

declare void @rlb_DrawRectangleGradientH(i32 noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, %struct.class_Color* noundef %4, %struct.class_Color* noundef %5)

declare void @rlb_DrawRectangleGradientEx(%struct.class_Vector4* noundef %0, %struct.class_Color* noundef %1, %struct.class_Color* noundef %2, %struct.class_Color* noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawRectangleLines(i32 noundef %0, i32 noundef %1, i32 noundef %2, i32 noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawRectangleLinesEx(%struct.class_Vector4* noundef %0, float noundef %1, %struct.class_Color* noundef %2)

declare void @rlb_DrawRectangleRounded(%struct.class_Vector4* noundef %0, float noundef %1, i32 noundef %2, %struct.class_Color* noundef %3)

declare void @rlb_DrawRectangleRoundedLines(%struct.class_Vector4* noundef %0, float noundef %1, i32 noundef %2, float noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawTriangle(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, %struct.class_Vector2* noundef %2, %struct.class_Color* noundef %3)

declare void @rlb_DrawTriangleLines(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, %struct.class_Vector2* noundef %2, %struct.class_Color* noundef %3)

declare void @rlb_DrawPoly(%struct.class_Vector2* noundef %0, i32 noundef %1, float noundef %2, float noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawPolyLines(%struct.class_Vector2* noundef %0, i32 noundef %1, float noundef %2, float noundef %3, %struct.class_Color* noundef %4)

declare void @rlb_DrawPolyLinesEx(%struct.class_Vector2* noundef %0, i32 noundef %1, float noundef %2, float noundef %3, float noundef %4, %struct.class_Color* noundef %5)

declare i1 @rlb_CheckCollisionRecs(%struct.class_Vector4* noundef %0, %struct.class_Vector4* noundef %1)

declare i1 @rlb_CheckCollisionCircles(%struct.class_Vector2* noundef %0, float noundef %1, %struct.class_Vector2* noundef %2, float noundef %3)

declare i1 @rlb_CheckCollisionCircleRec(%struct.class_Vector2* noundef %0, float noundef %1, %struct.class_Vector4* noundef %2)

declare i1 @rlb_CheckCollisionPointRec(%struct.class_Vector2* noundef %0, %struct.class_Vector4* noundef %1)

declare i1 @rlb_CheckCollisionPointCircle(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, float noundef %2)

declare i1 @rlb_CheckCollisionPointTriangle(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, %struct.class_Vector2* noundef %2, %struct.class_Vector2* noundef %3)

declare i1 @rlb_CheckCollisionPointLine(%struct.class_Vector2* noundef %0, %struct.class_Vector2* noundef %1, %struct.class_Vector2* noundef %2, i32 noundef %3)

declare %struct.class_Vector4* @rlb_GetCollisionRec(%struct.class_Vector4* noundef %0, %struct.class_Vector4* noundef %1)

define void @main() {
0:
	br label %semiroot

semiroot:
	%1 = getelementptr %struct.class_Color, %struct.class_Color* null, i32 1
	%2 = ptrtoint %struct.class_Color* %1 to i32
	%3 = call i8* @malloc(i32 %2)
	%4 = bitcast i8* %3 to %struct.class_Color*
	%5 = getelementptr %struct.class_Color, %struct.class_Color* %4, i32 0
	call void @Color_public_Constructor(%struct.class_Color* %5, i32 111, i32 222, i32 121, i32 255)
	%6 = bitcast %struct.class_Color* %5 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %6)
	store %struct.class_Color* %5, %struct.class_Color** @color
	%7 = load %struct.class_Color*, %struct.class_Color** @color
	%8 = bitcast %struct.class_Color* %7 to i8*
	call void @exc_ThrowIfNull(i8* %8)
	%9 = getelementptr %struct.class_Color, %struct.class_Color* %7, i32 0, i32 0
	%10 = load i32, %struct.String_vTable** %9
	store i32 %10, i32* @a
	%11 = getelementptr [11 x i8], [11 x i8]* @.str.0, i32 0, i32 0
	%12 = getelementptr %struct.class_String, %struct.class_String* null, i32 1
	%13 = ptrtoint %struct.class_String* %12 to i32
	%14 = call i8* @malloc(i32 %13)
	%15 = bitcast i8* %14 to %struct.class_String*
	%16 = getelementptr %struct.class_String, %struct.class_String* %15, i32 0
	call void @String_public_Constructor(%struct.class_String* %16)
	%17 = bitcast %struct.class_String* %16 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %17)
	call void @String_public_Load(%struct.class_String* %16, i8* %11)
	call void @rlb_InitWindow(i32 1000, i32 1000, %struct.class_String* %16)
	%18 = bitcast %struct.class_String* %16 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %18)
	br label %continue1

Label1:
	call void @rlb_BeginDrawing()
	%19 = getelementptr %struct.class_Color, %struct.class_Color* null, i32 1
	%20 = ptrtoint %struct.class_Color* %19 to i32
	%21 = call i8* @malloc(i32 %20)
	%22 = bitcast i8* %21 to %struct.class_Color*
	%23 = getelementptr %struct.class_Color, %struct.class_Color* %22, i32 0
	call void @Color_public_Constructor(%struct.class_Color* %23, i32 0, i32 0, i32 0, i32 255)
	%24 = bitcast %struct.class_Color* %23 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %24)
	call void @rlb_ClearBackground(%struct.class_Color* %23)
	%25 = bitcast %struct.class_Color* %23 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %25)
	%26 = getelementptr %struct.class_Color, %struct.class_Color* null, i32 1
	%27 = ptrtoint %struct.class_Color* %26 to i32
	%28 = call i8* @malloc(i32 %27)
	%29 = bitcast i8* %28 to %struct.class_Color*
	%30 = getelementptr %struct.class_Color, %struct.class_Color* %29, i32 0
	call void @Color_public_Constructor(%struct.class_Color* %30, i32 255, i32 0, i32 0, i32 255)
	%31 = bitcast %struct.class_Color* %30 to %struct.class_Any*
	call void @arc_RegisterReference(%struct.class_Any* %31)
	call void @rlb_DrawLine(i32 0, i32 0, i32 1000, i32 1000, %struct.class_Color* %30)
	%32 = bitcast %struct.class_Color* %30 to %struct.class_Any*
	call void @arc_UnregisterReference(%struct.class_Any* %32)
	call void @rlb_EndDrawing()
	br label %continue1

continue1:
	%33 = call i1 @rlb_WindowShouldClose()
	%34 = icmp ne i1 %33, 0
	%35 = xor i1 %34, true
	br i1 %35, label %Label1, label %break1

break1:
	call void @rlb_CloseWindow()
	; <ReturnARC>
	; </ReturnARC>
	ret void
}
