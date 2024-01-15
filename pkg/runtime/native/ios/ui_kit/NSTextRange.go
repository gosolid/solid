//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSTextRange : objc.NSObject
*/

type NSTextRange struct {
  *objc.NSObject

}

func (r *NSTextRange) TextRangeByFormingUnionWithTextRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) IsEmpty() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) EndLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) ContainsLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) ContainsRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) Location() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) InitWithLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) IsEqualToTextRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) IntersectsWithTextRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) TextRangeByIntersectingWithTextRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

