//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLResourceStatePassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLResourceStatePassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (*MTLResourceStatePassSampleBufferAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

