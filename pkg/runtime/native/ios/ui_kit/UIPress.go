//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPress : objc.NSObject
*/

type UIPress struct {
  *objc.NSObject

}

func (r *UIPress) GestureRecognizers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPress) Force() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPress) Key() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPress) Timestamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPress) Phase() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPress) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPress) Window() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPress) Responder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

