//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_PKCS5_PBKDF1_PARAMS_PTR
*/

type CSSMPKCS5PBKDF1PARAMSPTR struct {
  Passphrase SecAsn1Oid `v8:"passphrase"`
  InitVector SecAsn1Oid `v8:"initVector"`
}
