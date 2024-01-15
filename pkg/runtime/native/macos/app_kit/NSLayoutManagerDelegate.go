//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol AppKit.NSLayoutManagerDelegate
*/

type NSLayoutManagerDelegate struct {

}

func (r *NSLayoutManagerDelegate) LayoutManager() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSLayoutManagerDelegate) LayoutManagerDidInvalidateLayout() error {
  return fmt.Errorf("unimplemented")
}

