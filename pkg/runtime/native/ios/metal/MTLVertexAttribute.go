//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLVertexAttribute : objc.NSObject
*/

type MTLVertexAttribute struct {
  *objc.NSObject

}

func (r *MTLVertexAttribute) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) AttributeIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) AttributeType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) IsActive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) IsPatchData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexAttribute) IsPatchControlPointData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

