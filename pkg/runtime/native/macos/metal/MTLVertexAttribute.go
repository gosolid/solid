//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLVertexAttribute : objc.NSObject
*/

type MTLVertexAttribute struct {
  *objc.NSObject

}

func (r *MTLVertexAttribute) AttributeIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) AttributeType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) IsActive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) IsPatchData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) IsPatchControlPointData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

