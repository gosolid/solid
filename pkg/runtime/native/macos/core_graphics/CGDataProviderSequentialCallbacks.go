//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.CGDataProviderSequentialCallbacks
*/

type CGDataProviderSequentialCallbacks struct {
  Version uint `v8:"version"`
  GetBytes *func(...any) (any, error) `v8:"getBytes"`
  SkipForward *func(...any) (any, error) `v8:"skipForward"`
  Rewind *func(...any) (any, error) `v8:"rewind"`
  ReleaseInfo *func(...any) (any, error) `v8:"releaseInfo"`
}
