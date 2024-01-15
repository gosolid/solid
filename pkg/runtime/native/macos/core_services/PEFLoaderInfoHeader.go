//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.PEFLoaderInfoHeader
*/

type PEFLoaderInfoHeader struct {
  MainSection int `v8:"mainSection"`
  MainOffset uint `v8:"mainOffset"`
  InitSection int `v8:"initSection"`
  InitOffset uint `v8:"initOffset"`
  TermSection int `v8:"termSection"`
  TermOffset uint `v8:"termOffset"`
  ImportedLibraryCount uint `v8:"importedLibraryCount"`
  TotalImportedSymbolCount uint `v8:"totalImportedSymbolCount"`
  RelocSectionCount uint `v8:"relocSectionCount"`
  RelocInstrOffset uint `v8:"relocInstrOffset"`
  LoaderStringsOffset uint `v8:"loaderStringsOffset"`
  ExportHashOffset uint `v8:"exportHashOffset"`
  ExportHashTablePower uint `v8:"exportHashTablePower"`
  ExportedSymbolCount uint `v8:"exportedSymbolCount"`
}
