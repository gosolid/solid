//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLBlitPassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLBlitPassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLBlitPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (*MTLBlitPassSampleBufferAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBlitPassSampleBufferAttachmentDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

