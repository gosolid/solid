//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSUserDefaultsController : AppKit.NSController
*/

type NSUserDefaultsController struct {
  *NSController

}

func (r *NSUserDefaultsController) InitWithDefaults() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) InitialValues() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) SetAppliesImmediately() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) Values() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) Defaults() (*foundation.NSUserDefaults, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) HasUnappliedChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) Revert() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) Save() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) SharedUserDefaultsController() (*NSUserDefaultsController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) SetInitialValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) AppliesImmediately() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaultsController) RevertToInitialValues() error {
  return fmt.Errorf("unimplemented")
}

