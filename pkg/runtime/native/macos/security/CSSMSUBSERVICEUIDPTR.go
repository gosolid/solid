//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SUBSERVICE_UID_PTR
*/

type CSSMSUBSERVICEUIDPTR struct {
  Guid CSSMGUIDPTR `v8:"guid"`
  Version CSSMVERSIONPTR `v8:"version"`
  SubserviceId uint `v8:"subserviceId"`
  SubserviceType uint `v8:"subserviceType"`
}
