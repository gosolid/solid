//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSIndexPath : objc.NSObject
*/

type NSIndexPath struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSIndexPath) GetIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) Compare() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) Length() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexPathWithIndexes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) InitWithIndexes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) InitWithIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexAtPosition() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexPathWithIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexPathByAddingIndex() (*NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexPath) IndexPathByRemovingLastIndex() (*NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

