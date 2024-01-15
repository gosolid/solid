//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLResourceStatePassSampleBufferAttachmentDescriptor : objc.NSObject
*/

type MTLResourceStatePassSampleBufferAttachmentDescriptor struct {
  *objc.NSObject

}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) SampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) SetSampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResourceStatePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

