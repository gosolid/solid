//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRasterizationRateLayerArray : objc.NSObject
*/

type MTLRasterizationRateLayerArray struct {
  *objc.NSObject

}

func (r *MTLRasterizationRateLayerArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateLayerArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}
