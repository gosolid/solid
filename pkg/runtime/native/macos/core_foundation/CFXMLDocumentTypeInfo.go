//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFXMLDocumentTypeInfo
*/

type CFXMLDocumentTypeInfo struct {
  ExternalID CFXMLExternalID `v8:"externalID"`
}
