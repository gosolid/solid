//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSVisualEffectView : AppKit.NSView
*/

type NSVisualEffectView struct {
  *NSView

}

func (r *NSVisualEffectView) SetState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) MaskImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) Material() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) InteriorBackgroundStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) SetBlendingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) ViewWillMoveToWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) State() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) SetMaskImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) IsEmphasized() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) SetEmphasized() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) ViewDidMoveToWindow() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) SetMaterial() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSVisualEffectView) BlendingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

