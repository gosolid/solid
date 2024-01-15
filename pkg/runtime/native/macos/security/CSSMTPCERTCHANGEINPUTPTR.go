//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTCHANGE_INPUT_PTR
*/

type CSSMTPCERTCHANGEINPUTPTR struct {
  Action uint `v8:"action"`
  Reason uint `v8:"reason"`
  CLHandle int64 `v8:"cLHandle"`
  Cert *SecAsn1Oid `v8:"cert"`
  ChangeInfo *CSSMFIELDPTR `v8:"changeInfo"`
  StartTime *byte `v8:"startTime"`
  CallerCredentials *CSSMACCESSCREDENTIALSPTR `v8:"callerCredentials"`
}
