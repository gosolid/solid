//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CoreImage.CIPDF417BarcodeGenerator
*/

type CIPDF417BarcodeGenerator struct {

}

func (r *CIPDF417BarcodeGenerator) MinHeight() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) Rows() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetRows() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) PreferredAspectRatio() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) AlwaysSpecifyCompaction() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) MaxWidth() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetPreferredAspectRatio() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetCompactionMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetAlwaysSpecifyCompaction() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) MinWidth() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetMinHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) MaxHeight() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetMaxHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetCompactStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetCorrectionLevel() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) CompactionMode() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) CompactStyle() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) Message() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetMessage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetMinWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetMaxWidth() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) DataColumns() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) SetDataColumns() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPDF417BarcodeGenerator) CorrectionLevel() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

