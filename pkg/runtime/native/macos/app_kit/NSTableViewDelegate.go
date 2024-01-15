//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTableViewDelegate
*/

type NSTableViewDelegate struct {

}

func (r *NSTableViewDelegate) SelectionShouldChangeInTableView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDelegate) TableViewSelectionDidChange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewDelegate) TableViewColumnDidMove() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewDelegate) TableViewColumnDidResize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewDelegate) TableViewSelectionIsChanging() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewDelegate) TableView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

