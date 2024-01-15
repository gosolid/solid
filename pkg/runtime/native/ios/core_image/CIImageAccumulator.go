//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIImageAccumulator : objc.NSObject
*/

type CIImageAccumulator struct {
  *objc.NSObject

}

func (r *CIImageAccumulator) Clear() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) Extent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) Format() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) ImageAccumulatorWithExtent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) InitWithExtent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) Image() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) SetImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

