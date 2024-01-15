//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTISSUE_INPUT_PTR
*/

type CSSMTPCERTISSUEINPUTPTR struct {
  CSPSubserviceUid CSSMSUBSERVICEUIDPTR `v8:"cSPSubserviceUid"`
  CLHandle int64 `v8:"cLHandle"`
  NumberOfTemplateFields uint `v8:"numberOfTemplateFields"`
  SubjectCertFields *CSSMFIELDPTR `v8:"subjectCertFields"`
  MoreServiceRequests uint `v8:"moreServiceRequests"`
  NumberOfServiceControls uint `v8:"numberOfServiceControls"`
  ServiceControls *CSSMFIELDPTR `v8:"serviceControls"`
  UserCredentials *CSSMACCESSCREDENTIALSPTR `v8:"userCredentials"`
}
