//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIFilterConstructor
*/

type CIFilterConstructor struct {

}

func (r *CIFilterConstructor) FilterWithName() (*CIFilter, error) {
  return nil, fmt.Errorf("unimplemented")
}

