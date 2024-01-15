//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct CoreServices.LSItemInfoRecord
*/

type LSItemInfoRecord struct {
  Flags int `v8:"flags"`
  Filetype uint `v8:"filetype"`
  Creator uint `v8:"creator"`
  Extension *core_foundation.CFString `v8:"extension"`
}
