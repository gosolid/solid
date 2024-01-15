//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSScrubberImageItemView : AppKit.NSScrubberItemView
*/

type NSScrubberImageItemView struct {
  *NSScrubberItemView

}

func (r *NSScrubberImageItemView) SetImageAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberImageItemView) ImageView() (*NSImageView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberImageItemView) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScrubberImageItemView) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSScrubberImageItemView) ImageAlignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

