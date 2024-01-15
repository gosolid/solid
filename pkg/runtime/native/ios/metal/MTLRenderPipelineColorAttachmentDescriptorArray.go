//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPipelineColorAttachmentDescriptorArray : objc.NSObject
*/

type MTLRenderPipelineColorAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLRenderPipelineColorAttachmentDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPipelineColorAttachmentDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

