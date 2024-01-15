//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyStateRecord
*/

type UCKeyStateRecord struct {
  StateZeroCharData uint16 `v8:"stateZeroCharData"`
  StateZeroNextState uint16 `v8:"stateZeroNextState"`
  StateEntryCount uint16 `v8:"stateEntryCount"`
  StateEntryFormat uint16 `v8:"stateEntryFormat"`
  StateEntryData [1]uint `v8:"stateEntryData"`
}
