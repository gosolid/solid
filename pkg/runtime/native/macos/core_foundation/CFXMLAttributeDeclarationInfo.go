//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFXMLAttributeDeclarationInfo
*/

type CFXMLAttributeDeclarationInfo struct {
  AttributeName *CFString `v8:"attributeName"`
  TypeString *CFString `v8:"typeString"`
  DefaultString *CFString `v8:"defaultString"`
}
