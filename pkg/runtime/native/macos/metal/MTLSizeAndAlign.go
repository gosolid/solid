//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLSizeAndAlign
*/

type MTLSizeAndAlign struct {
  Size uint64 `v8:"size"`
  Align uint64 `v8:"align"`
}
