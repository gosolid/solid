//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_AccessDescription
*/

type CEAccessDescription struct {
  AccessMethod SecAsn1Oid `v8:"accessMethod"`
  AccessLocation CEGeneralName `v8:"accessLocation"`
}
