//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.AuthorizationItem
*/

type AuthorizationItem struct {
  Name *byte `v8:"name"`
  ValueLength uint64 `v8:"valueLength"`
  Value void `v8:"value"`
  Flags uint `v8:"flags"`
}
