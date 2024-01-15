//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLBlitPassSampleBufferAttachmentDescriptorArray : objc.NSObject
*/

type MTLBlitPassSampleBufferAttachmentDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLBlitPassSampleBufferAttachmentDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBlitPassSampleBufferAttachmentDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

