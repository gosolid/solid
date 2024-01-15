//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSWorkspaceOpenConfiguration : objc.NSObject
*/

type NSWorkspaceOpenConfiguration struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSWorkspaceOpenConfiguration) PromptsUserIfNeeded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetPromptsUserIfNeeded() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetHides() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) CreatesNewApplicationInstance() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetRequiresUniversalLinks() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) Arguments() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetForPrinting() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetArchitecture() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) HidesOthers() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetHidesOthers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetArguments() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) RequiresUniversalLinks() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) Configuration() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) Hides() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) IsForPrinting() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) AllowsRunningApplicationSubstitution() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) Environment() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) Architecture() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) Activates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetActivates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetAllowsRunningApplicationSubstitution() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetEnvironment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) AddsToRecentItems() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetAddsToRecentItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetCreatesNewApplicationInstance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) AppleEvent() (*foundation.NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWorkspaceOpenConfiguration) SetAppleEvent() error {
  return fmt.Errorf("unimplemented")
}

