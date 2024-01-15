//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.PEFImportedLibrary
*/

type PEFImportedLibrary struct {
  NameOffset uint `v8:"nameOffset"`
  OldImpVersion uint `v8:"oldImpVersion"`
  CurrentVersion uint `v8:"currentVersion"`
  ImportedSymbolCount uint `v8:"importedSymbolCount"`
  FirstImportedSymbol uint `v8:"firstImportedSymbol"`
  Options byte `v8:"options"`
  ReservedA byte `v8:"reservedA"`
  ReservedB uint16 `v8:"reservedB"`
}
