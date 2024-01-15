//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIEditMenuConfiguration : objc.NSObject
*/

type UIEditMenuConfiguration struct {
  *objc.NSObject

}

func (r *UIEditMenuConfiguration) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuConfiguration) SourcePoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuConfiguration) PreferredArrowDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuConfiguration) SetPreferredArrowDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuConfiguration) ConfigurationWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuConfiguration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIEditMenuConfiguration) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

