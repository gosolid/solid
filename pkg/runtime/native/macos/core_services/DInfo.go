//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.DInfo
*/

type DInfo struct {
  FrRect objc.Rect `v8:"frRect"`
  FrFlags uint16 `v8:"frFlags"`
  FrLocation objc.Point `v8:"frLocation"`
  FrView int16 `v8:"frView"`
}
