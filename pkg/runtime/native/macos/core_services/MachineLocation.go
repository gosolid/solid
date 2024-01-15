//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MachineLocation
*/

type MachineLocation struct {
  Latitude int `v8:"latitude"`
  Longitude int `v8:"longitude"`
  U void `v8:"u"`
}
