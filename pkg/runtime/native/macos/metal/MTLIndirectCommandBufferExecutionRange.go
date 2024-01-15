//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLIndirectCommandBufferExecutionRange
*/

type MTLIndirectCommandBufferExecutionRange struct {
  Location uint `v8:"location"`
  Length uint `v8:"length"`
}
