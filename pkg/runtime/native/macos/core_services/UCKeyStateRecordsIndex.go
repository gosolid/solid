//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyStateRecordsIndex
*/

type UCKeyStateRecordsIndex struct {
  KeyStateRecordsIndexFormat uint16 `v8:"keyStateRecordsIndexFormat"`
  KeyStateRecordCount uint16 `v8:"keyStateRecordCount"`
  KeyStateRecordOffsets [1]uint `v8:"keyStateRecordOffsets"`
}
