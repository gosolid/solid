//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.FPUInformationPowerPC
*/

type FPUInformationPowerPC struct {
  Registers [32]objc.UnsignedWide `v8:"registers"`
  FPSCR uint64 `v8:"fPSCR"`
  Reserved uint64 `v8:"reserved"`
}
