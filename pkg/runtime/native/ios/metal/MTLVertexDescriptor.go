//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLVertexDescriptor : objc.NSObject
*/

type MTLVertexDescriptor struct {
  *objc.NSObject

}

func (r *MTLVertexDescriptor) Layouts() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexDescriptor) Attributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexDescriptor) VertexDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLVertexDescriptor) Reset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

