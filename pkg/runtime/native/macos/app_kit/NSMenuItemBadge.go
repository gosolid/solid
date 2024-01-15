//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSMenuItemBadge : objc.NSObject
*/

type NSMenuItemBadge struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSMenuItemBadge) StringValue() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemBadge) NewItemsWithCount() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemBadge) InitWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemBadge) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemBadge) ItemCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemBadge) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemBadge) UpdatesWithCount() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemBadge) AlertsWithCount() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMenuItemBadge) InitWithCount() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

