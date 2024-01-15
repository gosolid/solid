//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_VERIFY_CONTEXT_PTR
*/

type CSSMTPVERIFYCONTEXTPTR struct {
  Action uint `v8:"action"`
  ActionData SecAsn1Oid `v8:"actionData"`
  Crls CSSMCRLGROUPPTR `v8:"crls"`
  Cred *CSSMTPCALLERAUTHCONTEXTPTR `v8:"cred"`
}
