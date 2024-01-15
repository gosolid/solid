//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_PolicyInformation
*/

type CEPolicyInformation struct {
  CertPolicyId SecAsn1Oid `v8:"certPolicyId"`
  NumPolicyQualifiers uint `v8:"numPolicyQualifiers"`
  PolicyQualifiers CEPolicyQualifierInfo `v8:"policyQualifiers"`
}
