//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFXMLElementInfo
*/

type CFXMLElementInfo struct {
  Attributes *CFDictionary `v8:"attributes"`
  AttributeOrder *CFArray `v8:"attributeOrder"`
  IsEmpty byte `v8:"isEmpty"`
  Reserved [3]byte `v8:"reserved"`
}
