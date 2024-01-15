//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIHighlightShadowAdjust
*/

type CIHighlightShadowAdjust struct {

}

func (r *CIHighlightShadowAdjust) SetShadowAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHighlightShadowAdjust) HighlightAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHighlightShadowAdjust) SetHighlightAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHighlightShadowAdjust) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIHighlightShadowAdjust) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHighlightShadowAdjust) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIHighlightShadowAdjust) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIHighlightShadowAdjust) ShadowAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

