//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_GeneralSubtrees
*/

type CEGeneralSubtrees struct {
  NumSubtrees uint `v8:"numSubtrees"`
  Subtrees CEGeneralSubtree `v8:"subtrees"`
}
