//js:package native/macos/core-graphics
package core_graphics

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreGraphics.CGDataConsumerCallbacks
*/

type CGDataConsumerCallbacks struct {
  PutBytes *func(...any) (any, error) `v8:"putBytes"`
  ReleaseConsumer *func(...any) (any, error) `v8:"releaseConsumer"`
}
