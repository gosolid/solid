//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLRenderPassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLRenderPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassSampleBufferAttachmentDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

