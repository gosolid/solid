//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLVertexAmplificationViewMapping
*/

type MTLVertexAmplificationViewMapping struct {
  ViewportArrayIndexOffset uint `v8:"viewportArrayIndexOffset"`
  RenderTargetArrayIndexOffset uint `v8:"renderTargetArrayIndexOffset"`
}
