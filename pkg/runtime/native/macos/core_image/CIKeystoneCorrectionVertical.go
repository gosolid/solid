//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIKeystoneCorrectionVertical
*/

type CIKeystoneCorrectionVertical struct {

}

func (r *CIKeystoneCorrectionVertical) FocalLength() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIKeystoneCorrectionVertical) SetFocalLength() error {
  return fmt.Errorf("unimplemented")
}

