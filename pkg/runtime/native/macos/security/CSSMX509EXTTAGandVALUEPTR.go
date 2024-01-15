//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_X509EXT_TAGandVALUE_PTR
*/

type CSSMX509EXTTAGandVALUEPTR struct {
  Type byte `v8:"type"`
  Value SecAsn1Oid `v8:"value"`
}
