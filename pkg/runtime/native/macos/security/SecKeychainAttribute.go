//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.SecKeychainAttribute
*/

type SecKeychainAttribute struct {
  Tag uint `v8:"tag"`
  Length uint `v8:"length"`
  Data void `v8:"data"`
}
