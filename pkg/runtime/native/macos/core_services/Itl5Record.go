//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.Itl5Record
*/

type Itl5Record struct {
  VersionNumber int `v8:"versionNumber"`
  NumberOfTables uint16 `v8:"numberOfTables"`
  Reserved [3]uint16 `v8:"reserved"`
  TableDirectory [1]TableDirectoryRecord `v8:"tableDirectory"`
}
