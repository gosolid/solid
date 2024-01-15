//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPasteConfiguration : objc.NSObject
*/

type UIPasteConfiguration struct {
  *objc.NSObject

}

func (r *UIPasteConfiguration) InitWithAcceptableTypeIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteConfiguration) AddAcceptableTypeIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteConfiguration) InitWithTypeIdentifiersForAcceptingClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteConfiguration) AddTypeIdentifiersForAcceptingClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteConfiguration) AcceptableTypeIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteConfiguration) SetAcceptableTypeIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPasteConfiguration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

