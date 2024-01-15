//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CERTNOTARIZE_OUTPUT_PTR
*/

type CSSMTPCERTNOTARIZEOUTPUTPTR struct {
  NotarizeStatus uint `v8:"notarizeStatus"`
  NotarizedCertGroup *CSSMCERTGROUPPTR `v8:"notarizedCertGroup"`
  PerformedServiceRequests uint `v8:"performedServiceRequests"`
}
