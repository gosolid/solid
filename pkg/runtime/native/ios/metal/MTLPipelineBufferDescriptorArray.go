//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLPipelineBufferDescriptorArray : objc.NSObject
*/

type MTLPipelineBufferDescriptorArray struct {
  *objc.NSObject

}

func (r *MTLPipelineBufferDescriptorArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPipelineBufferDescriptorArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

