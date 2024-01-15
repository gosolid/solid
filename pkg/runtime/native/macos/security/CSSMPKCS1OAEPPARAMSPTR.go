//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_PKCS1_OAEP_PARAMS_PTR
*/

type CSSMPKCS1OAEPPARAMSPTR struct {
  HashAlgorithm uint `v8:"hashAlgorithm"`
  HashParams SecAsn1Oid `v8:"hashParams"`
  MGF uint `v8:"mGF"`
  MGFParams SecAsn1Oid `v8:"mGFParams"`
  PSource uint `v8:"pSource"`
  PSourceParams SecAsn1Oid `v8:"pSourceParams"`
}
