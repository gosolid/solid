//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FVector
*/

type FVector struct {
  Start int16 `v8:"start"`
  Length int16 `v8:"length"`
}
