//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ENCODED_CRL_PTR
*/

type CSSMENCODEDCRLPTR struct {
  CrlType uint `v8:"crlType"`
  CrlEncoding uint `v8:"crlEncoding"`
  CrlBlob SecAsn1Oid `v8:"crlBlob"`
}
