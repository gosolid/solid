//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSStatusBar : objc.NSObject
*/

type NSStatusBar struct {
  *objc.NSObject

}

func (r *NSStatusBar) SystemStatusBar() (*NSStatusBar, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStatusBar) IsVertical() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSStatusBar) Thickness() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStatusBar) StatusItemWithLength() (*NSStatusItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStatusBar) RemoveStatusItem() error {
  return fmt.Errorf("unimplemented")
}

