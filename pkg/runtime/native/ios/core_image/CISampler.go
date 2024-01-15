//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CISampler : objc.NSObject
*/

type CISampler struct {
  *objc.NSObject

}

func (r *CISampler) SamplerWithImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISampler) InitWithImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISampler) Definition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISampler) Extent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

