//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIMenuElement : objc.NSObject
*/

type UIMenuElement struct {
  *objc.NSObject

}

func (r *UIMenuElement) Subtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuElement) SetSubtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuElement) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuElement) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuElement) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuElement) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuElement) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

