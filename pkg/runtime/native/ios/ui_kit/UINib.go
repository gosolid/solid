//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UINib : objc.NSObject
*/

type UINib struct {
  *objc.NSObject

}

func (r *UINib) InstantiateWithOwner() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINib) NibWithNibName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UINib) NibWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

