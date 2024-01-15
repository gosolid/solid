//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIBarItem : objc.NSObject
*/

type UIBarItem struct {
  *objc.NSObject

}

func (r *UIBarItem) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) LargeContentSizeImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) ImageInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetTag() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetLargeContentSizeImageInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) TitleTextAttributesForState() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) LandscapeImagePhone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetImageInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetLandscapeImagePhoneInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetLargeContentSizeImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) Tag() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) LandscapeImagePhoneInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) LargeContentSizeImageInsets() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetTitleTextAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarItem) SetLandscapeImagePhone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

