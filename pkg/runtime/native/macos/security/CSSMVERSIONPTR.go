//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_VERSION_PTR
*/

type CSSMVERSIONPTR struct {
  Major uint `v8:"major"`
  Minor uint `v8:"minor"`
}
