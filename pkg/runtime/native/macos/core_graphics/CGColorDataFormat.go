//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.CGColorDataFormat
*/

type CGColorDataFormat struct {
  Version uint `v8:"version"`
  ColorspaceInfo *void `v8:"colorspaceInfo"`
  BitmapInfo int `v8:"bitmapInfo"`
  BitsPerComponent uint64 `v8:"bitsPerComponent"`
  BytesPerRow uint64 `v8:"bytesPerRow"`
  Intent int `v8:"intent"`
  Decode float64 `v8:"decode"`
}
