//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLAttribute : objc.NSObject
*/

type MTLAttribute struct {
  *objc.NSObject

}

func (r *MTLAttribute) AttributeType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) IsActive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) IsPatchData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) IsPatchControlPointData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) AttributeIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

