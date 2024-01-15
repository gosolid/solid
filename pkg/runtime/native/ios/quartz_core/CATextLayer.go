//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CATextLayer : QuartzCore.CALayer
*/

type CATextLayer struct {
  *CALayer

}

func (r *CATextLayer) String() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) FontSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) Font() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) IsWrapped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetWrapped() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) AlignmentMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetAlignmentMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) AllowsFontSubpixelQuantization() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetAllowsFontSubpixelQuantization() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetFont() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetForegroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) TruncationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetFontSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) ForegroundColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATextLayer) SetTruncationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

