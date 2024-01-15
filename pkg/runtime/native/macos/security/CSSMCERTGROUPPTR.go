//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_CERTGROUP_PTR
*/

type CSSMCERTGROUPPTR struct {
  CertType uint `v8:"certType"`
  CertEncoding uint `v8:"certEncoding"`
  NumCerts uint `v8:"numCerts"`
  GroupList void `v8:"groupList"`
  CertGroupType uint `v8:"certGroupType"`
  Reserved void `v8:"reserved"`
}
