//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIAreaLogarithmicHistogram
*/

type CIAreaLogarithmicHistogram struct {

}

func (r *CIAreaLogarithmicHistogram) SetMaximumStop() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAreaLogarithmicHistogram) Scale() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAreaLogarithmicHistogram) SetScale() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAreaLogarithmicHistogram) Count() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAreaLogarithmicHistogram) SetCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAreaLogarithmicHistogram) MinimumStop() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAreaLogarithmicHistogram) SetMinimumStop() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAreaLogarithmicHistogram) MaximumStop() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

