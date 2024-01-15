//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTNOTARIZE_INPUT_PTR
*/

type CSSMTPCERTNOTARIZEINPUTPTR struct {
  CLHandle int64 `v8:"cLHandle"`
  NumberOfFields uint `v8:"numberOfFields"`
  MoreFields *CSSMFIELDPTR `v8:"moreFields"`
  SignScope *CSSMFIELDPTR `v8:"signScope"`
  ScopeSize uint `v8:"scopeSize"`
  MoreServiceRequests uint `v8:"moreServiceRequests"`
  NumberOfServiceControls uint `v8:"numberOfServiceControls"`
  ServiceControls *CSSMFIELDPTR `v8:"serviceControls"`
  UserCredentials *CSSMACCESSCREDENTIALSPTR `v8:"userCredentials"`
}
