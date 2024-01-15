//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSValidatedUserInterfaceItem
*/

type NSValidatedUserInterfaceItem struct {

}

func (r *NSValidatedUserInterfaceItem) Action() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSValidatedUserInterfaceItem) Tag() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

