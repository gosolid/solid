//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_PKCS5_PBKDF2_PARAMS_PTR
*/

type CSSMPKCS5PBKDF2PARAMSPTR struct {
  Passphrase SecAsn1Oid `v8:"passphrase"`
  PseudoRandomFunction uint `v8:"pseudoRandomFunction"`
}
