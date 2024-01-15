//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLPipelineBufferDescriptor : objc.NSObject
*/

type MTLPipelineBufferDescriptor struct {
  *objc.NSObject

}

func (r *MTLPipelineBufferDescriptor) Mutability() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLPipelineBufferDescriptor) SetMutability() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

