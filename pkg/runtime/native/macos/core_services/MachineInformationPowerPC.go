//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.MachineInformationPowerPC
*/

type MachineInformationPowerPC struct {
  CTR objc.UnsignedWide `v8:"cTR"`
  LR objc.UnsignedWide `v8:"lR"`
  PC objc.UnsignedWide `v8:"pC"`
  CRRegister uint64 `v8:"cRRegister"`
  XER uint64 `v8:"xER"`
  MSR uint64 `v8:"mSR"`
  MQ uint64 `v8:"mQ"`
  ExceptKind uint64 `v8:"exceptKind"`
  DSISR uint64 `v8:"dSISR"`
  DAR objc.UnsignedWide `v8:"dAR"`
  Reserved objc.UnsignedWide `v8:"reserved"`
}
