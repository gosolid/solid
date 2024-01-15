//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIContextMenuConfiguration : objc.NSObject
*/

type UIContextMenuConfiguration struct {
  *objc.NSObject

}

func (r *UIContextMenuConfiguration) SetBadgeCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuConfiguration) PreferredMenuElementOrder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuConfiguration) SetPreferredMenuElementOrder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuConfiguration) ConfigurationWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuConfiguration) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuConfiguration) SecondaryItemIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuConfiguration) SetSecondaryItemIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContextMenuConfiguration) BadgeCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

