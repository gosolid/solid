//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSObjectController : AppKit.NSController
*/

type NSObjectController struct {
  *NSController

}

func (r *NSObjectController) Selection() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) SetAutomaticallyPreparesContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObjectController) CanAdd() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) CanRemove() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) NewObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) Remove() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObjectController) AddObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObjectController) Content() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) SelectedObjects() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) AutomaticallyPreparesContent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) ObjectClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) SetEditable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObjectController) PrepareContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObjectController) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) RemoveObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObjectController) Add() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObjectController) ValidateUserInterfaceItem() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSObjectController) SetContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObjectController) SetObjectClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSObjectController) InitWithContent() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

