//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_APPLEDL_OPEN_PARAMETERS_PTR
*/

type CSSMAPPLEDLOPENPARAMETERSPTR struct {
  Length uint `v8:"length"`
  Version uint `v8:"version"`
  AutoCommit int `v8:"autoCommit"`
  Mask uint `v8:"mask"`
  Mode uint16 `v8:"mode"`
}
