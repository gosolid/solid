//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509EXT_POLICYQUALIFIERS_PTR
*/

type CSSMX509EXTPOLICYQUALIFIERSPTR struct {
  NumberOfPolicyQualifiers uint `v8:"numberOfPolicyQualifiers"`
  PolicyQualifier CSSMX509EXTPOLICYQUALIFIERINFOPTR `v8:"policyQualifier"`
}
