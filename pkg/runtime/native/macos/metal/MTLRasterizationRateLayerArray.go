//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLRasterizationRateLayerArray : objc.NSObject
*/

type MTLRasterizationRateLayerArray struct {
  *objc.NSObject

}

func (r *MTLRasterizationRateLayerArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerArray) ObjectAtIndexedSubscript() (*MTLRasterizationRateLayerDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

