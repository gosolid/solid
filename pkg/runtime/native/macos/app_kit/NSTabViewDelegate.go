//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSTabViewDelegate
*/

type NSTabViewDelegate struct {

}

func (r *NSTabViewDelegate) TabViewDidChangeNumberOfTabViewItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTabViewDelegate) TabView() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

