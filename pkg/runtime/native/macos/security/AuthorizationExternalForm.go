//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.AuthorizationExternalForm
*/

type AuthorizationExternalForm struct {
  Bytes [32]byte `v8:"bytes"`
}
