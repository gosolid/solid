//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Metal.MTLRasterizationRateLayerDescriptor : objc.NSObject
*/

type MTLRasterizationRateLayerDescriptor struct {
  *objc.NSObject

}

func (r *MTLRasterizationRateLayerDescriptor) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) InitWithSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) SampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) MaxSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) HorizontalSampleStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) VerticalSampleStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) Horizontal() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) Vertical() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

