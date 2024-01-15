//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSTextRange : objc.NSObject
*/

type NSTextRange struct {
  *objc.NSObject

}

func (r *NSTextRange) InitWithLocation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) IsEqualToTextRange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) ContainsLocation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) ContainsRange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) IntersectsWithTextRange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) TextRangeByIntersectingWithTextRange() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) IsEmpty() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) Location() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) TextRangeByFormingUnionWithTextRange() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextRange) EndLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

