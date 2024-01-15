//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSBrowserCell : AppKit.NSCell
*/

type NSBrowserCell struct {
  *NSCell

}

func (r *NSBrowserCell) InitTextCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) HighlightColorInView() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) BranchImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) HighlightedBranchImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) AlternateImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) SetAlternateImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) InitImageCell() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) Set() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) IsLeaf() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) SetLeaf() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) IsLoaded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBrowserCell) SetLoaded() error {
  return fmt.Errorf("unimplemented")
}

