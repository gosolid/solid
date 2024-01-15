//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (*MTLAccelerationStructurePassSampleBufferAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

