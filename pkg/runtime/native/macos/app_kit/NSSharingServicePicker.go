//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSSharingServicePicker : objc.NSObject
*/

type NSSharingServicePicker struct {
  *objc.NSObject

}

func (r *NSSharingServicePicker) InitWithItems() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePicker) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePicker) ShowRelativeToRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePicker) Close() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePicker) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePicker) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSharingServicePicker) StandardShareMenuItem() (*NSMenuItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

