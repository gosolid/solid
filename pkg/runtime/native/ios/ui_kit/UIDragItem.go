//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDragItem : objc.NSObject
*/

type UIDragItem struct {
  *objc.NSObject

}

func (r *UIDragItem) InitWithItemProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragItem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragItem) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragItem) ItemProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragItem) LocalObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragItem) SetLocalObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragItem) PreviewProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragItem) SetPreviewProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

