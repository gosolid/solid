//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509EXT_POLICYQUALIFIERINFO_PTR
*/

type CSSMX509EXTPOLICYQUALIFIERINFOPTR struct {
  PolicyQualifierId SecAsn1Oid `v8:"policyQualifierId"`
  Value SecAsn1Oid `v8:"value"`
}
