//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFXMLParserCallBacks
*/

type CFXMLParserCallBacks struct {
  Version int64 `v8:"version"`
  CreateXMLStructure *func(...any) (any, error) `v8:"createXMLStructure"`
  AddChild *func(...any) (any, error) `v8:"addChild"`
  EndXMLStructure *func(...any) (any, error) `v8:"endXMLStructure"`
  ResolveExternalEntity *func(...any) (any, error) `v8:"resolveExternalEntity"`
  HandleError *func(...any) (any, error) `v8:"handleError"`
}
