//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLBufferLayoutDescriptor : objc.NSObject
*/

type MTLBufferLayoutDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLBufferLayoutDescriptor) SetStepFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) StepRate() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) SetStepRate() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) Stride() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) SetStride() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBufferLayoutDescriptor) StepFunction() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

