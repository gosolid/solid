//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLDrawPatchIndirectArguments
*/

type MTLDrawPatchIndirectArguments struct {
  PatchCount uint `v8:"patchCount"`
  InstanceCount uint `v8:"instanceCount"`
  PatchStart uint `v8:"patchStart"`
  BaseInstance uint `v8:"baseInstance"`
}
