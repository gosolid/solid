//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFXMLExternalID
*/

type CFXMLExternalID struct {
  SystemID *CFURL `v8:"systemID"`
  PublicID *CFString `v8:"publicID"`
}
