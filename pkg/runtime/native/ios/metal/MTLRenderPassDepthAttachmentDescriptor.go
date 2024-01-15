//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLRenderPassDepthAttachmentDescriptor : Metal.MTLRenderPassAttachmentDescriptor
*/

type MTLRenderPassDepthAttachmentDescriptor struct {
  *MTLRenderPassAttachmentDescriptor

}

func (r *MTLRenderPassDepthAttachmentDescriptor) ClearDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDepthAttachmentDescriptor) SetClearDepth() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDepthAttachmentDescriptor) DepthResolveFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDepthAttachmentDescriptor) SetDepthResolveFilter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

