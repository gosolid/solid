//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "fmt"
)

/*
protocol CoreImage.CIPerspectiveTransformWithExtent
*/

type CIPerspectiveTransformWithExtent struct {

}

func (r *CIPerspectiveTransformWithExtent) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIPerspectiveTransformWithExtent) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

