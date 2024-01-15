//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLCounterResultStatistic
*/

type MTLCounterResultStatistic struct {
  TessellationInputPatches uint64 `v8:"tessellationInputPatches"`
  VertexInvocations uint64 `v8:"vertexInvocations"`
  PostTessellationVertexInvocations uint64 `v8:"postTessellationVertexInvocations"`
  ClipperInvocations uint64 `v8:"clipperInvocations"`
  ClipperPrimitivesOut uint64 `v8:"clipperPrimitivesOut"`
  FragmentInvocations uint64 `v8:"fragmentInvocations"`
  FragmentsPassed uint64 `v8:"fragmentsPassed"`
  ComputeKernelInvocations uint64 `v8:"computeKernelInvocations"`
}
