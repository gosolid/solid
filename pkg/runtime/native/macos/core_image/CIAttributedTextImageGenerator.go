//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CoreImage.CIAttributedTextImageGenerator
*/

type CIAttributedTextImageGenerator struct {

}

func (r *CIAttributedTextImageGenerator) Text() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIAttributedTextImageGenerator) SetText() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAttributedTextImageGenerator) ScaleFactor() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAttributedTextImageGenerator) SetScaleFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAttributedTextImageGenerator) Padding() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAttributedTextImageGenerator) SetPadding() error {
  return fmt.Errorf("unimplemented")
}

