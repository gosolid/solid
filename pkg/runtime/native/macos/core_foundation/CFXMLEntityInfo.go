//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFXMLEntityInfo
*/

type CFXMLEntityInfo struct {
  EntityType int `v8:"entityType"`
  ReplacementText *CFString `v8:"replacementText"`
  EntityID CFXMLExternalID `v8:"entityID"`
  NotationName *CFString `v8:"notationName"`
}
