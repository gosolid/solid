//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UICommandAlternate : objc.NSObject
*/

type UICommandAlternate struct {
  *objc.NSObject

}

func (r *UICommandAlternate) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommandAlternate) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommandAlternate) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommandAlternate) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommandAlternate) Action() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommandAlternate) ModifierFlags() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICommandAlternate) AlternateWithTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

