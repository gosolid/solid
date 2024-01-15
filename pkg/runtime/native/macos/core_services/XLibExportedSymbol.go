//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.XLibExportedSymbol
*/

type XLibExportedSymbol struct {
  ClassAndName uint `v8:"classAndName"`
  BpOffset uint `v8:"bpOffset"`
}
