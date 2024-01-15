//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.CGScreenUpdateMoveDelta
*/

type CGScreenUpdateMoveDelta struct {
  DX int `v8:"dX"`
  DY int `v8:"dY"`
}
