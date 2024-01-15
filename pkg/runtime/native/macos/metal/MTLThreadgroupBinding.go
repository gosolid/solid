//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLThreadgroupBinding
*/

type MTLThreadgroupBinding struct {

}

func (r *MTLThreadgroupBinding) ThreadgroupMemoryAlignment() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLThreadgroupBinding) ThreadgroupMemoryDataSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

