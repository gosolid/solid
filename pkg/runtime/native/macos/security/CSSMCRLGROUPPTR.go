//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CRLGROUP_PTR
*/

type CSSMCRLGROUPPTR struct {
  CrlType uint `v8:"crlType"`
  CrlEncoding uint `v8:"crlEncoding"`
  NumberOfCrls uint `v8:"numberOfCrls"`
  GroupCrlList void `v8:"groupCrlList"`
  CrlGroupType uint `v8:"crlGroupType"`
}
