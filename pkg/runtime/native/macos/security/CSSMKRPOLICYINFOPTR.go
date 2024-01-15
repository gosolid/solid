//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_KR_POLICY_INFO_PTR
*/

type CSSMKRPOLICYINFOPTR struct {
  KrbNotAllowed int `v8:"krbNotAllowed"`
  NumberOfEntries uint `v8:"numberOfEntries"`
  PolicyEntry CSSMKRPOLICYLISTITEMPTR `v8:"policyEntry"`
}
