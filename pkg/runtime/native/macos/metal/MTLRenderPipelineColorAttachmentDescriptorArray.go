//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPipelineColorAttachmentDescriptorArray : objc.NSObject
*/

type MTLRenderPipelineColorAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLRenderPipelineColorAttachmentDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript() (*MTLRenderPipelineColorAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

