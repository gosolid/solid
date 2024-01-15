//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TableDirectoryRecord
*/

type TableDirectoryRecord struct {
  TableSignature uint `v8:"tableSignature"`
  Reserved uint `v8:"reserved"`
  TableStartOffset uint `v8:"tableStartOffset"`
  TableSize uint `v8:"tableSize"`
}
