//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTVERIFY_INPUT_PTR
*/

type CSSMTPCERTVERIFYINPUTPTR struct {
  CLHandle int64 `v8:"cLHandle"`
  Cert *SecAsn1Oid `v8:"cert"`
  VerifyContext *CSSMTPVERIFYCONTEXTPTR `v8:"verifyContext"`
}
