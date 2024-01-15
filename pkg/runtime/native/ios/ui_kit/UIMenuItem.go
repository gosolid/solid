//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIMenuItem : objc.NSObject
*/

type UIMenuItem struct {
  *objc.NSObject

}

func (r *UIMenuItem) InitWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuItem) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuItem) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuItem) Action() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuItem) SetAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

