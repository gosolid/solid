//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIDeferredMenuElement : UIKit.UIMenuElement
*/

type UIDeferredMenuElement struct {
  *UIMenuElement

}

func (r *UIDeferredMenuElement) ElementWithUncachedProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDeferredMenuElement) ElementWithProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

