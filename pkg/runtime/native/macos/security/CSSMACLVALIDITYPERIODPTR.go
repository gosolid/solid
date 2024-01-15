//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ACL_VALIDITY_PERIOD_PTR
*/

type CSSMACLVALIDITYPERIODPTR struct {
  StartDate SecAsn1Oid `v8:"startDate"`
  EndDate SecAsn1Oid `v8:"endDate"`
}
