//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CRLISSUE_INPUT_PTR
*/

type CSSMTPCRLISSUEINPUTPTR struct {
  CLHandle int64 `v8:"cLHandle"`
  CrlIdentifier uint `v8:"crlIdentifier"`
  CrlThisTime *byte `v8:"crlThisTime"`
  PolicyIdentifier *CSSMFIELDPTR `v8:"policyIdentifier"`
  CallerCredentials *CSSMACCESSCREDENTIALSPTR `v8:"callerCredentials"`
}
