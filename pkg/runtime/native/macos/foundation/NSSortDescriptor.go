//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSSortDescriptor : objc.NSObject
*/

type NSSortDescriptor struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
}

func (r *NSSortDescriptor) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) AllowEvaluation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) Ascending() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) Key() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) Selector() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) Comparator() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) ReversedSortDescriptor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) SortDescriptorWithKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) InitWithKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) CompareObject() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

