//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.CGDeviceColor
*/

type CGDeviceColor struct {
  Red float32 `v8:"red"`
  Green float32 `v8:"green"`
  Blue float32 `v8:"blue"`
}
