//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIBarAppearance : objc.NSObject
*/

type UIBarAppearance struct {
  *objc.NSObject

}

func (r *UIBarAppearance) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) ConfigureWithDefaultBackground() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) SetBackgroundEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) InitWithIdiom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) Copy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) BackgroundEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) SetShadowColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) SetShadowImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) InitWithBarAppearance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) ConfigureWithOpaqueBackground() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) ConfigureWithTransparentBackground() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) SetBackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) SetBackgroundImageContentMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) Idiom() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) BackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) SetBackgroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) BackgroundImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) BackgroundImageContentMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) ShadowColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBarAppearance) ShadowImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

