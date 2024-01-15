//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLResourceStatePassSampleBufferAttachmentDescriptor : objc.NSObject
*/

type MTLResourceStatePassSampleBufferAttachmentDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) SetSampleBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) SampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

