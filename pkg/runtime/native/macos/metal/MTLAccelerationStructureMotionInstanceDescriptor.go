//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLAccelerationStructureMotionInstanceDescriptor
*/

type MTLAccelerationStructureMotionInstanceDescriptor struct {
  Options int `v8:"options"`
  Mask uint `v8:"mask"`
  IntersectionFunctionTableOffset uint `v8:"intersectionFunctionTableOffset"`
  AccelerationStructureIndex uint `v8:"accelerationStructureIndex"`
  UserID uint `v8:"userID"`
  MotionTransformsStartIndex uint `v8:"motionTransformsStartIndex"`
  MotionTransformsCount uint `v8:"motionTransformsCount"`
  MotionStartBorderMode int `v8:"motionStartBorderMode"`
  MotionEndBorderMode int `v8:"motionEndBorderMode"`
  MotionStartTime float32 `v8:"motionStartTime"`
  MotionEndTime float32 `v8:"motionEndTime"`
}
