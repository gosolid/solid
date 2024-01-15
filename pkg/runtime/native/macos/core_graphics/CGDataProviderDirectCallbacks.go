//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.CGDataProviderDirectCallbacks
*/

type CGDataProviderDirectCallbacks struct {
  Version uint `v8:"version"`
  GetBytePointer *func(...any) (any, error) `v8:"getBytePointer"`
  ReleaseBytePointer *func(...any) (any, error) `v8:"releaseBytePointer"`
  GetBytesAtPosition *func(...any) (any, error) `v8:"getBytesAtPosition"`
  ReleaseInfo *func(...any) (any, error) `v8:"releaseInfo"`
}
