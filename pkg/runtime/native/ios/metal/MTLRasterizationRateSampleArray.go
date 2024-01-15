//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLRasterizationRateSampleArray : objc.NSObject
*/

type MTLRasterizationRateSampleArray struct {
  *objc.NSObject

}

func (r *MTLRasterizationRateSampleArray) ObjectAtIndexedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateSampleArray) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

