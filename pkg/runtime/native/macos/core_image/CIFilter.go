//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CoreImage.CIFilter
*/

type CIFilter struct {

}

func (r *CIFilter) CustomAttributes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) OutputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

