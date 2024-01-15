//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLBufferLayoutDescriptor : objc.NSObject
*/

type MTLBufferLayoutDescriptor struct {
  *objc.NSObject

}

func (r *MTLBufferLayoutDescriptor) Stride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) SetStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) StepFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) SetStepFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) StepRate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) SetStepRate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

