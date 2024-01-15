//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct ApplicationServices.PixMap
*/

type PixMap struct {
  BaseAddr *byte `v8:"baseAddr"`
  RowBytes int16 `v8:"rowBytes"`
  Bounds objc.Rect `v8:"bounds"`
  PmVersion int16 `v8:"pmVersion"`
  PackType int16 `v8:"packType"`
  PackSize int `v8:"packSize"`
  HRes int `v8:"hRes"`
  VRes int `v8:"vRes"`
  PixelType int16 `v8:"pixelType"`
  PixelSize int16 `v8:"pixelSize"`
  CmpCount int16 `v8:"cmpCount"`
  CmpSize int16 `v8:"cmpSize"`
  PixelFormat uint `v8:"pixelFormat"`
  PmTable **ColorTable `v8:"pmTable"`
  PmExt void `v8:"pmExt"`
}
