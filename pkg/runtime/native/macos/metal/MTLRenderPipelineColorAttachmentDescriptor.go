//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLRenderPipelineColorAttachmentDescriptor : objc.NSObject
*/

type MTLRenderPipelineColorAttachmentDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) PixelFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) IsBlendingEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetAlphaBlendOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetWriteMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetDestinationRGBBlendFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SourceAlphaBlendFactor() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetSourceAlphaBlendFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetDestinationAlphaBlendFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetBlendingEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetSourceRGBBlendFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) DestinationRGBBlendFactor() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) RgbBlendOperation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) WriteMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SourceRGBBlendFactor() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) SetRgbBlendOperation() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) DestinationAlphaBlendFactor() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptor) AlphaBlendOperation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

