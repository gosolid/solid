//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

