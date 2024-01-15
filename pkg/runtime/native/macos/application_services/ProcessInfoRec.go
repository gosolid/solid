//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_services"
)

/*
struct ApplicationServices.ProcessInfoRec
*/

type ProcessInfoRec struct {
  ProcessInfoLength uint `v8:"processInfoLength"`
  ProcessName *byte `v8:"processName"`
  ProcessNumber objc.ProcessSerialNumber `v8:"processNumber"`
  ProcessType uint `v8:"processType"`
  ProcessSignature uint `v8:"processSignature"`
  ProcessMode uint `v8:"processMode"`
  ProcessLocation *byte `v8:"processLocation"`
  ProcessSize uint `v8:"processSize"`
  ProcessFreeMem uint `v8:"processFreeMem"`
  ProcessLauncher objc.ProcessSerialNumber `v8:"processLauncher"`
  ProcessLaunchDate uint `v8:"processLaunchDate"`
  ProcessActiveTime uint `v8:"processActiveTime"`
  ProcessAppRef *core_services.FSRef `v8:"processAppRef"`
}
