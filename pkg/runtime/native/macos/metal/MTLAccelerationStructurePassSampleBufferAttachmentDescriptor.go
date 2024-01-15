//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLAccelerationStructurePassSampleBufferAttachmentDescriptor : objc.NSObject
*/

type MTLAccelerationStructurePassSampleBufferAttachmentDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SetSampleBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

