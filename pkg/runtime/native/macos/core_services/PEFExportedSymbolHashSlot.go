//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.PEFExportedSymbolHashSlot
*/

type PEFExportedSymbolHashSlot struct {
  CountAndStart uint `v8:"countAndStart"`
}
