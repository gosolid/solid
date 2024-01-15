//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLRenderPassSampleBufferAttachmentDescriptor : objc.NSObject
*/

type MTLRenderPassSampleBufferAttachmentDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) StartOfVertexSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetStartOfVertexSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetEndOfVertexSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetSampleBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) EndOfVertexSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) StartOfFragmentSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetStartOfFragmentSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) EndOfFragmentSampleIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetEndOfFragmentSampleIndex() error {
  return fmt.Errorf("unimplemented")
}

