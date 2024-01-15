//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSController : objc.NSObject
*/

type NSController struct {
  *objc.NSObject
  *foundation.NSCoding
  *NSEditor
  *NSEditorRegistration
}

func (r *NSController) ObjectDidBeginEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSController) ObjectDidEndEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSController) DiscardEditing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSController) CommitEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSController) CommitEditingWithDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSController) IsEditing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSController) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSController) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

