//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRasterizationRateMapDescriptor : objc.NSObject
*/

type MTLRasterizationRateMapDescriptor struct {
  *objc.NSObject

}

func (r *MTLRasterizationRateMapDescriptor) SetScreenSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) RasterizationRateMapDescriptorWithScreenSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) LayerAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) ScreenSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) SetLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) LayerCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) SetLayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateMapDescriptor) Layers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

