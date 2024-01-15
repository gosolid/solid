//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSOpenPanel : AppKit.NSSavePanel
*/

type NSOpenPanel struct {
  *NSSavePanel

}

func (r *NSOpenPanel) SetAllowsMultipleSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) CanResolveUbiquitousConflicts() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) AllowsMultipleSelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) ResolvesAliases() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) SetResolvesAliases() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) CanChooseFiles() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) SetCanResolveUbiquitousConflicts() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) SetCanDownloadUbiquitousContents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) IsAccessoryViewDisclosed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) OpenPanel() (*NSOpenPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) SetCanChooseFiles() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) SetAccessoryViewDisclosed() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) URLs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) SetCanChooseDirectories() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) CanDownloadUbiquitousContents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOpenPanel) CanChooseDirectories() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

