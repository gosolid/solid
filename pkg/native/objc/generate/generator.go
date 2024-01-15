package main

/*
#cgo CFLAGS: -I /opt/homebrew/Cellar/llvm/17.0.2/include
#cgo LDFLAGS: -L /opt/homebrew/Cellar/llvm/17.0.2/lib -l clang
#include <clang-c/Index.h>
#include <stdlib.h>

extern void visitChildren(const CXCursor cursor, const void *context);
extern void ParseHeader(const char *target, const char *isysroot, const char *header, const void * index, const void *translationUnit, const void *cursor);
*/
import "C"

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strings"
	"unsafe"

	"gopkg.in/yaml.v2"
)

type SDK string

var MacOS SDK = "macosx"
var IPhoneOS SDK = "iphoneos"
var IPhoneOSSimulator SDK = "iossimulator"

type TAPI []map[any]any

func (t TAPI) ReexportedFrameworks() []string {
	frameworks := map[string]bool{}

	for _, tapi := range t {
		if l, ok := tapi["reexported-libraries"]; ok {
			for _, i := range l.([]any) {
				if k, ok := i.(map[any]any)["libraries"]; ok {
					for _, s := range k.([]any) {
						frameworks[s.(string)] = true
					}
				}
			}
		}
	}

	f := []string{}

	for framework := range frameworks {
		if !strings.HasPrefix(framework, "/System/Library/Frameworks/") {
			continue
		}

		framework = strings.TrimSuffix(strings.Split(strings.TrimPrefix(framework, "/System/Library/Frameworks/"), "/")[0], ".framework")

		f = append(f, framework)
	}

	return f
}

