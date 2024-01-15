//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.SecAsn1PubKeyInfo
*/

type SecAsn1PubKeyInfo struct {
  Algorithm SecAsn1AlgId `v8:"algorithm"`
  SubjectPublicKey SecAsn1Oid `v8:"subjectPublicKey"`
}
