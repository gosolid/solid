//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIPopoverController : objc.NSObject
*/

type UIPopoverController struct {
  *objc.NSObject

}

func (r *UIPopoverController) SetPopoverBackgroundViewClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) DismissPopoverAnimated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) ContentViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) SetPopoverLayoutMargins() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) IsPopoverVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) PopoverLayoutMargins() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) PopoverBackgroundViewClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) SetContentViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) SetPopoverContentSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) PopoverContentSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) PassthroughViews() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) SetPassthroughViews() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) InitWithContentViewController() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) PresentPopoverFromRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) PresentPopoverFromBarButtonItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPopoverController) PopoverArrowDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

