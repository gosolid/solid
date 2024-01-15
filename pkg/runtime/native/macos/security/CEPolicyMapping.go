//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_PolicyMapping
*/

type CEPolicyMapping struct {
  IssuerDomainPolicy SecAsn1Oid `v8:"issuerDomainPolicy"`
  SubjectDomainPolicy SecAsn1Oid `v8:"subjectDomainPolicy"`
}
