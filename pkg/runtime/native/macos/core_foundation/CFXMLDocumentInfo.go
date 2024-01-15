//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFXMLDocumentInfo
*/

type CFXMLDocumentInfo struct {
  SourceURL *CFURL `v8:"sourceURL"`
  Encoding uint `v8:"encoding"`
}
