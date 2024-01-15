//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSPageController : AppKit.NSViewController
*/

type NSPageController struct {
  *NSViewController
  *NSAnimatablePropertyContainer
  *foundation.NSCoding
}

func (r *NSPageController) TakeSelectedIndexFrom() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageController) ArrangedObjects() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPageController) SetSelectedIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageController) SelectedIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPageController) NavigateBack() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageController) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageController) SetTransitionStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageController) SelectedViewController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPageController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPageController) TransitionStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPageController) SetArrangedObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageController) NavigateForwardToObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageController) CompleteTransition() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPageController) NavigateForward() error {
  return fmt.Errorf("unimplemented")
}

