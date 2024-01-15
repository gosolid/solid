//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLAccelerationStructureSizes
*/

type MTLAccelerationStructureSizes struct {
  AccelerationStructureSize uint64 `v8:"accelerationStructureSize"`
  BuildScratchBufferSize uint64 `v8:"buildScratchBufferSize"`
  RefitScratchBufferSize uint64 `v8:"refitScratchBufferSize"`
}
