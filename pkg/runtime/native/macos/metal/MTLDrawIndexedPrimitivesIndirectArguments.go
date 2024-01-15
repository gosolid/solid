//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLDrawIndexedPrimitivesIndirectArguments
*/

type MTLDrawIndexedPrimitivesIndirectArguments struct {
  IndexCount uint `v8:"indexCount"`
  InstanceCount uint `v8:"instanceCount"`
  IndexStart uint `v8:"indexStart"`
  BaseVertex int `v8:"baseVertex"`
  BaseInstance uint `v8:"baseInstance"`
}
