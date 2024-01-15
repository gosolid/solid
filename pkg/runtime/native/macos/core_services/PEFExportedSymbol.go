//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.PEFExportedSymbol
*/

type PEFExportedSymbol struct {
  ClassAndName uint `v8:"classAndName"`
  SymbolValue uint `v8:"symbolValue"`
  SectionIndex int16 `v8:"sectionIndex"`
}
