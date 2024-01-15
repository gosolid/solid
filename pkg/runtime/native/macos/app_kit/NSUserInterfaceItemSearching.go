//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSUserInterfaceItemSearching
*/

type NSUserInterfaceItemSearching struct {

}

func (r *NSUserInterfaceItemSearching) SearchForItemsWithSearchString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceItemSearching) LocalizedTitlesForItem() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceItemSearching) PerformActionForItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserInterfaceItemSearching) ShowAllHelpTopicsForSearchString() error {
  return fmt.Errorf("unimplemented")
}

