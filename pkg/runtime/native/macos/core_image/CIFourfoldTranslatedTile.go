//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIFourfoldTranslatedTile
*/

type CIFourfoldTranslatedTile struct {

}

func (r *CIFourfoldTranslatedTile) SetAcuteAngle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldTranslatedTile) SetCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldTranslatedTile) Angle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldTranslatedTile) SetWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldTranslatedTile) AcuteAngle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldTranslatedTile) Width() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldTranslatedTile) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldTranslatedTile) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFourfoldTranslatedTile) Center() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIFourfoldTranslatedTile) SetAngle() error {
  return fmt.Errorf("unimplemented")
}

