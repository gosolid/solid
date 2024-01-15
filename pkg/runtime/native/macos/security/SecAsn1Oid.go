//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.SecAsn1Oid
*/

type SecAsn1Oid struct {
  Length uint64 `v8:"length"`
  Data byte `v8:"data"`
}
