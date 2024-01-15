//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_GeneralSubtree
*/

type CEGeneralSubtree struct {
  Base CEGeneralNames `v8:"base"`
  Minimum uint `v8:"minimum"`
  MaximumPresent int `v8:"maximumPresent"`
  Maximum uint `v8:"maximum"`
}
