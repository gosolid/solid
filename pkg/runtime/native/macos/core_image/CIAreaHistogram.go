//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIAreaHistogram
*/

type CIAreaHistogram struct {

}

func (r *CIAreaHistogram) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAreaHistogram) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAreaHistogram) Count() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAreaHistogram) SetCount() error {
  return fmt.Errorf("unimplemented")
}

