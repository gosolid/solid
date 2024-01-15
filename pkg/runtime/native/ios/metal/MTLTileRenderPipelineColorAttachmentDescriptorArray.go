//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLTileRenderPipelineColorAttachmentDescriptorArray : objc.NSObject
*/

type MTLTileRenderPipelineColorAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLTileRenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLTileRenderPipelineColorAttachmentDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

