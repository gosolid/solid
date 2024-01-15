//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDynamicItemGroup : objc.NSObject
*/

type UIDynamicItemGroup struct {
  *objc.NSObject

}

func (r *UIDynamicItemGroup) InitWithItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDynamicItemGroup) Items() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

