//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_RANGE_PTR
*/

type CSSMRANGEPTR struct {
  Min uint `v8:"min"`
  Max uint `v8:"max"`
}
