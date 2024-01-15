//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIBarButtonItemAppearance : objc.NSObject
*/

type UIBarButtonItemAppearance struct {
  *objc.NSObject

}

func (r *UIBarButtonItemAppearance) Copy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemAppearance) ConfigureWithDefaultForStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemAppearance) Normal() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemAppearance) Highlighted() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemAppearance) Disabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemAppearance) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemAppearance) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemAppearance) Focused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarButtonItemAppearance) InitWithStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

