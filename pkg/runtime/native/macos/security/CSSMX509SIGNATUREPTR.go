//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_SIGNATURE_PTR
*/

type CSSMX509SIGNATUREPTR struct {
  AlgorithmIdentifier SecAsn1AlgId `v8:"algorithmIdentifier"`
  Encrypted SecAsn1Oid `v8:"encrypted"`
}
