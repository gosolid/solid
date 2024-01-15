//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLStageInRegionIndirectArguments
*/

type MTLStageInRegionIndirectArguments struct {
  StageInOrigin [3]uint `v8:"stageInOrigin"`
  StageInSize [3]uint `v8:"stageInSize"`
}
