//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ENCODED_CERT_PTR
*/

type CSSMENCODEDCERTPTR struct {
  CertType uint `v8:"certType"`
  CertEncoding uint `v8:"certEncoding"`
  CertBlob SecAsn1Oid `v8:"certBlob"`
}
