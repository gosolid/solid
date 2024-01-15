//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIHoverHighlightEffect : objc.NSObject
*/

type UIHoverHighlightEffect struct {
  *objc.NSObject

}

func (r *UIHoverHighlightEffect) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverHighlightEffect) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIHoverHighlightEffect) Effect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

