//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPreviewTarget : objc.NSObject
*/

type UIPreviewTarget struct {
  *objc.NSObject

}

func (r *UIPreviewTarget) Center() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewTarget) Transform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewTarget) InitWithContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewTarget) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewTarget) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewTarget) Container() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

