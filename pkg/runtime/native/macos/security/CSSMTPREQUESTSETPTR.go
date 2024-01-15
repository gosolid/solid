//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_REQUEST_SET_PTR
*/

type CSSMTPREQUESTSETPTR struct {
  NumberOfRequests uint `v8:"numberOfRequests"`
  Requests void `v8:"requests"`
}
