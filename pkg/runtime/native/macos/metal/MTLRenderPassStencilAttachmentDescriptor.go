//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLRenderPassStencilAttachmentDescriptor : Metal.MTLRenderPassAttachmentDescriptor
*/

type MTLRenderPassStencilAttachmentDescriptor struct {
  *MTLRenderPassAttachmentDescriptor

}

func (r *MTLRenderPassStencilAttachmentDescriptor) SetClearStencil() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassStencilAttachmentDescriptor) StencilResolveFilter() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassStencilAttachmentDescriptor) SetStencilResolveFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassStencilAttachmentDescriptor) ClearStencil() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

