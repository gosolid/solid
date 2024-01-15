//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIKeystoneCorrectionCombined
*/

type CIKeystoneCorrectionCombined struct {

}

func (r *CIKeystoneCorrectionCombined) SetFocalLength() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKeystoneCorrectionCombined) FocalLength() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

