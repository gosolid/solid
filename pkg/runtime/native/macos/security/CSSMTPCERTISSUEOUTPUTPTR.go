//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTISSUE_OUTPUT_PTR
*/

type CSSMTPCERTISSUEOUTPUTPTR struct {
  IssueStatus uint `v8:"issueStatus"`
  CertGroup *CSSMCERTGROUPPTR `v8:"certGroup"`
  PerformedServiceRequests uint `v8:"performedServiceRequests"`
}
