//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTouchBarItem : objc.NSObject
*/

type NSTouchBarItem struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSTouchBarItem) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBarItem) ViewController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBarItem) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBarItem) Identifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBarItem) VisibilityPriority() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTouchBarItem) SetVisibilityPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTouchBarItem) InitWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBarItem) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBarItem) CustomizationLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouchBarItem) IsVisible() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

