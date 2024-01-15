//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPreviewParameters : objc.NSObject
*/

type UIPreviewParameters struct {
  *objc.NSObject

}

func (r *UIPreviewParameters) InitWithTextLineRects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewParameters) VisiblePath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewParameters) SetVisiblePath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewParameters) ShadowPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewParameters) SetShadowPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewParameters) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewParameters) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPreviewParameters) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

