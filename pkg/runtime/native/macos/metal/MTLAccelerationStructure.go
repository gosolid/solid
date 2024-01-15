//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLAccelerationStructure
*/

type MTLAccelerationStructure struct {

}

func (r *MTLAccelerationStructure) Size() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLAccelerationStructure) GpuResourceID() (MTLResourceID, error) {
  return MTLResourceID{}, fmt.Errorf("unimplemented")
}

