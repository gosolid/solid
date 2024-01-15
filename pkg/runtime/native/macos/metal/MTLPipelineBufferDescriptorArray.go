//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLPipelineBufferDescriptorArray : objc.NSObject
*/

type MTLPipelineBufferDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLPipelineBufferDescriptorArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLPipelineBufferDescriptorArray) ObjectAtIndexedSubscript() (*MTLPipelineBufferDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

