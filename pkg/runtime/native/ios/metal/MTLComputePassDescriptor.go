//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLComputePassDescriptor : objc.NSObject
*/

type MTLComputePassDescriptor struct {
  *objc.NSObject

}

func (r *MTLComputePassDescriptor) SetDispatchType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassDescriptor) SampleBufferAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassDescriptor) ComputePassDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassDescriptor) DispatchType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

