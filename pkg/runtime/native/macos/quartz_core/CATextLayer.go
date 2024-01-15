//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface QuartzCore.CATextLayer : QuartzCore.CALayer
*/

type CATextLayer struct {
  *CALayer

}

func (r *CATextLayer) ForegroundColor() (*core_graphics.CGColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetTruncationMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATextLayer) AlignmentMode() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetAllowsFontSubpixelQuantization() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetFontSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATextLayer) Font() (*void, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) FontSize() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetWrapped() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATextLayer) TruncationMode() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) AllowsFontSubpixelQuantization() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) String() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetString() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetFont() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetForegroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATextLayer) IsWrapped() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetAlignmentMode() error {
  return fmt.Errorf("unimplemented")
}

