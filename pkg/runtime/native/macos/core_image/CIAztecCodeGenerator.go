//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CoreImage.CIAztecCodeGenerator
*/

type CIAztecCodeGenerator struct {

}

func (r *CIAztecCodeGenerator) Message() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIAztecCodeGenerator) SetMessage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAztecCodeGenerator) CorrectionLevel() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAztecCodeGenerator) SetCorrectionLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAztecCodeGenerator) Layers() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAztecCodeGenerator) SetLayers() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAztecCodeGenerator) CompactStyle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAztecCodeGenerator) SetCompactStyle() error {
  return fmt.Errorf("unimplemented")
}

