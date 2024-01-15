//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIContentUnavailableConfigurationState : objc.NSObject
*/

type UIContentUnavailableConfigurationState struct {
  *objc.NSObject

}

func (r *UIContentUnavailableConfigurationState) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableConfigurationState) TraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableConfigurationState) SetTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableConfigurationState) SearchText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableConfigurationState) SetSearchText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableConfigurationState) InitWithTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableConfigurationState) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableConfigurationState) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

