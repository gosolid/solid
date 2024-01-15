//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.SecAsn1Template
*/

type SecAsn1Template struct {
  Kind uint `v8:"kind"`
  Offset uint `v8:"offset"`
  Sub void `v8:"sub"`
  Size uint `v8:"size"`
}
