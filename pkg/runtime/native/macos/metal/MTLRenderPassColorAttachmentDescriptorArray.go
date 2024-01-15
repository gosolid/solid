//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPassColorAttachmentDescriptorArray : objc.NSObject
*/

type MTLRenderPassColorAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLRenderPassColorAttachmentDescriptorArray) ObjectAtIndexedSubscript() (*MTLRenderPassColorAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassColorAttachmentDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

