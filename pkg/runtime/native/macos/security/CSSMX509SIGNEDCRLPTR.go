//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_SIGNED_CRL_PTR
*/

type CSSMX509SIGNEDCRLPTR struct {
  TbsCertList CSSMX509TBSCERTLISTPTR `v8:"tbsCertList"`
  Signature CSSMX509SIGNATUREPTR `v8:"signature"`
}
