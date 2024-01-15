//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLCounterResultStageUtilization
*/

type MTLCounterResultStageUtilization struct {
  TotalCycles uint64 `v8:"totalCycles"`
  VertexCycles uint64 `v8:"vertexCycles"`
  TessellationCycles uint64 `v8:"tessellationCycles"`
  PostTessellationVertexCycles uint64 `v8:"postTessellationVertexCycles"`
  FragmentCycles uint64 `v8:"fragmentCycles"`
  RenderTargetCycles uint64 `v8:"renderTargetCycles"`
}
