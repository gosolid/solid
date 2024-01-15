//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_NAME_PTR
*/

type CSSMX509NAMEPTR struct {
  NumberOfRDNs uint `v8:"numberOfRDNs"`
  RelativeDistinguishedName *CSSMX509RDNPTR `v8:"relativeDistinguishedName"`
}
