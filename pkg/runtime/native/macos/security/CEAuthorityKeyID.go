//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_AuthorityKeyID
*/

type CEAuthorityKeyID struct {
  KeyIdentifierPresent int `v8:"keyIdentifierPresent"`
  KeyIdentifier SecAsn1Oid `v8:"keyIdentifier"`
  GeneralNamesPresent int `v8:"generalNamesPresent"`
  GeneralNames CEGeneralNames `v8:"generalNames"`
  SerialNumberPresent int `v8:"serialNumberPresent"`
  SerialNumber SecAsn1Oid `v8:"serialNumber"`
}
