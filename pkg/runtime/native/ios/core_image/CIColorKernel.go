//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CIColorKernel : CoreImage.CIKernel
*/

type CIColorKernel struct {
  *CIKernel

}

func (r *CIColorKernel) KernelWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorKernel) ApplyWithExtent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

