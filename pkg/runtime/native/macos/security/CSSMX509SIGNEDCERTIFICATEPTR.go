//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_SIGNED_CERTIFICATE_PTR
*/

type CSSMX509SIGNEDCERTIFICATEPTR struct {
  Certificate CSSMX509TBSCERTIFICATEPTR `v8:"certificate"`
  Signature CSSMX509SIGNATUREPTR `v8:"signature"`
}
