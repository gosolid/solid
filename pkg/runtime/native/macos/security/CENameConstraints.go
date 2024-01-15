//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_NameConstraints
*/

type CENameConstraints struct {
  Permitted CEGeneralSubtrees `v8:"permitted"`
  Excluded CEGeneralSubtrees `v8:"excluded"`
}
