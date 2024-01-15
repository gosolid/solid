//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSTextStorage : Foundation.NSMutableAttributedString
*/

type NSTextStorage struct {
  *foundation.NSMutableAttributedString
  *foundation.NSSecureCoding
}

func (r *NSTextStorage) RemoveLayoutManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) Edited() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) LayoutManagers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) EditedMask() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) ChangeInLength() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) TextStorageObserver() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) SetTextStorageObserver() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) InvalidateAttributesInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) AddLayoutManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) EditedRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) ProcessEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) EnsureAttributesAreFixedInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextStorage) FixesAttributesLazily() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

