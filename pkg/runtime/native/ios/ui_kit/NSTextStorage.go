//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/foundation"
  "fmt"
)

/*
interface UIKit.NSTextStorage : Foundation.NSMutableAttributedString
*/

type NSTextStorage struct {
  *foundation.NSMutableAttributedString

}

func (r *NSTextStorage) AddLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) ProcessEditing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) ChangeInLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) FixesAttributesLazily() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) RemoveLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) Edited() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) InvalidateAttributesInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) LayoutManagers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) EditedMask() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) TextStorageObserver() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) EnsureAttributesAreFixedInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) EditedRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) SetTextStorageObserver() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

