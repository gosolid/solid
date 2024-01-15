//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTextFinderBarContainer
*/

type NSTextFinderBarContainer struct {

}

func (r *NSTextFinderBarContainer) ContentView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderBarContainer) FindBarView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderBarContainer) SetFindBarView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinderBarContainer) IsFindBarVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTextFinderBarContainer) SetFindBarVisible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTextFinderBarContainer) FindBarViewDidChangeHeight() error {
  return fmt.Errorf("unimplemented")
}

