//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_RDN_PTR
*/

type CSSMX509RDNPTR struct {
  NumberOfPairs uint `v8:"numberOfPairs"`
  AttributeTypeAndValue *CSSMX509TYPEVALUEPAIRPTR `v8:"attributeTypeAndValue"`
}
