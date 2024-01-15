//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FXInfo
*/

type FXInfo struct {
  FdIconID int16 `v8:"fdIconID"`
  FdReserved [3]int16 `v8:"fdReserved"`
  FdScript byte `v8:"fdScript"`
  FdXFlags byte `v8:"fdXFlags"`
  FdComment int16 `v8:"fdComment"`
  FdPutAway int `v8:"fdPutAway"`
}
