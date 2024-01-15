//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.RoutineDescriptor
*/

type RoutineDescriptor struct {
  GoMixedModeTrap uint16 `v8:"goMixedModeTrap"`
  Version byte `v8:"version"`
  RoutineDescriptorFlags byte `v8:"routineDescriptorFlags"`
  Reserved1 uint `v8:"reserved1"`
  Reserved2 byte `v8:"reserved2"`
  SelectorInfo byte `v8:"selectorInfo"`
  RoutineCount uint16 `v8:"routineCount"`
  RoutineRecords [1]RoutineRecord `v8:"routineRecords"`
}
