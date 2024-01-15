//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLComputePassDescriptor : objc.NSObject
*/

type MTLComputePassDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLComputePassDescriptor) ComputePassDescriptor() (*MTLComputePassDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassDescriptor) DispatchType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassDescriptor) SetDispatchType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePassDescriptor) SampleBufferAttachments() (*MTLComputePassSampleBufferAttachmentDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

