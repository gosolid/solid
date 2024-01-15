//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLStructMember : objc.NSObject
*/

type MTLStructMember struct {
  *objc.NSObject

}

func (r *MTLStructMember) TextureReferenceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) PointerType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) Offset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) DataType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) ArgumentIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) StructType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) ArrayType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

