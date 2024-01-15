//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRenderPassColorAttachmentDescriptorArray : objc.NSObject
*/

type MTLRenderPassColorAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLRenderPassColorAttachmentDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRenderPassColorAttachmentDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

