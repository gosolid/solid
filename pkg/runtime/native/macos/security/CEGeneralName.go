//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_GeneralName
*/

type CEGeneralName struct {
  NameType int `v8:"nameType"`
  BerEncoded int `v8:"berEncoded"`
  Name SecAsn1Oid `v8:"name"`
}
