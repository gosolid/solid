//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CE_DataAndType
*/

type CEDataAndType struct {
  Type int `v8:"type"`
  Extension void `v8:"extension"`
  Critical int `v8:"critical"`
}
