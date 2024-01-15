//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTextContentStorage : AppKit.NSTextContentManager
*/

type NSTextContentStorage struct {
  *NSTextContentManager
  *NSTextStorageObserving
}

func (r *NSTextContentStorage) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) SetAttributedString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) AttributedStringForTextElement() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) TextElementForAttributedString() (*NSTextElement, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) LocationFromLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) AdjustedRangeFromRange() (*NSTextRange, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) OffsetFromLocation() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) AttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

