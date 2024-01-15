//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509_EXTENSION_PTR
*/

type CSSMX509EXTENSIONPTR struct {
  ExtnId SecAsn1Oid `v8:"extnId"`
  Critical int `v8:"critical"`
  Format int `v8:"format"`
  Value void `v8:"value"`
  BERvalue SecAsn1Oid `v8:"bERvalue"`
}
