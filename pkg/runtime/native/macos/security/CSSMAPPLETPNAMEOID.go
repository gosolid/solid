//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLE_TP_NAME_OID
*/

type CSSMAPPLETPNAMEOID struct {
  String byte `v8:"string"`
  Oid SecAsn1Oid `v8:"oid"`
}
