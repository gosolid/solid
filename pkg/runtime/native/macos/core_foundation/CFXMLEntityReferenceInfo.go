//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFXMLEntityReferenceInfo
*/

type CFXMLEntityReferenceInfo struct {
  EntityType int `v8:"entityType"`
}
