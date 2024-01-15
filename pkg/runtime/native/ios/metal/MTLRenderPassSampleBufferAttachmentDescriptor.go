//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPassSampleBufferAttachmentDescriptor : objc.NSObject
*/

type MTLRenderPassSampleBufferAttachmentDescriptor struct {
  *objc.NSObject

}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetStartOfVertexSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) StartOfFragmentSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetEndOfFragmentSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetSampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) StartOfVertexSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) EndOfVertexSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetEndOfVertexSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SetStartOfFragmentSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) EndOfFragmentSampleIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptor) SampleBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

