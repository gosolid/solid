//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UICommand : UIKit.UIMenuElement
*/

type UICommand struct {
  *UIMenuElement

}

func (r *UICommand) SetAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) State() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) SetState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) DiscoverabilityTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) Attributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) SetDiscoverabilityTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) PropertyList() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) CommandWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) Action() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommand) Alternates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

