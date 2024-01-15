//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLVertexBufferLayoutDescriptor : objc.NSObject
*/

type MTLVertexBufferLayoutDescriptor struct {
  *objc.NSObject

}

func (r *MTLVertexBufferLayoutDescriptor) StepFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) SetStepFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) StepRate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) SetStepRate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) Stride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexBufferLayoutDescriptor) SetStride() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

