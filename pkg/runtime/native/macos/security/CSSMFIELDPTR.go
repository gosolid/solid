//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_FIELD_PTR
*/

type CSSMFIELDPTR struct {
  FieldOid SecAsn1Oid `v8:"fieldOid"`
  FieldValue SecAsn1Oid `v8:"fieldValue"`
}
