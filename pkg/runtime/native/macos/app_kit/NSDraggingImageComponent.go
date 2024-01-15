//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSDraggingImageComponent : objc.NSObject
*/

type NSDraggingImageComponent struct {
  *objc.NSObject

}

func (r *NSDraggingImageComponent) DraggingImageComponentWithKey() (*NSDraggingImageComponent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingImageComponent) SetFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingImageComponent) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSDraggingImageComponent) InitWithKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingImageComponent) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingImageComponent) Key() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingImageComponent) SetKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDraggingImageComponent) Contents() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDraggingImageComponent) SetContents() error {
  return fmt.Errorf("unimplemented")
}

