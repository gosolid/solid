//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLBlitPassSampleBufferAttachmentDescriptor : objc.NSObject
*/

type MTLBlitPassSampleBufferAttachmentDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLBlitPassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitPassSampleBufferAttachmentDescriptor) SampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBlitPassSampleBufferAttachmentDescriptor) SetSampleBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitPassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLBlitPassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBlitPassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

