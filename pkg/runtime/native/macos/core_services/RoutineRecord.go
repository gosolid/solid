//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.RoutineRecord
*/

type RoutineRecord struct {
  ProcInfo uint64 `v8:"procInfo"`
  Reserved1 byte `v8:"reserved1"`
  ISA byte `v8:"iSA"`
  RoutineFlags uint16 `v8:"routineFlags"`
  ProcDescriptor *func(...any) (any, error) `v8:"procDescriptor"`
  Reserved2 uint `v8:"reserved2"`
  Selector uint `v8:"selector"`
}
