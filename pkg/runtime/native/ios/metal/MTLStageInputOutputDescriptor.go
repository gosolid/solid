//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLStageInputOutputDescriptor : objc.NSObject
*/

type MTLStageInputOutputDescriptor struct {
  *objc.NSObject

}

func (r *MTLStageInputOutputDescriptor) Reset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) Layouts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) Attributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) IndexType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) SetIndexType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) IndexBufferIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) SetIndexBufferIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStageInputOutputDescriptor) StageInputOutputDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

