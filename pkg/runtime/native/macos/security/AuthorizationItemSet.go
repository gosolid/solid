//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.AuthorizationItemSet
*/

type AuthorizationItemSet struct {
  Count uint `v8:"count"`
  Items AuthorizationItem `v8:"items"`
}
