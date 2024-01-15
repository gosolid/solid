//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.AuthorizationValue
*/

type AuthorizationValue struct {
  Length uint64 `v8:"length"`
  Data void `v8:"data"`
}
