//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSViewController : AppKit.NSResponder
*/

type NSViewController struct {
  *NSResponder
  *NSEditor
  *NSSeguePerforming
  *NSUserInterfaceItemIdentification
}

func (r *NSViewController) CommitEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSViewController) CommitEditingWithDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) ViewDidLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) RepresentedObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewController) PreferredContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSViewController) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewController) LoadView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) ViewDidLoad() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) ViewDidDisappear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) NibName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewController) InitWithNibName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewController) UpdateViewConstraints() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) ViewIfLoaded() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewController) IsViewLoaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSViewController) ViewDidAppear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) ViewWillDisappear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) ViewWillLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewController) NibBundle() (*foundation.NSBundle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewController) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewController) SetPreferredContentSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) LoadViewIfNeeded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) DiscardEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) ViewWillAppear() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) SetRepresentedObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSViewController) SetView() error {
  return fmt.Errorf("unimplemented")
}

