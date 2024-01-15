//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSSortDescriptor : objc.NSObject
*/

type NSSortDescriptor struct {
  *objc.NSObject

}

func (r *NSSortDescriptor) SortDescriptorWithKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) Key() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) Selector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) InitWithKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) AllowEvaluation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) CompareObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) Ascending() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) Comparator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSortDescriptor) ReversedSortDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

