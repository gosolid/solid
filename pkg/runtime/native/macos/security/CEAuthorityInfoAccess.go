//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_AuthorityInfoAccess
*/

type CEAuthorityInfoAccess struct {
  NumAccessDescriptions uint `v8:"numAccessDescriptions"`
  AccessDescriptions CEAccessDescription `v8:"accessDescriptions"`
}
