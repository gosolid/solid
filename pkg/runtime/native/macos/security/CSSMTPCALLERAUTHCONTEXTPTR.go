//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CALLERAUTH_CONTEXT_PTR
*/

type CSSMTPCALLERAUTHCONTEXTPTR struct {
  Policy CSSMTPPOLICYINFOPTR `v8:"policy"`
  VerifyTime *byte `v8:"verifyTime"`
  VerificationAbortOn uint `v8:"verificationAbortOn"`
  CallbackWithVerifiedCert *func(...any) (any, error) `v8:"callbackWithVerifiedCert"`
  NumberOfAnchorCerts uint `v8:"numberOfAnchorCerts"`
  AnchorCerts *SecAsn1Oid `v8:"anchorCerts"`
  DBList *CSSMDLDBLISTPTR `v8:"dBList"`
  CallerCredentials *CSSMACCESSCREDENTIALSPTR `v8:"callerCredentials"`
}
