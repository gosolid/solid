//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLAccelerationStructurePassDescriptor : objc.NSObject
*/

type MTLAccelerationStructurePassDescriptor struct {
  *objc.NSObject

}

func (r *MTLAccelerationStructurePassDescriptor) AccelerationStructurePassDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassDescriptor) SampleBufferAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

