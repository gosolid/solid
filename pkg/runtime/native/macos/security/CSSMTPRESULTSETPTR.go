//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_RESULT_SET_PTR
*/

type CSSMTPRESULTSETPTR struct {
  NumberOfResults uint `v8:"numberOfResults"`
  Results void `v8:"results"`
}
