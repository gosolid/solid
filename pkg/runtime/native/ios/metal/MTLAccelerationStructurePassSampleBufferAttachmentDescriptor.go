//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLAccelerationStructurePassSampleBufferAttachmentDescriptor : objc.NSObject
*/

type MTLAccelerationStructurePassSampleBufferAttachmentDescriptor struct {
  *objc.NSObject

}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SetSampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

