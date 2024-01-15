//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLSamplePosition
*/

type MTLSamplePosition struct {
  X float32 `v8:"x"`
  Y float32 `v8:"y"`
}
