//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.FInfo
*/

type FInfo struct {
  FdType uint `v8:"fdType"`
  FdCreator uint `v8:"fdCreator"`
  FdFlags uint16 `v8:"fdFlags"`
  FdLocation objc.Point `v8:"fdLocation"`
  FdFldr int16 `v8:"fdFldr"`
}
