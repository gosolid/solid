//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIKeyCommand : UIKit.UICommand
*/

type UIKeyCommand struct {
  *UICommand

}

func (r *UIKeyCommand) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) PropertyList() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) State() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) Action() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) WantsPriorityOverSystemBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) CommandWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) KeyCommandWithInput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) ModifierFlags() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) SetAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) SetAllowsAutomaticMirroring() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) DiscoverabilityTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) SetDiscoverabilityTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) SetState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) SetWantsPriorityOverSystemBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) SetAllowsAutomaticLocalization() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) AllowsAutomaticMirroring() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) Attributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) Alternates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) AllowsAutomaticLocalization() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIKeyCommand) Input() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

