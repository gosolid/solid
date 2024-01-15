//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLTileRenderPipelineColorAttachmentDescriptorArray : objc.NSObject
*/

type MTLTileRenderPipelineColorAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLTileRenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript() (*MTLTileRenderPipelineColorAttachmentDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineColorAttachmentDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

