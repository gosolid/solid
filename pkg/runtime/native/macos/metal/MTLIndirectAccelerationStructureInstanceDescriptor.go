//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLIndirectAccelerationStructureInstanceDescriptor
*/

type MTLIndirectAccelerationStructureInstanceDescriptor struct {
  TransformationMatrix MTLPackedFloat4x3 `v8:"transformationMatrix"`
  Options int `v8:"options"`
  Mask uint `v8:"mask"`
  IntersectionFunctionTableOffset uint `v8:"intersectionFunctionTableOffset"`
  UserID uint `v8:"userID"`
  AccelerationStructureID MTLResourceID `v8:"accelerationStructureID"`
}
