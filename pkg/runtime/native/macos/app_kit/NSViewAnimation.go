//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSViewAnimation : AppKit.NSAnimation
*/

type NSViewAnimation struct {
  *NSAnimation

}

func (r *NSViewAnimation) InitWithViewAnimations() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewAnimation) ViewAnimations() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSViewAnimation) SetViewAnimations() error {
  return fmt.Errorf("unimplemented")
}

