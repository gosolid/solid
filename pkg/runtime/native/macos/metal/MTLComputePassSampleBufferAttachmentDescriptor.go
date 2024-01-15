//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLComputePassSampleBufferAttachmentDescriptor : objc.NSObject
*/

type MTLComputePassSampleBufferAttachmentDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) SampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) SetSampleBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLComputePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

