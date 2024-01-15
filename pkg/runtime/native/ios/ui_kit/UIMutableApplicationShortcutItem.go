//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIMutableApplicationShortcutItem : UIKit.UIApplicationShortcutItem
*/

type UIMutableApplicationShortcutItem struct {
  *UIApplicationShortcutItem

}

func (r *UIMutableApplicationShortcutItem) SetIcon() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) SetUserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) LocalizedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) SetLocalizedSubtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) Icon() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) TargetContentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) SetTargetContentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) SetType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) SetLocalizedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableApplicationShortcutItem) LocalizedSubtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

