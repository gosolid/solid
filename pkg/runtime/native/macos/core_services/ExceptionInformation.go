//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ExceptionInformation
*/

type ExceptionInformation struct {
  TheKind uint64 `v8:"theKind"`
  MachineState MachineInformation `v8:"machineState"`
  RegisterImage RegisterInformation `v8:"registerImage"`
  FPUImage FPUInformation `v8:"fPUImage"`
  Info void `v8:"info"`
  VectorImage VectorInformation `v8:"vectorImage"`
}
