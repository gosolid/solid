//js:package native/macos/metal
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

func (r *MTLRenderPassDepthAttachmentDescriptor) DepthResolveFilter() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDepthAttachmentDescriptor) SetDepthResolveFilter() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDepthAttachmentDescriptor) ClearDepth() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassDepthAttachmentDescriptor) SetClearDepth() error {
  return fmt.Errorf("unimplemented")
}

