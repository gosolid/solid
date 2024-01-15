//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MemoryExceptionInformation
*/

type MemoryExceptionInformation struct {
  TheArea *OpaqueAreaID `v8:"theArea"`
  TheAddress *void `v8:"theAddress"`
  TheError int `v8:"theError"`
  TheReference uint64 `v8:"theReference"`
}
