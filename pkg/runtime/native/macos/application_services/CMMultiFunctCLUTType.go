//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMultiFunctCLUTType
*/

type CMMultiFunctCLUTType struct {
  GridPoints [16]byte `v8:"gridPoints"`
  EntrySize byte `v8:"entrySize"`
  Reserved [3]byte `v8:"reserved"`
  Data [2]byte `v8:"data"`
}
