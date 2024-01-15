//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIAccordionFoldTransition
*/

type CIAccordionFoldTransition struct {

}

func (r *CIAccordionFoldTransition) BottomHeight() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAccordionFoldTransition) SetBottomHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAccordionFoldTransition) NumberOfFolds() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAccordionFoldTransition) SetNumberOfFolds() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAccordionFoldTransition) FoldShadowAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIAccordionFoldTransition) SetFoldShadowAmount() error {
  return fmt.Errorf("unimplemented")
}

