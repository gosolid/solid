//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_PolicyQualifierInfo
*/

type CEPolicyQualifierInfo struct {
  PolicyQualifierId SecAsn1Oid `v8:"policyQualifierId"`
  Qualifier SecAsn1Oid `v8:"qualifier"`
}
