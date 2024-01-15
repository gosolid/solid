//js:package native/ios/metal
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

func (r *MTLRenderPassColorAttachmentDescriptor) ClearColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassColorAttachmentDescriptor) SetClearColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

