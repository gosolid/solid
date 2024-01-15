//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLScissorRect
*/

type MTLScissorRect struct {
  X uint64 `v8:"x"`
  Y uint64 `v8:"y"`
  Width uint64 `v8:"width"`
  Height uint64 `v8:"height"`
}
