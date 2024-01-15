//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.SecAsn1AlgId
*/

type SecAsn1AlgId struct {
  Algorithm SecAsn1Oid `v8:"algorithm"`
  Parameters SecAsn1Oid `v8:"parameters"`
}
