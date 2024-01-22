//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.CGPatternCallbacks
*/

type CGPatternCallbacks struct {
  Version uint `v8:"version"`
  DrawPattern *func(...any) (any, error) `v8:"drawPattern"`
  ReleaseInfo *func(...any) (any, error) `v8:"releaseInfo"`
}