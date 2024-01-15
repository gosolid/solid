//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLComputePassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLComputePassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLComputePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (*MTLComputePassSampleBufferAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

