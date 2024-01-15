//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol AppKit.NSAccessibilityTable
*/

type NSAccessibilityTable struct {

}

func (r *NSAccessibilityTable) AccessibilitySelectedRows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilityLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilityColumns() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilitySelectedColumns() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilitySelectedCells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilityVisibleCells() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilityVisibleColumns() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilityHeaderGroup() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilityColumnHeaderUIElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilityRows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) SetAccessibilitySelectedRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilityVisibleRows() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityTable) AccessibilityRowHeaderUIElements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