func (t TAPI) IsExportedClass(cls string) bool {
	for _, tapi := range t {
		if exports, ok := tapi["exports"]; ok {
			for _, export := range exports.([]any) {
				if classes, ok := export.(map[any]any)["objc-classes"]; ok {
					for _, c := range classes.([]any) {
						if c == cls {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

type listener struct {
	seen           map[string]bool
	target         string
	sdkPath        string
	frameworksPath string
	currentFile    string
	tapi           map[string]TAPI
}

func main() {
	parser := &listener{
		seen: map[string]bool{"framework:objc": true},
		tapi: map[string]TAPI{},
	}

	target := os.Args[1]
	sdk := SDK(os.Args[2])
	frameworks := os.Args[3:]

	if b, err := exec.Command("xcrun", "-sdk", string(sdk), "--show-sdk-path").Output(); err != nil {
		panic(err)
	} else {
		parser.target = target
		parser.sdkPath = strings.TrimSpace(string(b))
		parser.frameworksPath = path.Join(strings.TrimSpace(string(b)), "System", "Library", "Frameworks")

		for i, framework := range frameworks {
			if framework == "*" {
				if entries, err := os.ReadDir(parser.frameworksPath); err != nil {
					panic(err)
				} else {
					toAdd := frameworks[:i]

					for _, entry := range entries {
						if entry.IsDir() && strings.HasSuffix(entry.Name(), ".framework") {
							framework := strings.TrimSuffix(path.Base(entry.Name()), ".framework")
							parser.seen["framework:"+framework] = true
							toAdd = append(toAdd, framework)
						}
					}

					frameworks = append(toAdd, frameworks[i+1:]...)
				}
			}
		}

		// if err := parser.ReadHeader(path.Join(parser.sdkPath, "usr", "include", "objc", "objc.h")); err != nil {
		// 	log.Println("WARN: unable to parse objc:", err)
		// }

		if err := parser.ReadHeader(path.Join(parser.sdkPath, "usr", "include", "os", "object.h")); err != nil {
			log.Println("WARN: unable to parse os/object:", err)
		}

		for _, framework := range frameworks {
			if err := parser.ReadFramework(framework); err != nil {
				log.Println("error importing framework", framework+":", err)
			}
		}
	}
}

func (p *listener) SearchHeader(header string, relative bool) error {
	if relative {
		if stats, err := os.Stat(path.Join(path.Dir(p.currentFile), header)); err != nil && !os.IsNotExist(err) {
			return err
		} else if err == nil && !stats.IsDir() {
			return p.ReadHeader(path.Join(path.Dir(p.currentFile), header))
		}
	} else if strings.Contains(header, "/") {
		parts := strings.Split(header, "/")
		headerPath := path.Join(p.frameworksPath, parts[0]+".framework", "Headers", strings.Join(parts[1:], "/"))

		if stats, err := os.Stat(headerPath); err != nil && !os.IsNotExist(err) {
			return err
		} else if err == nil && !stats.IsDir() {
			return p.ReadHeader(headerPath)
		}
	} else {
		headerPath := path.Join(p.sdkPath, "usr", "include", header)

		if stats, err := os.Stat(headerPath); err != nil && !os.IsNotExist(err) {
			return err
		} else if err == nil && !stats.IsDir() {
			return p.ReadHeader(headerPath)
		}
	}

	return os.ErrNotExist
}

func (p *listener) LoadTAPI(framework string) (TAPI, error) {
	if tapi, ok := p.tapi[framework]; ok {
		return tapi, nil
	} else {
		tapi := []map[any]any{}
		seen := map[string]bool{}
		var loadFramework func(framework string) error
		loadFramework = func(framework string) error {
			if seen[framework] {
				return nil
			}
			seen[framework] = true

			if tapiBuf, err := os.ReadFile(path.Join(p.frameworksPath, framework+".framework", framework+".tbd")); err != nil {
				return err
			} else {
				dec := yaml.NewDecoder(bytes.NewReader(tapiBuf))

				for {
					var doc map[any]any
					if dec.Decode(&doc) != nil {
						break
					}
					tapi = append(tapi, doc)
				}

				for _, framework := range TAPI(tapi).ReexportedFrameworks() {
					if err := loadFramework(framework); err != nil {
						return err
					}
				}

				return nil
			}
		}

		p.tapi[framework] = tapi

		if err := loadFramework(framework); err != nil {
			return nil, err
		} else {
			return tapi, nil
		}
	}
}

var reHeaders = regexp.MustCompile(`\s+header\s*"([^"]*)"`)

func (p *listener) ReadFramework(framework string) error {
	if _, ok := p.seen["framework:"+framework]; !ok {
		p.seen["framework:"+framework] = true

		log.Println("📚 importing framework", framework)

		if b, err := os.ReadFile(path.Join(p.frameworksPath, framework+".framework", "Modules", "module.modulemap")); err != nil {
			return err
		} else {
			for _, m := range reHeaders.FindAllSubmatch(b, -1) {
				header := framework + "/" + string(m[1])

				if err := p.SearchHeader(header, false); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

type VisitorPhase int

const (
	VisitorPhaseInitial VisitorPhase = iota
	VisitorPhaseDetail
)

type VisitorContext struct {
	tapi       TAPI
	framework  string
	phase      VisitorPhase
	primitives []string
	structs    map[uint]*Struct
	interfaces map[uint]*Interface
	methods    map[uint]*Method
	protocols  map[uint]*Protocol

	Visitor func(cursor C.CXCursor, parent C.CXCursor, context *VisitorContext) C.int
}

type Loc struct {
	Framework string
	Filename  string
	Line      uint
	Column    uint
}

func (l *Loc) String() string {
	return fmt.Sprintf("%s:%d:%d", l.Filename, l.Line, l.Column)
}

type Pointer struct {
	InnerType any
}

func (p *Pointer) String() string {
	return fmt.Sprintf("%s*", p.InnerType)
}

type VariableArray struct {
	InnerType any
}

func (a *VariableArray) String() string {
	return fmt.Sprintf("%s[]", a.InnerType)
}

type FixedArray struct {
	InnerType any
	Length    int
}

func (a *FixedArray) String() string {
	return fmt.Sprintf("%s[%d]", a.InnerType, a.Length)
}

type Interface struct {
	Loc         *Loc
	ForwardDecl bool
	Name        string
	Super       *Interface
	Protocols   []*Protocol
	Properties  []*Property
	Methods     []*Method
}

type Protocol struct {
	Loc     *Loc
	Name    string
	Methods []*Method
}

type Struct struct {
	Loc         *Loc
	Canonical   bool
	Name        string
	IsAnonymous bool
	Fields      []*Field
}

func (s *Struct) String() string {
	if s.IsAnonymous {
		return "struct (anonymous)"
	}

	return fmt.Sprintf("struct %s.%s", s.Loc.Framework, s.Name)
}

func (s *Protocol) CanonicalName() string {
	return unHyphenate(s.Name, "_")
}

func (s *Interface) CanonicalName() string {
	return unHyphenate(s.Name, "_")
}

func (s *Struct) CanonicalName() string {
	return unHyphenate(s.Name, "_")
}

type Field struct {
	Loc  *Loc
	Name string
	Type any
}

func (f *Field) String() string {
	var type_ string
	if t, ok := f.Type.(*Interface); ok {
		type_ = fmt.Sprintf("%s.%s", t.Loc.Framework, t.Name)
	} else if t, ok := f.Type.(*Struct); ok {
		type_ = fmt.Sprintf("%s.%s", t.Loc.Framework, t.Name)
	} else {
		panic("ERR: invalid field type")
	}

	return fmt.Sprintf("%s %s", type_, f.Name)
}

func (p *Protocol) String() string {
	return fmt.Sprintf("protocol %s.%s", p.Loc.Framework, p.Name)
}

func (i *Interface) String() string {
	if i.Super == nil {
		return fmt.Sprintf("interface %s.%s", i.Loc.Framework, i.Name)
	} else {
		return fmt.Sprintf("interface %s.%s : %s.%s", i.Loc.Framework, i.Name, i.Super.Loc.Framework, i.Super.Name)
	}
}

type Property struct {
	Loc  *Loc
	Name string
	Type any
}

type Method struct {
	Loc        *Loc
	Name       string
	ReturnType any
	Instance   bool
	Interface  *Interface
	Parameters []*Parameter
}

func (m *Method) String() string {
	var returnType string
	if t, ok := m.ReturnType.(*Interface); ok {
		returnType = t.Name
	} else if t, ok := m.ReturnType.(*Struct); ok {
		returnType = t.Name
	} else {
		panic("ERR: invalid return type")
	}

	if m.Instance {
		return fmt.Sprintf("- (%s)%s", returnType, m.Name)
	} else {
		return fmt.Sprintf("+ (%s)%s", returnType, m.Name)
	}
}

type Parameter struct {
	Loc  *Loc
	Name string
	Type any
}

func (p *Parameter) String() string {
	var type_ string
	if t, ok := p.Type.(*Interface); ok {
		type_ = fmt.Sprintf("%s.%s", t.Loc.Framework, t.Name)
	} else if t, ok := p.Type.(*Struct); ok {
		type_ = fmt.Sprintf("%s.%s", t.Loc.Framework, t.Name)
	} else {
		panic("ERR: invalid field type")
	}

	return fmt.Sprintf("(%s)%s", type_, p.Name)
}

//export visitor
func visitor(cursor C.CXCursor, parent C.CXCursor, context unsafe.Pointer) C.int {
	visitorContext := (*VisitorContext)(context)
	return visitorContext.Visitor(cursor, parent, visitorContext)
}

func (v *VisitorContext) ResolveType(t C.CXType) any {
	pointers := 0
	arrays := []int{}

	adornType := func(t any) any {
		for i := 0; i < pointers; i++ {
			t = &Pointer{InnerType: t}
		}
		for i := len(arrays) - 1; i >= 0; i-- {
			a := arrays[i]

			if a == -1 {
				t = &VariableArray{InnerType: t}
			} else {
				t = &FixedArray{InnerType: t, Length: a}
			}
		}
		return t
	}

	for t.kind == C.CXType_Pointer || t.kind == C.CXType_ObjCObjectPointer {
		pointers++
		t = C.clang_getPointeeType(t)
	}
	t = C.clang_getUnqualifiedType(t)

	if t.kind == C.CXType_Enum {
		return "int"
	}

	spelling := C.clang_getTypeSpelling(t)
	ts := C.GoString(C.clang_getCString(spelling))
	C.clang_disposeString(spelling)

	if ts == "BOOL" {
		return adornType("BOOL")
	}

	pointers = 0

	canonicalType := C.clang_getCanonicalType(t)
	for canonicalType.kind == C.CXType_Vector || canonicalType.kind == C.CXType_IncompleteArray || canonicalType.kind == C.CXType_ConstantArray || canonicalType.kind == C.CXType_VariableArray || canonicalType.kind == C.CXType_DependentSizedArray {
		if canonicalType.kind == C.CXType_ConstantArray {
			size := C.clang_getArraySize(canonicalType)
			arrays = append(arrays, int(size))
		} else {
			arrays = append(arrays, -1)
		}
		canonicalType = C.clang_getElementType(canonicalType)
	}
	for canonicalType.kind == C.CXType_Pointer || canonicalType.kind == C.CXType_ObjCObjectPointer {
		pointers++
		canonicalType = C.clang_getPointeeType(canonicalType)
	}
	canonicalType = C.clang_getUnqualifiedType(canonicalType)
	if canonicalType.kind == C.CXType_Elaborated {
		canonicalType = C.clang_Type_getNamedType(canonicalType)
	}
	if canonicalType.kind == C.CXType_ObjCObject {
		canonicalType = C.clang_Type_getObjCObjectBaseType(canonicalType)
	}

	if canonicalType.kind == C.CXType_Enum {
		return adornType("int")
	}
	if canonicalType.kind == C.CXType_FunctionProto || canonicalType.kind == C.CXType_FunctionNoProto || canonicalType.kind == C.CXType_BlockPointer {
		return adornType("func")
	}

	ctypeSpelling := C.clang_getTypeSpelling(canonicalType)
	cts := C.GoString(C.clang_getCString(ctypeSpelling))
	C.clang_disposeString(ctypeSpelling)

	for _, primitive := range v.primitives {
		if primitive == cts {
			return adornType(primitive)
		}
	}

	c := C.clang_getTypeDeclaration(canonicalType)

	// cursorType := C.clang_getCursorType(c)
	// spelling = C.clang_getTypeSpelling(cursorType)
	// cts = C.GoString(C.clang_getCString(spelling))
	// C.clang_disposeString(spelling)

	if c.kind == C.CXCursor_UnionDecl {
		return nil
	}

	ch := uint(C.clang_hashCursor(c))

	if struct_, ok := v.structs[ch]; ok {
		return adornType(struct_)
	}

	if protocol, ok := v.structs[ch]; ok {
		return adornType(protocol)
	}

	if iface, ok := v.interfaces[ch]; ok {
		return adornType(iface)
	}

	for _, struct_ := range v.structs {
		if struct_.Name == cts {
			return adornType(struct_)
		}
	}

	for _, protocol := range v.protocols {
		if protocol.Name == cts {
			return adornType(protocol)
		}
	}

	for _, iface := range v.interfaces {
		if iface.Name == cts {
			return adornType(iface)
		}
	}

	log.Println(v.framework, v.structs, canonicalType)
	panic("ERR: previous declaration for type not found: " + ts + " (" + cts + ")")
}

func (p *listener) ReadHeader(header string) error {
	if p.seen[header] {
		return nil
	} else {
		p.seen[header] = true
	}

	oldCurrentFile := p.currentFile
	p.currentFile = header
	defer func() {
		p.currentFile = oldCurrentFile
	}()

	var currentFramework string
	if strings.HasPrefix(header, p.frameworksPath) {
		currentFramework = strings.TrimSuffix(strings.Split(strings.TrimPrefix(header, p.frameworksPath), "/")[1], ".framework")
	} else {
		currentFramework = "objc"
	}

	var tapi TAPI
	if currentFramework != "objc" {
		if t, err := p.LoadTAPI(currentFramework); err != nil {
			return err
		} else {
			tapi = t
		}
	}

	ch := C.CString(header)
	defer C.free(unsafe.Pointer(ch))

	target := C.CString(p.target)
	defer C.free(unsafe.Pointer(target))

	isysroot := C.CString(p.sdkPath)
	defer C.free(unsafe.Pointer(isysroot))

	context := &VisitorContext{
		tapi:      tapi,
		framework: currentFramework,
		methods:   map[uint]*Method{},
		primitives: []string{
			"void",
			"instancetype",
			"Class",
			"SEL",
			"Protocol",
			"id",
			"char",
			"signed char",
			"struct _NSZone",
			"unsigned __int128",
			"unsigned char",
			"signed char",
			"short",
			"unsigned short",
			"int",
			"unsigned int",
			"long",
			"unsigned long",
			"long long",
			"unsigned long long",
			"float",
			"double",
			"long double",
		},
		structs:    map[uint]*Struct{},
		interfaces: map[uint]*Interface{},
		protocols:  map[uint]*Protocol{},
		Visitor: func(cursor, parent C.CXCursor, context *VisitorContext) C.int {
			if cursor.kind == 0 {
				return 0
			}
			kind := C.clang_getCursorKind(cursor)

			r := C.clang_getCursorExtent(cursor)
			location := C.clang_getRangeStart(r)
			var file C.CXFile
			var line C.uint
			var column C.uint
			C.clang_getFileLocation(location, &file, &line, &column, nil)
			cxFileName := C.clang_getFileName(file)
			filename := C.GoString(C.clang_getCString(cxFileName))
			C.clang_disposeString(cxFileName)

			var framework string

			if strings.HasPrefix(filename, p.frameworksPath) {
				framework = strings.TrimSuffix(strings.Split(strings.TrimPrefix(filename, p.frameworksPath), "/")[1], ".framework")
			} else {
				framework = "objc"
			}

			loc := &Loc{
				Framework: framework,
				Filename:  filename,
				Line:      uint(line),
				Column:    uint(column),
			}

			if loc.Framework != currentFramework {
				if err := p.ReadFramework(loc.Framework); err != nil {
					log.Println("error importing framework", loc.Framework+":", err)
				}
			}

			hash := C.clang_hashCursor(cursor)
			var parentHash C.uint
			if parent.kind != 0 {
				parentHash = C.clang_hashCursor(parent)
			}

			accessSpecifier := C.clang_getCXXAccessSpecifier(cursor)
			if accessSpecifier != C.CX_CXXPublic && accessSpecifier != C.CX_CXXInvalidAccessSpecifier {
				return C.CXChildVisit_Continue
			}

			if kind == C.CXCursor_TypedefDecl {
				// cxCursorName := C.clang_getCursorDisplayName(cursor)
				// cursorName := C.GoString(C.clang_getCString(cxCursorName))
				// C.clang_disposeString(cxCursorName)
				// cxType := C.clang_getTypedefDeclUnderlyingType(cursor)

				return C.CXChildVisit_Recurse
			} else if kind == C.CXCursor_StructDecl {
				var cursorName string
				var isAnonymous bool

				if C.clang_Cursor_isAnonymousRecordDecl(cursor) != 0 {
					isAnonymous = true
				} else {
					var cxCursorName C.CXString
					if parent.kind == C.CXCursor_TypedefDecl {
						cxCursorName = C.clang_getCursorDisplayName(parent)
					} else {
						cxCursorName = C.clang_getCursorDisplayName(cursor)
					}
					cursorName = C.GoString(C.clang_getCString(cxCursorName))
					C.clang_disposeString(cxCursorName)
				}

				ccursor := C.clang_getCanonicalCursor(cursor)
				var canonical bool
				if C.clang_hashCursor(ccursor) == hash {
					canonical = true
				}

				struct_ := &Struct{
					Loc:         loc,
					Name:        cursorName,
					Canonical:   canonical,
					IsAnonymous: isAnonymous,
				}

				context.structs[uint(hash)] = struct_

				return C.CXChildVisit_Recurse
			} else if kind == C.CXCursor_ObjCInterfaceDecl {
				cxCursorName := C.clang_getCursorDisplayName(cursor)
				cursorName := C.GoString(C.clang_getCString(cxCursorName))
				C.clang_disposeString(cxCursorName)

				// hash := C.clang_hashCursor(cxTypeCursor)

				for k, iface := range context.interfaces {
					if iface.Name == cursorName {
						delete(context.interfaces, k)
						break
					}
				}

				iface := &Interface{
					Loc:  loc,
					Name: cursorName,
				}

				context.interfaces[uint(hash)] = iface

				return C.CXChildVisit_Recurse
			} else if kind == C.CXCursor_ObjCClassRef {
				cxCursorName := C.clang_getCursorDisplayName(cursor)
				cursorName := C.GoString(C.clang_getCString(cxCursorName))
				C.clang_disposeString(cxCursorName)

				seen := false
				for _, iface := range context.interfaces {
					if iface.Name == cursorName {
						seen = true
						break
					}
				}

				if !seen {
					iface := &Interface{
						Loc:         loc,
						Name:        cursorName,
						ForwardDecl: true,
					}

					context.interfaces[uint(hash)] = iface
				}

				return C.CXChildVisit_Recurse
			} else if kind == C.CXCursor_ObjCProtocolDecl {
				cxCursorName := C.clang_getCursorDisplayName(cursor)
				cursorName := C.GoString(C.clang_getCString(cxCursorName))
				C.clang_disposeString(cxCursorName)

				protocol := &Protocol{
					Loc:  loc,
					Name: cursorName,
				}

				context.protocols[uint(hash)] = protocol

				return C.CXChildVisit_Recurse
			} else if context.phase == VisitorPhaseDetail {
				if kind == C.CXCursor_ObjCProtocolRef {
					cxCursorName := C.clang_getCursorDisplayName(cursor)
					protocolName := C.GoString(C.clang_getCString(cxCursorName))
					C.clang_disposeString(cxCursorName)

					var protocol *Protocol
					for _, p := range context.protocols {
						if p.Name == protocolName {
							protocol = p
						}
					}

					iface := context.interfaces[uint(parentHash)]

					if protocol != nil && iface != nil {
						iface.Protocols = append(iface.Protocols, protocol)
					}
				} else if kind == C.CXCursor_ObjCSuperClassRef {
					cxCursorType := C.clang_getCursorType(cursor)
					cxSuperTypeSpelling := C.clang_getTypeSpelling(cxCursorType)
					superType := C.GoString(C.clang_getCString(cxSuperTypeSpelling))
					C.clang_disposeString(cxSuperTypeSpelling)

					found := false
					for _, iface := range context.interfaces {
						if iface.Name == superType {
							found = true
							context.interfaces[uint(parentHash)].Super = iface
						}
					}

					if !found {
						panic("ERR: previous declaration for superclass interface not found")
					}

				} else if kind == C.CXCursor_ObjCInstanceMethodDecl {
					cxCursorName := C.clang_getCursorDisplayName(cursor)
					cursorName := C.GoString(C.clang_getCString(cxCursorName))
					C.clang_disposeString(cxCursorName)

					cxResultType := C.clang_getCursorResultType(cursor)

					method := &Method{
						Loc:        loc,
						Instance:   true,
						Name:       cursorName,
						Interface:  context.interfaces[uint(parentHash)],
						ReturnType: context.ResolveType(cxResultType),
					}

					context.methods[uint(hash)] = method

					if iface, ok := context.interfaces[uint(parentHash)]; ok {
						iface.Methods = append(iface.Methods, method)
					} else if protocol, ok := context.protocols[uint(parentHash)]; ok {
						protocol.Methods = append(protocol.Methods, method)
					}

					return C.CXChildVisit_Recurse
				} else if kind == C.CXCursor_ObjCClassMethodDecl {
					cxCursorName := C.clang_getCursorDisplayName(cursor)
					cursorName := C.GoString(C.clang_getCString(cxCursorName))
					C.clang_disposeString(cxCursorName)

					cxResultType := C.clang_getCursorResultType(cursor)

					method := &Method{
						Loc:        loc,
						Instance:   false,
						Name:       cursorName,
						ReturnType: context.ResolveType(cxResultType),
					}

					if iface, ok := method.ReturnType.(*Interface); ok {
						if err := p.ReadFramework(iface.Loc.Framework); err != nil {
							log.Println(err)
						}
					} else if protocol, ok := method.ReturnType.(*Protocol); ok {
						if err := p.ReadFramework(protocol.Loc.Framework); err != nil {
							log.Println(err)
						}
					} else if st, ok := method.ReturnType.(*Struct); ok {
						if err := p.ReadFramework(st.Loc.Framework); err != nil {
							log.Println(err)
						}
					}

					context.methods[uint(hash)] = method

					if iface, ok := context.interfaces[uint(parentHash)]; ok {
						iface.Methods = append(iface.Methods, method)
					} else if protocol, ok := context.protocols[uint(parentHash)]; ok {
						protocol.Methods = append(protocol.Methods, method)
					}
					return C.CXChildVisit_Recurse
				} else if kind == C.CXCursor_ParmDecl {
					cxCursorName := C.clang_getCursorDisplayName(cursor)
					cursorName := C.GoString(C.clang_getCString(cxCursorName))
					C.clang_disposeString(cxCursorName)

					cxCursorType := C.clang_getCursorType(cursor)

					parameter := &Parameter{
						Loc:  loc,
						Name: cursorName,
						Type: context.ResolveType(cxCursorType),
					}

					if iface, ok := parameter.Type.(*Interface); ok {
						if err := p.ReadFramework(iface.Loc.Framework); err != nil {
							log.Println(err)
						}
					} else if protocol, ok := parameter.Type.(*Protocol); ok {
						if err := p.ReadFramework(protocol.Loc.Framework); err != nil {
							log.Println(err)
						}
					} else if st, ok := parameter.Type.(*Struct); ok {
						if err := p.ReadFramework(st.Loc.Framework); err != nil {
							log.Println(err)
						}
					}

					if method, ok := context.methods[uint(parentHash)]; ok {
						method.Parameters = append(method.Parameters, parameter)
					}
				} else if kind == C.CXCursor_FieldDecl {
					cxCursorName := C.clang_getCursorDisplayName(cursor)
					cursorName := C.GoString(C.clang_getCString(cxCursorName))
					C.clang_disposeString(cxCursorName)

					cxCursorType := C.clang_getCursorType(cursor)

					field := &Field{
						Loc:  loc,
						Name: cursorName,
						Type: context.ResolveType(cxCursorType),
					}

					if iface, ok := field.Type.(*Interface); ok {
						if err := p.ReadFramework(iface.Loc.Framework); err != nil {
							log.Println(err)
						}
					} else if protocol, ok := field.Type.(*Protocol); ok {
						if err := p.ReadFramework(protocol.Loc.Framework); err != nil {
							log.Println(err)
						}
					} else if st, ok := field.Type.(*Struct); ok {
						if err := p.ReadFramework(st.Loc.Framework); err != nil {
							log.Println(err)
						}
					}

					if st, ok := context.structs[uint(parentHash)]; ok {
						st.Fields = append(st.Fields, field)
					}
				}
			}

			return C.CXChildVisit_Continue
		},
	}

	var pinner runtime.Pinner
	pinner.Pin(context)
	defer pinner.Unpin()
	_cgoCheckPointer := func(interface{}, interface{}) {}
	_cgoCheckPointer(nil, nil)

	var index C.CXIndex
	var translationUnit C.CXTranslationUnit
	var cursor C.CXCursor

	C.ParseHeader(target, isysroot, ch, unsafe.Pointer(&index), unsafe.Pointer(&translationUnit), unsafe.Pointer(&cursor))

	context.phase = VisitorPhaseInitial
	C.visitChildren(cursor, unsafe.Pointer(context))

	context.phase = VisitorPhaseDetail
	C.visitChildren(cursor, unsafe.Pointer(context))

	C.clang_disposeTranslationUnit(translationUnit)
	C.clang_disposeIndex(index)

	seen := map[string]bool{}
	for _, st := range context.structs {
		if !st.Canonical || st.Loc.Framework != currentFramework {
			continue
		}

		if seen[st.Name] {
			panic(fmt.Sprintf("ERR: %s %s", "duplicate struct", st.Name))
		}
		seen[st.Name] = true

		if err := p.WriteStruct(st); err != nil {
			return err
		}
	}

	seen = map[string]bool{}
	for _, iface := range context.interfaces {
		if seen[iface.Name] {
			panic(fmt.Sprintf("ERR: %s %s", "duplicate interface", iface.Name))
		}
		seen[iface.Name] = true

		if iface.ForwardDecl || iface.Loc.Framework != currentFramework {
			continue
		}

		if err := p.WriteInterface(iface); err != nil {
			return err
		}
	}

	seen = map[string]bool{}
	for _, protocol := range context.protocols {
		if seen[protocol.Name] {
			panic(fmt.Sprintf("ERR: %s %s", "duplicate protocol", protocol.Name))
		}
		seen[protocol.Name] = true

		if protocol.Loc.Framework != currentFramework {
			continue
		}

		if err := p.WriteProtocol(protocol); err != nil {
			return err
		}
	}

	// log.Println(interfaces, protocols)
	// if b, err := exec.Command("clang", "-target", p.target, "-isysroot", p.sdkPath, "-x", "objective-c", "-fsyntax-only", "-Xclang", "-ast-dump=json", "-E", header).Output(); err != nil {
	// 	return err
	// } else {
	// 	os.WriteFile("ast/"+currentFramework+".json", b, 0644)

	// 	ast := TranslationUnitDecl{
	// 		"loc": map[string]any{
	// 			"file": header,
	// 		},
	// 	}
	// 	json.Unmarshal(b, &ast)

	// 	// for _, s := range ast.Structs() {
	// 	// 	file := (Node)(s).SourceFile()
	// 	// 	if path.Dir(file) == path.Dir(header) {
	// 	// 		if err := p.WriteFile(ast, (Node)(s)); err != nil {
	// 	// 			return err
	// 	// 		}
	// 	// 	}
	// 	// }

	// 	pseen := map[string]bool{}

	// 	for _, protocol := range ast.Protocols() {
	// 		if pseen[protocol.Name()] {
	// 			continue
	// 		}
	// 		pseen[protocol.Name()] = true

	// 		protocols := protocol.Declarations(ast)

	// 		if len(protocols) == 0 {
	// 			continue
	// 		}

	// 		protocol = protocols[0].(ObjCProtocolDeclNode)

	// 		file := Node(protocol).SourceFile()

	// 		if _, ok := protocol["inner"]; !ok {
	// 			continue
	// 		}
	// 		if file == "" || strings.HasPrefix(file, path.Join(p.frameworksPath, currentFramework+".framework")) {
	// 			if err := p.WriteFile(ast, protocols); err != nil {
	// 				return err
	// 			}
	// 		} else if framework != "objc" {
	// 			if err := p.ReadFramework(framework); err != nil {
	// 				log.Println("error importing framework", framework+":", err)
	// 			}
	// 		}
	// 	}

	// 	iseen := map[string]bool{}

	// 	for _, iface := range ast.Interfaces() {
	// 		if iseen[iface.Name()] {
	// 			continue
	// 		}
	// 		iseen[iface.Name()] = true

	// 		interfaces := iface.Declarations(ast)

	// 		if iface.Name() == "EAGLContext" {
	// 			log.Println(interfaces)
	// 		}
	// 		if len(interfaces) == 0 {
	// 			continue
	// 		}

	// 		found := false
	// 		for _, i := range interfaces {
	// 			file := Node(i.(ObjCInterfaceDeclNode)).SourceFile()
	// 			var framework string

	// 			if file != "" {
	// 				if strings.HasPrefix(file, p.frameworksPath) {
	// 					framework = strings.TrimSuffix(strings.Split(strings.TrimPrefix(file, p.frameworksPath), "/")[1], ".framework")
	// 				} else {
	// 					framework = "objc"
	// 				}
	// 			} else {
	// 				framework = "objc"
	// 			}

	// 			if framework == currentFramework {
	// 				log.Println(framework, currentFramework, i.Name(), tapi == nil || tapi.IsExportedClass(i.Name()))
	// 			} else if i.Name() == "EAGLContext" {
	// 				log.Println(i, file, framework)
	// 			}

	// 			if framework == currentFramework && (tapi == nil || tapi.IsExportedClass(i.Name())) {
	// 				if !found {
	// 					if err := p.WriteFile(ast, interfaces); err != nil {
	// 						return err
	// 					}
	// 					found = true
	// 				}
	// 			} else if framework != currentFramework {
	// 				if err := p.ReadFramework(framework); err != nil {
	// 					log.Println("error importing framework", framework+":", err)
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return nil
}

func reHyphenate(in string, sep string) string {
	var out bytes.Buffer
	w := io.StringWriter(&out)

	for i, b := range in {
		if strings.ToUpper(string(b)) == string(b) {
			if i < len(in)-1 && strings.ToLower(string(in[i+1])) == string(in[i+1]) {
				w.WriteString(sep)
			} else if i > 0 && strings.ToLower(string(in[i-1])) == string(in[i-1]) {
				w.WriteString(sep)
			}
		}

		out.WriteString(string(b))
	}

	return strings.ToLower(strings.TrimPrefix(string(out.Bytes()), sep))
}

func unHyphenate(in string, sep string) string {
	parts := strings.Split(in, sep)
	out := make([]string, len(parts))
	for i, p := range parts {
		if len(p) > 0 {
			out[i] = strings.ToUpper(p[0:1]) + p[1:]
		}
	}

	return strings.Join(out, "")
}

func (p *listener) WriteStruct(st *Struct) error {
	wd, _ := os.Getwd()
	target := path.Base(wd)

	filename := fmt.Sprintf("%s/%s.go", reHyphenate(st.Loc.Framework, "_"), st.CanonicalName())
	dirname := path.Dir(filename)

	var hb bytes.Buffer
	var b bytes.Buffer
	var h = io.StringWriter(&hb)
	var w = io.StringWriter(&b)

	h.WriteString(fmt.Sprintf("//js:package native/%s/%s\n", target, reHyphenate(st.Loc.Framework, "-")))
	h.WriteString(fmt.Sprintf("package %s\n\n", reHyphenate(st.Loc.Framework, "_")))
	h.WriteString("//go:generate go run github.com/grexie/isolates/codegen\n\n")

	imports := map[string]string{}

	w.WriteString(fmt.Sprintf("/*\n%s\n*/\n\n", st.String()))

	w.WriteString(fmt.Sprintf("type %s struct {\n", st.CanonicalName()))
	for _, f := range st.Fields {
		if t, err := p.GetType(f.Type, "any", imports, st.Loc.Framework); err != nil {
			return err
		} else {
			name := unHyphenate(f.Name, "_")
			jsName := strings.ToLower(name[0:1]) + name[1:]
			w.WriteString(fmt.Sprintf("  %s %s `v8:\"%s\"`\n", name, t, jsName))
		}
	}
	w.WriteString("}\n")

	if len(imports) > 0 {
		h.WriteString("import (\n")
		for imp, name := range imports {
			if strings.HasSuffix(imp, "/"+name) || imp == name {
				h.WriteString(fmt.Sprintf("  \"%s\"\n", imp))
			} else {
				h.WriteString(fmt.Sprintf("  %s \"%s\"\n", name, imp))
			}
		}
		h.WriteString(")\n\n")
	}

	hb.Write(b.Bytes())
	// log.Println("writing", dirname, filename)
	os.MkdirAll(dirname, 0755)
	os.WriteFile(filename, hb.Bytes(), 0644)

	return nil
}

func (p *listener) WriteInterface(iface *Interface) error {
	wd, _ := os.Getwd()
	target := path.Base(wd)

	filename := fmt.Sprintf("%s/%s.go", reHyphenate(iface.Loc.Framework, "_"), iface.CanonicalName())
	dirname := path.Dir(filename)

	var hb bytes.Buffer
	var b bytes.Buffer
	var h = io.StringWriter(&hb)
	var w = io.StringWriter(&b)

	h.WriteString(fmt.Sprintf("//js:package native/%s/%s\n", target, reHyphenate(iface.Loc.Framework, "-")))
	h.WriteString(fmt.Sprintf("package %s\n\n", reHyphenate(iface.Loc.Framework, "_")))
	h.WriteString("//go:generate go run github.com/grexie/isolates/codegen\n\n")

	var super string

	imports := map[string]string{}

	w.WriteString(fmt.Sprintf("/*\n%s\n*/\n\n", iface.String()))

	if iface.Super != nil {
		if iface.Super.Loc.Framework != iface.Loc.Framework {
			suppackagename := reHyphenate(iface.Super.Loc.Framework, "_")
			suppackage := fmt.Sprintf("github.com/gosolid/solid/pkg/runtime/native/%s/%s", target, suppackagename)
			super = fmt.Sprintf("*%s.%s", suppackagename, iface.Super.CanonicalName())
			imports[suppackage] = suppackagename
		} else {
			super = "*" + iface.Super.CanonicalName()
		}
	}

	protocols := []string{}

	for _, protocol := range iface.Protocols {
		if protocol.Loc.Framework != iface.Loc.Framework {
			protocolpackagename := reHyphenate(protocol.Loc.Framework, "_")
			protocolpackage := fmt.Sprintf("github.com/gosolid/solid/pkg/runtime/native/%s/%s", target, protocolpackagename)
			protocolRef := fmt.Sprintf("  *%s.%s", protocolpackagename, protocol.CanonicalName())
			imports[protocolpackage] = protocolpackagename
			protocols = append(protocols, protocolRef)
		} else {
			protocolRef := "  *" + protocol.CanonicalName()
			protocols = append(protocols, protocolRef)
		}
	}

	w.WriteString(fmt.Sprintf("type %s struct {\n  %s\n%s\n}\n", iface.Name, super, strings.Join(protocols, "\n")))

	methods := map[string][]*Method{}

	for _, m := range iface.Methods {
		methodName := strings.Split(m.Name, ":")[0]
		if _, ok := methods[methodName]; !ok {
			methods[methodName] = []*Method{}
		}
		methods[methodName] = append(methods[methodName], m)
	}

	for _, m := range methods {
		w.WriteString("\n")
		if err := p.WriteInterfaceMethod(iface, m, imports, w); err != nil {
			return err
		}
	}

	w.WriteString("\n")

	if len(imports) > 0 {
		h.WriteString("import (\n")
		for imp, name := range imports {
			if strings.HasSuffix(imp, "/"+name) || imp == name {
				h.WriteString(fmt.Sprintf("  \"%s\"\n", imp))
			} else {
				h.WriteString(fmt.Sprintf("  %s \"%s\"\n", name, imp))
			}
		}
		h.WriteString(")\n\n")
	}

	hb.Write(b.Bytes())
	// log.Println("writing", dirname, filename)
	os.MkdirAll(dirname, 0755)
	os.WriteFile(filename, hb.Bytes(), 0644)

	return nil
}

func (p *listener) WriteProtocol(protocol *Protocol) error {
	wd, _ := os.Getwd()
	target := path.Base(wd)

	filename := fmt.Sprintf("%s/%s.go", reHyphenate(protocol.Loc.Framework, "_"), protocol.CanonicalName())
	dirname := path.Dir(filename)

	var hb bytes.Buffer
	var b bytes.Buffer
	var h = io.StringWriter(&hb)
	var w = io.StringWriter(&b)

	h.WriteString(fmt.Sprintf("//js:package native/%s/%s\n", target, reHyphenate(protocol.Loc.Framework, "-")))
	h.WriteString(fmt.Sprintf("package %s\n\n", reHyphenate(protocol.Loc.Framework, "_")))
	h.WriteString("//go:generate go run github.com/grexie/isolates/codegen\n\n")

	imports := map[string]string{}

	w.WriteString(fmt.Sprintf("/*\n%s\n*/\n\n", protocol.String()))

	protocols := []string{}

	w.WriteString(fmt.Sprintf("type %s struct {\n%s\n}\n", protocol.CanonicalName(), strings.Join(protocols, "\n")))

	methods := map[string][]*Method{}

	for _, m := range protocol.Methods {
		methodName := strings.Split(m.Name, ":")[0]
		if _, ok := methods[methodName]; !ok {
			methods[methodName] = []*Method{}
		}
		methods[methodName] = append(methods[methodName], m)
	}

	for _, m := range methods {
		w.WriteString("\n")
		if err := p.WriteProtocolMethod(protocol, m, imports, w); err != nil {
			return err
		}
	}

	w.WriteString("\n")

	if len(imports) > 0 {
		h.WriteString("import (\n")
		for imp, name := range imports {
			if strings.HasSuffix(imp, "/"+name) || imp == name {
				h.WriteString(fmt.Sprintf("  \"%s\"\n", imp))
			} else {
				h.WriteString(fmt.Sprintf("  %s \"%s\"\n", name, imp))
			}
		}
		h.WriteString(")\n\n")
	}

	hb.Write(b.Bytes())
	// log.Println("writing", dirname, filename)
	os.MkdirAll(dirname, 0755)
	os.WriteFile(filename, hb.Bytes(), 0644)

	return nil
}

var primitiveTypesObjCToGo = map[string]string{
	"void":               "void",
	"id":                 "any",
	"Protocol":           "any",
	"Class":              "any",
	"NSMethodSignature":  "any",
	"SEL":                "any",
	"_NSZone":            "*foundation.NSZone",
	"struct _NSZone":     "*foundation.NSZone",
	"NSInvocation":       "*foundation.NSInvocation",
	"NSString":           "string",
	"func":               "func(...any) (any, error)",
	"BOOL":               "bool",
	"char":               "byte",
	"unsigned char":      "byte",
	"short":              "int16",
	"unsigned short":     "uint16",
	"int":                "int",
	"unsigned int":       "uint",
	"unsigned __int128":  "uint64",
	"signed char":        "byte",
	"long":               "int64",
	"unsigned long":      "uint64",
	"long long":          "int64",
	"unsigned long long": "uint64",
	"float":              "float32",
	"double":             "float64",
}

func (p *listener) GetType(t any, instanceType string, imports map[string]string, currentFramework string) (string, error) {
	if t == nil {
		return "void", nil
	}

	wd, _ := os.Getwd()
	target := path.Base(wd)

	if t == "id" {
		return "any", nil
	} else if t == "instancetype" {
		return instanceType, nil
	} else if ot, ok := t.(string); ok {
		if gt, ok := primitiveTypesObjCToGo[ot]; ok {
			return gt, nil
		} else {
			return "", fmt.Errorf("unknown primitive type: %s", t)
		}
	} else if ptype, ok := t.(*Pointer); ok {
		if innerType, err := p.GetType(ptype.InnerType, instanceType, imports, currentFramework); err != nil {
			return "", err
		} else {
			return fmt.Sprintf("*%s", innerType), nil
		}
	} else if atype, ok := t.(*VariableArray); ok {
		if innerType, err := p.GetType(atype.InnerType, instanceType, imports, currentFramework); err != nil {
			return "", err
		} else {
			return fmt.Sprintf("[]%s", innerType), nil
		}
	} else if atype, ok := t.(*FixedArray); ok {
		if innerType, err := p.GetType(atype.InnerType, instanceType, imports, currentFramework); err != nil {
			return "", err
		} else {
			return fmt.Sprintf("[%d]%s", atype.Length, innerType), nil
		}
	} else if iface, ok := t.(*Interface); ok {
		if iface.Loc.Framework != currentFramework {
			packageName := reHyphenate(iface.Loc.Framework, "_")
			pkg := fmt.Sprintf("github.com/gosolid/solid/pkg/runtime/native/%s/%s", target, packageName)
			returnType := fmt.Sprintf("*%s.%s", packageName, iface.CanonicalName())
			imports[pkg] = packageName
			return returnType, nil
		} else {
			return "*" + iface.CanonicalName(), nil
		}
	} else if protocol, ok := t.(*Protocol); ok {
		if protocol.Loc.Framework != currentFramework {
			packageName := reHyphenate(protocol.Loc.Framework, "_")
			pkg := fmt.Sprintf("github.com/gosolid/solid/pkg/runtime/native/%s/%s", target, packageName)
			returnType := fmt.Sprintf("*%s.%s", packageName, protocol.CanonicalName())
			imports[pkg] = packageName
			return returnType, nil
		} else {
			return "*" + protocol.CanonicalName(), nil
		}
	} else if st, ok := t.(*Struct); ok {
		if st.Loc.Framework != currentFramework {
			packageName := reHyphenate(st.Loc.Framework, "_")
			pkg := fmt.Sprintf("github.com/gosolid/solid/pkg/runtime/native/%s/%s", target, packageName)
			returnType := fmt.Sprintf("%s.%s", packageName, st.CanonicalName())
			imports[pkg] = packageName
			return returnType, nil
		} else {
			return st.CanonicalName(), nil
		}
	} else {
		return "", fmt.Errorf("unknown type: %s", t)
	}
}

var primitiveTypesObjCToGoZero = map[string]string{
	"void":               "nil",
	"id":                 "nil",
	"Protocol":           "nil",
	"Class":              "nil",
	"NSMethodSignature":  "nil",
	"SEL":                "nil",
	"_NSZone":            "nil",
	"struct _NSZone":     "nil",
	"NSInvocation":       "nil",
	"NSString":           "\"\"",
	"func":               "nil",
	"BOOL":               "false",
	"char":               "0",
	"unsigned char":      "0",
	"short":              "0",
	"unsigned short":     "0",
	"int":                "0",
	"unsigned int":       "0",
	"long":               "0",
	"unsigned long":      "0",
	"long long":          "0",
	"unsigned long long": "0",
	"float":              "0",
	"double":             "0",
}

func (p *listener) GetTypeZero(t any, instanceType string, imports map[string]string, currentFramework string) (string, error) {
	if t == nil {
		return "nil", nil
	}

	wd, _ := os.Getwd()
	target := path.Base(wd)

	if t == "id" {
		return "nil", nil
	} else if t == "instancetype" {
		return "nil", nil
	} else if ot, ok := t.(string); ok {
		if gt, ok := primitiveTypesObjCToGoZero[ot]; ok {
			return gt, nil
		} else {
			return "", fmt.Errorf("unknown primitive type: %s", t)
		}
	} else if _, ok := t.(*Pointer); ok {
		return "nil", nil
	} else if _, ok := t.(*VariableArray); ok {
		return "nil", nil
	} else if _, ok := t.(*FixedArray); ok {
		return "nil", nil
	} else if _, ok := t.(*Interface); ok {
		return "nil", nil
	} else if _, ok := t.(*Protocol); ok {
		return "nil", nil
	} else if st, ok := t.(*Struct); ok {
		if st.Loc.Framework != currentFramework {
			packageName := reHyphenate(st.Loc.Framework, "_")
			pkg := fmt.Sprintf("github.com/gosolid/solid/pkg/runtime/native/%s/%s", target, packageName)
			returnType := fmt.Sprintf("%s.%s{}", packageName, st.CanonicalName())
			imports[pkg] = packageName
			return returnType, nil
		} else {
			return st.CanonicalName() + "{}", nil
		}
	} else {
		return "", fmt.Errorf("unknown type: %s", t)
	}

}

func (p *listener) WriteProtocolMethod(protocol *Protocol, m []*Method, imports map[string]string, w io.StringWriter) error {
	methodName := strings.Split(m[0].Name, ":")[0]
	methodName = strings.ToUpper(methodName[0:1]) + methodName[1:]

	imports["fmt"] = "fmt"

	if returnType, err := p.GetType(m[0].ReturnType, protocol.Name, imports, protocol.Loc.Framework); err != nil {
		return err
	} else {
		if returnType == "void" {
			w.WriteString(fmt.Sprintf("func (r *%s) %s() error {\n", protocol.Name, methodName))
		} else {
			w.WriteString(fmt.Sprintf("func (r *%s) %s() (%s, error) {\n", protocol.Name, methodName, returnType))
		}

		if returnType == "void" {
			w.WriteString("  return fmt.Errorf(\"unimplemented\")\n")
		} else if zero, err := p.GetTypeZero(m[0].ReturnType, protocol.Name, imports, protocol.Loc.Framework); err != nil {
			return err
		} else {
			w.WriteString(fmt.Sprintf("  return %s, fmt.Errorf(\"unimplemented\")\n", zero))
		}
		w.WriteString("}\n")
	}

	return nil
}

func (p *listener) WriteInterfaceMethod(iface *Interface, m []*Method, imports map[string]string, w io.StringWriter) error {
	methodName := strings.Split(m[0].Name, ":")[0]
	methodName = strings.ToUpper(methodName[0:1]) + methodName[1:]

	imports["fmt"] = "fmt"

	if returnType, err := p.GetType(m[0].ReturnType, iface.Name, imports, iface.Loc.Framework); err != nil {
		return err
	} else {
		if returnType == "void" {
			w.WriteString(fmt.Sprintf("func (r *%s) %s() error {\n", iface.Name, methodName))
		} else {
			w.WriteString(fmt.Sprintf("func (r *%s) %s() (%s, error) {\n", iface.Name, methodName, returnType))
		}

		if returnType == "void" {
			w.WriteString("  return fmt.Errorf(\"unimplemented\")\n")
		} else if zero, err := p.GetTypeZero(m[0].ReturnType, iface.Name, imports, iface.Loc.Framework); err != nil {
			return err
		} else {
			w.WriteString(fmt.Sprintf("  return %s, fmt.Errorf(\"unimplemented\")\n", zero))
		}
		w.WriteString("}\n")
	}

	return nil
}
