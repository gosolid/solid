//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLResourceStatePassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLResourceStatePassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

