//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIOpenURLContext : objc.NSObject
*/

type UIOpenURLContext struct {
  *objc.NSObject

}

func (r *UIOpenURLContext) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIOpenURLContext) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIOpenURLContext) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIOpenURLContext) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

