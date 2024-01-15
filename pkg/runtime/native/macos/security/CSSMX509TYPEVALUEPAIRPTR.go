//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_TYPE_VALUE_PAIR_PTR
*/

type CSSMX509TYPEVALUEPAIRPTR struct {
  Type SecAsn1Oid `v8:"type"`
  ValueType byte `v8:"valueType"`
  Value SecAsn1Oid `v8:"value"`
}
