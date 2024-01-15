//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLQuadTessellationFactorsHalf
*/

type MTLQuadTessellationFactorsHalf struct {
  EdgeTessellationFactor [4]uint16 `v8:"edgeTessellationFactor"`
  InsideTessellationFactor [2]uint16 `v8:"insideTessellationFactor"`
}
