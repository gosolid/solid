//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLComputePassSampleBufferAttachmentDescriptor : objc.NSObject
*/

type MTLComputePassSampleBufferAttachmentDescriptor struct {
  *objc.NSObject

}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) SetSampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) SampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

