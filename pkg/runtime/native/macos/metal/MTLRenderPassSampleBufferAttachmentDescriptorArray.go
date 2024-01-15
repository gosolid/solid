//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLRenderPassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLRenderPassSampleBufferAttachmentDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (*MTLRenderPassSampleBufferAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

