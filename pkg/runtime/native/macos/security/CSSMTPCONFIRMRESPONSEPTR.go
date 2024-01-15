//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_CONFIRM_RESPONSE_PTR
*/

type CSSMTPCONFIRMRESPONSEPTR struct {
  NumberOfResponses uint `v8:"numberOfResponses"`
  Responses *uint `v8:"responses"`
}
