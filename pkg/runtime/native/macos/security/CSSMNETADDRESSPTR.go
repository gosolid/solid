//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_NET_ADDRESS_PTR
*/

type CSSMNETADDRESSPTR struct {
  AddressType uint `v8:"addressType"`
  Address SecAsn1Oid `v8:"address"`
}
