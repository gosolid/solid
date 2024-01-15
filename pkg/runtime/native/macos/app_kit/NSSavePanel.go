//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSSavePanel : AppKit.NSPanel
*/

type NSSavePanel struct {
  *NSPanel

}

func (r *NSSavePanel) ShowsTagField() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) URL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) IsExpanded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetCanSelectHiddenExtension() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) Message() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetTagNames() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) Ok() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) Cancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetAccessoryView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) CanSelectHiddenExtension() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) IsExtensionHidden() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) BeginWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetCanCreateDirectories() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) ShowsHiddenFiles() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SavePanel() (*NSSavePanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) ValidateVisibleColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) AllowedContentTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetTitle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) NameFieldStringValue() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetMessage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) BeginSheetModalForWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) RunModal() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetDirectoryURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) CanCreateDirectories() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetPrompt() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetNameFieldLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetNameFieldStringValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetShowsTagField() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) Identifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) AllowsOtherFileTypes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetExtensionHidden() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) NameFieldLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) TagNames() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) Title() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetShowsHiddenFiles() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetTreatsFilePackagesAsDirectories() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) Prompt() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) DirectoryURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetAllowedContentTypes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) SetAllowsOtherFileTypes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) AccessoryView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSavePanel) TreatsFilePackagesAsDirectories() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

