//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIContentUnavailableButtonProperties : objc.NSObject
*/

type UIContentUnavailableButtonProperties struct {
  *objc.NSObject

}

func (r *UIContentUnavailableButtonProperties) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableButtonProperties) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableButtonProperties) Role() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableButtonProperties) SetRole() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableButtonProperties) PrimaryAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableButtonProperties) SetPrimaryAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableButtonProperties) Menu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableButtonProperties) SetMenu() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

