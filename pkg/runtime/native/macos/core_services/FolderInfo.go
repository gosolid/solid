//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.FolderInfo
*/

type FolderInfo struct {
  WindowBounds objc.Rect `v8:"windowBounds"`
  FinderFlags uint16 `v8:"finderFlags"`
  Location objc.Point `v8:"location"`
  ReservedField uint16 `v8:"reservedField"`
}
