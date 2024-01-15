//js:package native/ios/metal
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

func (r *MTLRenderPassStencilAttachmentDescriptor) ClearStencil() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassStencilAttachmentDescriptor) SetClearStencil() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassStencilAttachmentDescriptor) StencilResolveFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassStencilAttachmentDescriptor) SetStencilResolveFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

