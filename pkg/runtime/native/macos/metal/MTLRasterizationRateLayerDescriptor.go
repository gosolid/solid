//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLRasterizationRateLayerDescriptor : objc.NSObject
*/

type MTLRasterizationRateLayerDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLRasterizationRateLayerDescriptor) HorizontalSampleStorage() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) VerticalSampleStorage() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) Horizontal() (*MTLRasterizationRateSampleArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) Vertical() (*MTLRasterizationRateSampleArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) InitWithSampleCount() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) SampleCount() (MTLSize, error) {
  return MTLSize{}, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerDescriptor) MaxSampleCount() (MTLSize, error) {
  return MTLSize{}, fmt.Errorf("unimplemented")
}

