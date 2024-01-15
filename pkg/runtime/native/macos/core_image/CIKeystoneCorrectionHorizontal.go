//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIKeystoneCorrectionHorizontal
*/

type CIKeystoneCorrectionHorizontal struct {

}

func (r *CIKeystoneCorrectionHorizontal) FocalLength() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIKeystoneCorrectionHorizontal) SetFocalLength() error {
  return fmt.Errorf("unimplemented")
}

