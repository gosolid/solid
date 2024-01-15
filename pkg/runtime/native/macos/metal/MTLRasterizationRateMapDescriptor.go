//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLRasterizationRateMapDescriptor : objc.NSObject
*/

type MTLRasterizationRateMapDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLRasterizationRateMapDescriptor) RasterizationRateMapDescriptorWithScreenSize() (*MTLRasterizationRateMapDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) LayerAtIndex() (*MTLRasterizationRateLayerDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) SetLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) SetScreenSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) Layers() (*MTLRasterizationRateLayerArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) ScreenSize() (MTLSize, error) {
  return MTLSize{}, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) LayerCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

