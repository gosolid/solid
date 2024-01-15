//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CoreImage.CITextImageGenerator
*/

type CITextImageGenerator struct {

}

func (r *CITextImageGenerator) Text() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITextImageGenerator) FontSize() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITextImageGenerator) ScaleFactor() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITextImageGenerator) SetScaleFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITextImageGenerator) Padding() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITextImageGenerator) SetText() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITextImageGenerator) FontName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITextImageGenerator) SetFontName() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITextImageGenerator) SetFontSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITextImageGenerator) SetPadding() error {
  return fmt.Errorf("unimplemented")
}

