//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_TP_POLICYINFO_PTR
*/

type CSSMTPPOLICYINFOPTR struct {
  NumberOfPolicyIds uint `v8:"numberOfPolicyIds"`
  PolicyIds *CSSMFIELDPTR `v8:"policyIds"`
  PolicyControl void `v8:"policyControl"`
}
