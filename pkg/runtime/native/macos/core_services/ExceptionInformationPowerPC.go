//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ExceptionInformationPowerPC
*/

type ExceptionInformationPowerPC struct {
  TheKind uint64 `v8:"theKind"`
  MachineState MachineInformationPowerPC `v8:"machineState"`
  RegisterImage RegisterInformationPowerPC `v8:"registerImage"`
  FPUImage FPUInformationPowerPC `v8:"fPUImage"`
  Info void `v8:"info"`
  VectorImage VectorInformationPowerPC `v8:"vectorImage"`
}
