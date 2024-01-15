//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLOrigin
*/

type MTLOrigin struct {
  X uint64 `v8:"x"`
  Y uint64 `v8:"y"`
  Z uint64 `v8:"z"`
}
