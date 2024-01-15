//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSplitViewController : AppKit.NSViewController
*/

type NSSplitViewController struct {
  *NSViewController
  *NSSplitViewDelegate
  *NSUserInterfaceValidations
}

func (r *NSSplitViewController) AddSplitViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) ValidateUserInterfaceItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) SetSplitViewItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) SetMinimumThicknessForInlineSidebars() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) SplitViewItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) MinimumThicknessForInlineSidebars() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) InsertSplitViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) RemoveSplitViewItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) SplitViewItemForViewController() (*NSSplitViewItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) ViewDidLoad() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) SplitView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewController) SetSplitView() error {
  return fmt.Errorf("unimplemented")
}

