//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLComputePassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLComputePassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLComputePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

