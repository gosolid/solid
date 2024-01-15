//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.NSTextContentStorage : UIKit.NSTextContentManager
*/

type NSTextContentStorage struct {
  *NSTextContentManager

}

func (r *NSTextContentStorage) SetAttributedString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) LocationFromLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) AttributedString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) OffsetFromLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) AdjustedRangeFromRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) AttributedStringForTextElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextContentStorage) TextElementForAttributedString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

