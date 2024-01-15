//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIAction : UIKit.UIMenuElement
*/

type UIAction struct {
  *UIMenuElement

}

func (r *UIAction) Sender() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) ActionWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) DiscoverabilityTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) State() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) SetDiscoverabilityTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) SetState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) Attributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) SetAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) ActionWithHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAction) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

