//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.ExtendedFolderInfo
*/

type ExtendedFolderInfo struct {
  ScrollPosition objc.Point `v8:"scrollPosition"`
  Reserved1 int `v8:"reserved1"`
  ExtendedFinderFlags uint16 `v8:"extendedFinderFlags"`
  Reserved2 int16 `v8:"reserved2"`
  PutAwayFolderID int `v8:"putAwayFolderID"`
}
