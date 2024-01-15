//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLResourceStatePassDescriptor : objc.NSObject
*/

type MTLResourceStatePassDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLResourceStatePassDescriptor) ResourceStatePassDescriptor() (*MTLResourceStatePassDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassDescriptor) SampleBufferAttachments() (*MTLResourceStatePassSampleBufferAttachmentDescriptorArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

