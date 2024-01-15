//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIMenuSystem : objc.NSObject
*/

type UIMenuSystem struct {
  *objc.NSObject

}

func (r *UIMenuSystem) MainSystem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuSystem) ContextSystem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuSystem) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuSystem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuSystem) SetNeedsRebuild() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMenuSystem) SetNeedsRevalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

