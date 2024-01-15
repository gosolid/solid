//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDragPreview : objc.NSObject
*/

type UIDragPreview struct {
  *objc.NSObject

}

func (r *UIDragPreview) InitWithView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragPreview) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragPreview) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragPreview) View() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragPreview) Parameters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

