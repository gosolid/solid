//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSPrintPanelAccessorizing
*/

type NSPrintPanelAccessorizing struct {

}

func (r *NSPrintPanelAccessorizing) KeyPathsForValuesAffectingPreview() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPrintPanelAccessorizing) LocalizedSummaryItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

