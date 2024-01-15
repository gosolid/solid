//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.BitMap
*/

type BitMap struct {
  BaseAddr *byte `v8:"baseAddr"`
  RowBytes int16 `v8:"rowBytes"`
  Bounds objc.Rect `v8:"bounds"`
}
