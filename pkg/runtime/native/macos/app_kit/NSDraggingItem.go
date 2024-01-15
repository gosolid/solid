//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSDraggingItem : objc.NSObject
*/

type NSDraggingItem struct {
  *objc.NSObject

}

func (r *NSDraggingItem) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingItem) SetDraggingFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingItem) Item() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingItem) DraggingFrame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSDraggingItem) ImageComponentsProvider() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingItem) SetImageComponentsProvider() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingItem) ImageComponents() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingItem) InitWithPasteboardWriter() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

