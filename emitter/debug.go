package emitter

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

var CurrentVersion string
var SourceFile string

// Declare llvm.dbg.declare function.
// declare void @llvm.dbg.declare(metadata, metadata, metadata)
var llvmDbgDeclare *ir.Func

func (emt *Emitter) InitDbg() {
	llvmDbgDeclare = emt.Module.NewFunc(
		"llvm.dbg.declare",
		types.Void,
		ir.NewParam("", types.Metadata),
		ir.NewParam("", types.Metadata),
		ir.NewParam("", types.Metadata),
	)
}

//func (emt *Emitter) AddMetadata() {
//	// DICompileUnit
//	// !0 = distinct !DICompileUnit(language: DW_LANG_C99, file: !1, producer: "clang version 8.0.0 (tags/RELEASE_800/final)", isOptimized: false, runtimeVersion: 0, emissionKind: FullDebug, enums: !2, nameTableKind: None)
//	diCompileUnit := &metadata.DICompileUnit{
//		MetadataID:   -1,
//		Distinct:     true,
//		Language:     enum.DwarfLangC99, // when the impostor is C
//		Producer:     "rgoc version " + CurrentVersion,
//		EmissionKind: enum.EmissionKindFullDebug, // buggin'
//	}
//
//	// DIFile
//	// !1 = !DIFile(filename: "foo.c", directory: "/home/u/Desktop/foo")
//	diFile := &metadata.DIFile{
//		MetadataID: -1,
//		Filename:   filepath.Base(SourceFile),
//		Directory:  filepath.Dir(SourceFile),
//	}
//	diCompileUnit.File = diFile
//
//	// Empty tuple.
//	//    !2 = !{}
//	emptyTuple := &metadata.Tuple{
//		MetadataID: -1,
//	}
//	diCompileUnit.Enums = emptyTuple
//
//	// Dwarf metadata.
//	//    !3 = !{i32 2, !"Dwarf Version", i32 4}
//	dwarfVersion := &metadata.Tuple{
//		MetadataID: -1,
//		Fields:     []metadata.Field{CI32(2), &metadata.String{Value: "Dwarf Version"}, CI32(4)},
//	}
//	//    !4 = !{i32 2, !"Debug Info Version", i32 3}
//	debugInfoVersion := &metadata.Tuple{
//		MetadataID: -1,
//		Fields:     []metadata.Field{CI32(2), &metadata.String{Value: "Debug Info Version"}, CI32(3)},
//	}
//	//    !5 = !{i32 1, !"wchar_size", i32 4}
//	wcharSize := &metadata.Tuple{
//		MetadataID: -1,
//		Fields:     []metadata.Field{CI32(1), &metadata.String{Value: "wchar_size"}, CI32(4)},
//	}
//	//    !6 = !{i32 7, !"PIC Level", i32 2}
//	picLevel := &metadata.Tuple{
//		MetadataID: -1,
//		Fields:     []metadata.Field{CI32(7), &metadata.String{Value: "PIC Level"}, CI32(2)},
//	}
//	//    !7 = !{i32 7, !"PIE Level", i32 2}
//	pieLevel := &metadata.Tuple{
//		MetadataID: -1,
//		Fields:     []metadata.Field{CI32(7), &metadata.String{Value: "PIE Level"}, CI32(2)},
//	}
//	//    !8 = !{!"clang version 8.0.0 (tags/RELEASE_800/final)"}
//	clangVersion := &metadata.Tuple{
//		MetadataID: -1,
//		Fields:     []metadata.Field{&metadata.String{Value: "rgoc version " + CurrentVersion}},
//	}
//
//	// DISubprogram
//	//    !9 = distinct !DISubprogram(name: "foo", scope: !1, file: !1, line: 1, type: !10, scopeLine: 1, flags: DIFlagPrototyped, spFlags: DISPFlagDefinition, unit: !0, retainedNodes: !2)
//	diSubprogramFoo := &metadata.DISubprogram{
//		MetadataID:    -1,
//		Distinct:      true,
//		Name:          "foo",
//		Scope:         diFile,
//		File:          diFile,
//		Line:          1,
//		ScopeLine:     1,
//		Flags:         enum.DIFlagPrototyped,
//		SPFlags:       enum.DISPFlagDefinition,
//		Unit:          diCompileUnit,
//		RetainedNodes: emptyTuple,
//	}
//
//	// DISubroutineType
//	//    !10 = !DISubroutineType(types: !11)
//	diSubroutineType := &metadata.DISubroutineType{
//		MetadataID: -1,
//	}
//	diSubprogramFoo.Type = diSubroutineType
//
//	// Types tuple.
//	//    !11 = !{!12, !12, !12}
//	typesTuple := &metadata.Tuple{
//		MetadataID: -1,
//	}
//	diSubroutineType.Types = typesTuple
//
//	// DIBasicType
//	//    !12 = !DIBasicType(name: "int", size: 32, encoding: DW_ATE_signed)
//	diBasicTypeI32 := &metadata.DIBasicType{
//		MetadataID: -1,
//		Name:       "int",
//		Size:       32,
//		Encoding:   enum.DwarfAttEncodingSigned,
//	}
//	typesTuple.Fields = []metadata.Field{diBasicTypeI32, diBasicTypeI32, diBasicTypeI32}
//
//	// DILocalVariable
//	//    !13 = !DILocalVariable(name: "a", arg: 1, scope: !9, file: !1, line: 1, type: !12)
//	diLocalVarA = &metadata.DILocalVariable{
//		MetadataID: -1,
//		Name:       "a",
//		Arg:        1,
//		Scope:      diSubprogramFoo,
//		File:       diFile,
//		Line:       1,
//		Type:       diBasicTypeI32,
//	}
//
//	// DILocation
//	//    !14 = !DILocation(line: 1, column: 13, scope: !9)
//	diLocA = &metadata.DILocation{
//		MetadataID: -1,
//		Line:       1,
//		Column:     13,
//		Scope:      diSubprogramFoo,
//	}
//
//	// DILocalVariable
//	//    !15 = !DILocalVariable(name: "b", arg: 2, scope: !9, file: !1, line: 1, type: !12)
//	diLocalVarB = &metadata.DILocalVariable{
//		MetadataID: -1,
//		Name:       "b",
//		Arg:        2,
//		Scope:      diSubprogramFoo,
//		File:       diFile,
//		Line:       1,
//		Type:       diBasicTypeI32,
//	}
//
//	// DILocation
//	//    !16 = !DILocation(line: 1, column: 20, scope: !9)
//	diLocB = &metadata.DILocation{
//		MetadataID: -1,
//		Line:       1,
//		Column:     20,
//		Scope:      diSubprogramFoo,
//	}
//
//	// DILocalVariable
//	//    !17 = !DILocalVariable(name: "sum", scope: !9, file: !1, line: 2, type: !12)
//	diLocalVarSum = &metadata.DILocalVariable{
//		MetadataID: -1,
//		Name:       "sum",
//		Scope:      diSubprogramFoo,
//		File:       diFile,
//		Line:       2,
//		Type:       diBasicTypeI32,
//	}
//
//	// DILocation
//	//    !18 = !DILocation(line: 2, column: 6, scope: !9)
//	diLocSum = &metadata.DILocation{
//		MetadataID: -1,
//		Line:       2,
//		Column:     6,
//		Scope:      diSubprogramFoo,
//	}
//
//	m.MetadataDefs = append(m.MetadataDefs, diCompileUnit)
//	m.MetadataDefs = append(m.MetadataDefs, diFile)
//	m.MetadataDefs = append(m.MetadataDefs, emptyTuple)
//	m.MetadataDefs = append(m.MetadataDefs, dwarfVersion)
//	m.MetadataDefs = append(m.MetadataDefs, debugInfoVersion)
//	m.MetadataDefs = append(m.MetadataDefs, wcharSize)
//	m.MetadataDefs = append(m.MetadataDefs, picLevel)
//	m.MetadataDefs = append(m.MetadataDefs, pieLevel)
//	m.MetadataDefs = append(m.MetadataDefs, clangVersion)
//	m.MetadataDefs = append(m.MetadataDefs, diSubprogramFoo)
//	m.MetadataDefs = append(m.MetadataDefs, diSubroutineType)
//	m.MetadataDefs = append(m.MetadataDefs, typesTuple)
//	m.MetadataDefs = append(m.MetadataDefs, diBasicTypeI32)
//	m.MetadataDefs = append(m.MetadataDefs, diLocalVarA)
//	m.MetadataDefs = append(m.MetadataDefs, diLocA)
//	m.MetadataDefs = append(m.MetadataDefs, diLocalVarB)
//	m.MetadataDefs = append(m.MetadataDefs, diLocB)
//	m.MetadataDefs = append(m.MetadataDefs, diLocalVarSum)
//	m.MetadataDefs = append(m.MetadataDefs, diLocSum)
//
//	// Named metadata definitions.
//	//    !llvm.dbg.cu = !{!0}
//	llvmDbgCu := &metadata.NamedDef{
//		Name:  "llvm.dbg.cu",
//		Nodes: []metadata.Node{diCompileUnit},
//	}
//	m.NamedMetadataDefs["llvm.dbg.cu"] = llvmDbgCu
//	//    !llvm.module.flags = !{!3, !4, !5, !6, !7}
//	llvmModuleFlags := &metadata.NamedDef{
//		Name:  "llvm.module.flags",
//		Nodes: []metadata.Node{dwarfVersion, debugInfoVersion, wcharSize, picLevel, pieLevel},
//	}
//	m.NamedMetadataDefs["llvm.module.flags"] = llvmModuleFlags
//	//    !llvm.ident = !{!8}
//	llvmIdent := &metadata.NamedDef{
//		Name:  "llvm.ident",
//		Nodes: []metadata.Node{clangVersion},
//	}
//	m.NamedMetadataDefs["llvm.ident"] = llvmIdent
//}
