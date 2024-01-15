//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.CGFunctionCallbacks
*/

type CGFunctionCallbacks struct {
  Version uint `v8:"version"`
  Evaluate *func(...any) (any, error) `v8:"evaluate"`
  ReleaseInfo *func(...any) (any, error) `v8:"releaseInfo"`
}
