//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.AuthorizationValueVector
*/

type AuthorizationValueVector struct {
  Count uint `v8:"count"`
  Values AuthorizationValue `v8:"values"`
}
