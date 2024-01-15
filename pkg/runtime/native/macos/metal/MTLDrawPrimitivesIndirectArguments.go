//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLDrawPrimitivesIndirectArguments
*/

type MTLDrawPrimitivesIndirectArguments struct {
  VertexCount uint `v8:"vertexCount"`
  InstanceCount uint `v8:"instanceCount"`
  VertexStart uint `v8:"vertexStart"`
  BaseInstance uint `v8:"baseInstance"`
}
