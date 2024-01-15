//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSWindowController : AppKit.NSResponder
*/

type NSWindowController struct {
  *NSResponder
  *NSSeguePerforming
}

func (r *NSWindowController) SetWindowFrameAutosaveName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) SetWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) InitWithWindowNibPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) LoadWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) ShouldCloseDocument() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) SynchronizeWindowTitleWithDocumentName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) Owner() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) WindowNibName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) WindowFrameAutosaveName() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) PreviewRepresentableActivityItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) SetPreviewRepresentableActivityItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) WindowWillLoad() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) ShowWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) Close() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) ContentViewController() (*NSViewController, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) Window() (*NSWindow, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) InitWithWindow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) WindowDidLoad() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) WindowNibPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) SetShouldCloseDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) IsWindowLoaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) InitWithWindowNibName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) WindowTitleForDocumentDisplayName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) SetContentViewController() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) SetDocumentEdited() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) ShouldCascadeWindows() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSWindowController) SetDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) SetShouldCascadeWindows() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWindowController) Document() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

