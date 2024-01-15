//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLTriangleTessellationFactorsHalf
*/

type MTLTriangleTessellationFactorsHalf struct {
  EdgeTessellationFactor [3]uint16 `v8:"edgeTessellationFactor"`
  InsideTessellationFactor uint16 `v8:"insideTessellationFactor"`
}
