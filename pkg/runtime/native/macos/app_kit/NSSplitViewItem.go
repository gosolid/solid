//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSSplitViewItem : objc.NSObject
*/

type NSSplitViewItem struct {
  *objc.NSObject
  *NSAnimatablePropertyContainer
  *foundation.NSCoding
}

func (r *NSSplitViewItem) SidebarWithViewController() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) MaximumThickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetHoldingPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetAllowsFullHeightLayout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) Behavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) CollapseBehavior() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) IsSpringLoaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) ViewController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) MinimumThickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetPreferredThicknessFraction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetSpringLoaded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetTitlebarSeparatorStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) ContentListWithViewController() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetMaximumThickness() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetMinimumThickness() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) AutomaticMaximumThickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SplitViewItemWithViewController() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) IsCollapsed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) AllowsFullHeightLayout() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) InspectorWithViewController() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) PreferredThicknessFraction() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetCollapseBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetAutomaticMaximumThickness() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) CanCollapseFromWindowResize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) TitlebarSeparatorStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetCollapsed() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetCanCollapse() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) HoldingPriority() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetCanCollapseFromWindowResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) SetViewController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSplitViewItem) CanCollapse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

