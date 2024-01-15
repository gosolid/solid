//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLPipelineBufferDescriptor : objc.NSObject
*/

type MTLPipelineBufferDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLPipelineBufferDescriptor) Mutability() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLPipelineBufferDescriptor) SetMutability() error {
  return fmt.Errorf("unimplemented")
}

