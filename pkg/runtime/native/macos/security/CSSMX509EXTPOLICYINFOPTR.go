//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509EXT_POLICYINFO_PTR
*/

type CSSMX509EXTPOLICYINFOPTR struct {
  PolicyIdentifier SecAsn1Oid `v8:"policyIdentifier"`
  PolicyQualifiers CSSMX509EXTPOLICYQUALIFIERSPTR `v8:"policyQualifiers"`
}
