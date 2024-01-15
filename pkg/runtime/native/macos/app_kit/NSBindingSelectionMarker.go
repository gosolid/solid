//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSBindingSelectionMarker : objc.NSObject
*/

type NSBindingSelectionMarker struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSBindingSelectionMarker) NotApplicableSelectionMarker() (*NSBindingSelectionMarker, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBindingSelectionMarker) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBindingSelectionMarker) SetDefaultPlaceholder() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBindingSelectionMarker) DefaultPlaceholderForMarker() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBindingSelectionMarker) MultipleValuesSelectionMarker() (*NSBindingSelectionMarker, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBindingSelectionMarker) NoSelectionMarker() (*NSBindingSelectionMarker, error) {
  return nil, fmt.Errorf("unimplemented")
}

