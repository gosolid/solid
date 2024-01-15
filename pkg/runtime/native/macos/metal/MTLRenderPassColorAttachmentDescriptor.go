//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Metal.MTLRenderPassColorAttachmentDescriptor : Metal.MTLRenderPassAttachmentDescriptor
*/

type MTLRenderPassColorAttachmentDescriptor struct {
  *MTLRenderPassAttachmentDescriptor

}

func (r *MTLRenderPassColorAttachmentDescriptor) ClearColor() (MTLClearColor, error) {
  return MTLClearColor{}, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassColorAttachmentDescriptor) SetClearColor() error {
  return fmt.Errorf("unimplemented")
}

