//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyModifiersToTableNum
*/

type UCKeyModifiersToTableNum struct {
  KeyModifiersToTableNumFormat uint16 `v8:"keyModifiersToTableNumFormat"`
  DefaultTableNum uint16 `v8:"defaultTableNum"`
  ModifiersCount uint `v8:"modifiersCount"`
  TableNum [1]byte `v8:"tableNum"`
}
