//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIFilter : objc.NSObject
*/

type CIFilter struct {
  *objc.NSObject

}

func (r *CIFilter) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) OutputImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) InputKeys() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) OutputKeys() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) Attributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) SetDefaults() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) Apply() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilter) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

