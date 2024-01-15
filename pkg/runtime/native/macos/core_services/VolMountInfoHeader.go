//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.VolMountInfoHeader
*/

type VolMountInfoHeader struct {
  Length int16 `v8:"length"`
  Media uint `v8:"media"`
}
