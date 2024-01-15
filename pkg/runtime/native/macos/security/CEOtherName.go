//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_OtherName
*/

type CEOtherName struct {
  TypeId SecAsn1Oid `v8:"typeId"`
  Value SecAsn1Oid `v8:"value"`
}
