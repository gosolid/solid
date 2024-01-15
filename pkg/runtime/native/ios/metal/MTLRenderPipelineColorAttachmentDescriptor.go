//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPipelineColorAttachmentDescriptor : objc.NSObject
*/

type MTLRenderPipelineColorAttachmentDescriptor struct {
  *objc.NSObject

}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) IsBlendingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetWriteMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) PixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetBlendingEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) RgbBlendOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) WriteMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

