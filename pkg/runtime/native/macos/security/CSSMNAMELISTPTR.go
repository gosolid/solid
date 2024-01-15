//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_NAME_LIST_PTR
*/

type CSSMNAMELISTPTR struct {
  NumStrings uint `v8:"numStrings"`
  String byte `v8:"string"`
}
