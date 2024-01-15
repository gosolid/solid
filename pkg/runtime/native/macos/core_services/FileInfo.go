//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.FileInfo
*/

type FileInfo struct {
  FileType uint `v8:"fileType"`
  FileCreator uint `v8:"fileCreator"`
  FinderFlags uint16 `v8:"finderFlags"`
  Location objc.Point `v8:"location"`
  ReservedField uint16 `v8:"reservedField"`
}
