//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSIndexPath : objc.NSObject
*/

type NSIndexPath struct {
  *objc.NSObject

}

func (r *NSIndexPath) InitWithIndexes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexPathByAddingIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexAtPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) GetIndexes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexPathWithIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexPathWithIndexes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) InitWithIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexPathByRemovingLastIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) Compare() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) Length() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

