//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSOutlineViewDelegate
*/

type NSOutlineViewDelegate struct {

}

func (r *NSOutlineViewDelegate) OutlineView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOutlineViewDelegate) OutlineViewSelectionDidChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineViewDelegate) OutlineViewColumnDidResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineViewDelegate) OutlineViewSelectionIsChanging() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineViewDelegate) OutlineViewItemDidCollapse() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineViewDelegate) SelectionShouldChangeInOutlineView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOutlineViewDelegate) OutlineViewColumnDidMove() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineViewDelegate) OutlineViewItemWillExpand() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineViewDelegate) OutlineViewItemDidExpand() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOutlineViewDelegate) OutlineViewItemWillCollapse() error {
  return fmt.Errorf("unimplemented")
}

