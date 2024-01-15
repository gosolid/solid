//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_ExtendedKeyUsage
*/

type CEExtendedKeyUsage struct {
  NumPurposes uint `v8:"numPurposes"`
  Purposes *SecAsn1Oid `v8:"purposes"`
}
