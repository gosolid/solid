//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFXMLAttributeListDeclarationInfo
*/

type CFXMLAttributeListDeclarationInfo struct {
  NumberOfAttributes int64 `v8:"numberOfAttributes"`
  Attributes CFXMLAttributeDeclarationInfo `v8:"attributes"`
}
