//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_GeneralNames
*/

type CEGeneralNames struct {
  NumNames uint `v8:"numNames"`
  GeneralName CEGeneralName `v8:"generalName"`
}
