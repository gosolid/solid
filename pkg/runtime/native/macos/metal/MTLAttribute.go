//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface Metal.MTLAttribute : objc.NSObject
*/

type MTLAttribute struct {
  *objc.NSObject

}

func (r *MTLAttribute) IsPatchControlPointData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) AttributeIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) AttributeType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) IsActive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLAttribute) IsPatchData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

