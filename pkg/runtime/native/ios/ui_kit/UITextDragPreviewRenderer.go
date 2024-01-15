//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextDragPreviewRenderer : objc.NSObject
*/

type UITextDragPreviewRenderer struct {
  *objc.NSObject

}

func (r *UITextDragPreviewRenderer) BodyRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDragPreviewRenderer) LastLineRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDragPreviewRenderer) InitWithLayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDragPreviewRenderer) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDragPreviewRenderer) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDragPreviewRenderer) LayoutManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDragPreviewRenderer) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDragPreviewRenderer) FirstLineRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextDragPreviewRenderer) AdjustFirstLineRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

