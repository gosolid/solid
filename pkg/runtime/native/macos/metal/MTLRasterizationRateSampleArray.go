//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLRasterizationRateSampleArray : objc.NSObject
*/

type MTLRasterizationRateSampleArray struct {
  *objc.NSObject

}

func (r *MTLRasterizationRateSampleArray) ObjectAtIndexedSubscript() (*foundation.NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLRasterizationRateSampleArray) SetObject() error {
  return fmt.Errorf("unimplemented")
}

