//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSHelpManager : objc.NSObject
*/

type NSHelpManager struct {
  *objc.NSObject

}

func (r *NSHelpManager) FindString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHelpManager) RegisterBooksInBundle() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHelpManager) SharedHelpManager() (*NSHelpManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHelpManager) SetContextHelpModeActive() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHelpManager) RemoveContextHelpForObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHelpManager) ContextHelpForObject() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHelpManager) ShowContextHelpForObject() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHelpManager) OpenHelpAnchor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHelpManager) SetContextHelp() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHelpManager) IsContextHelpModeActive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

