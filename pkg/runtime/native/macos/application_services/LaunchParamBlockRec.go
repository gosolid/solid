//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_services"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.LaunchParamBlockRec
*/

type LaunchParamBlockRec struct {
  Reserved1 uint `v8:"reserved1"`
  Reserved2 uint16 `v8:"reserved2"`
  LaunchBlockID uint16 `v8:"launchBlockID"`
  LaunchEPBLength uint `v8:"launchEPBLength"`
  LaunchFileFlags uint16 `v8:"launchFileFlags"`
  LaunchControlFlags uint16 `v8:"launchControlFlags"`
  LaunchAppRef *core_services.FSRef `v8:"launchAppRef"`
  LaunchProcessSN objc.ProcessSerialNumber `v8:"launchProcessSN"`
  LaunchPreferredSize uint `v8:"launchPreferredSize"`
  LaunchMinimumSize uint `v8:"launchMinimumSize"`
  LaunchAvailableSize uint `v8:"launchAvailableSize"`
  LaunchAppParameters *AppParameters `v8:"launchAppParameters"`
}
