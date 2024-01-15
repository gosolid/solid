//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyStateTerminators
*/

type UCKeyStateTerminators struct {
  KeyStateTerminatorsFormat uint16 `v8:"keyStateTerminatorsFormat"`
  KeyStateTerminatorCount uint16 `v8:"keyStateTerminatorCount"`
  KeyStateTerminators [1]uint16 `v8:"keyStateTerminators"`
}
