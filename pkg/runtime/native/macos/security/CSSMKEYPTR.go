//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_KEY_PTR
*/

type CSSMKEYPTR struct {
  KeyHeader CSSMKEYHEADERPTR `v8:"keyHeader"`
  KeyData SecAsn1Oid `v8:"keyData"`
}
