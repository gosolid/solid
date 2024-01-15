//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLStructMember : objc.NSObject
*/

type MTLStructMember struct {
  *objc.NSObject

}

func (r *MTLStructMember) ArgumentIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) StructType() (*MTLStructType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) ArrayType() (*MTLArrayType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) TextureReferenceType() (*MTLTextureReferenceType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) PointerType() (*MTLPointerType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) Offset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLStructMember) DataType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

