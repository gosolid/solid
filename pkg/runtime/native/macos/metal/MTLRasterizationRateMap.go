//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLRasterizationRateMap
*/

type MTLRasterizationRateMap struct {

}

func (r *MTLRasterizationRateMap) MapScreenToPhysicalCoordinates() (MTLSamplePosition, error) {
  return MTLSamplePosition{}, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMap) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMap) ScreenSize() (MTLSize, error) {
  return MTLSize{}, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMap) PhysicalGranularity() (MTLSize, error) {
  return MTLSize{}, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMap) ParameterBufferSizeAndAlign() (MTLSizeAndAlign, error) {
  return MTLSizeAndAlign{}, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMap) CopyParameterDataToBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMap) PhysicalSizeForLayer() (MTLSize, error) {
  return MTLSize{}, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMap) MapPhysicalToScreenCoordinates() (MTLSamplePosition, error) {
  return MTLSamplePosition{}, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMap) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMap) LayerCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

