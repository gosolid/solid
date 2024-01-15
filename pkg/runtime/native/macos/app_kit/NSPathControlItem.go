//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPathControlItem : objc.NSObject
*/

type NSPathControlItem struct {
  *objc.NSObject

}

func (r *NSPathControlItem) AttributedTitle() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControlItem) SetAttributedTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControlItem) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControlItem) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControlItem) URL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControlItem) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControlItem) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

