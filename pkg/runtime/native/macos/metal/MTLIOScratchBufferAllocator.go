//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLIOScratchBufferAllocator
*/

type MTLIOScratchBufferAllocator struct {

}

func (r *MTLIOScratchBufferAllocator) NewScratchBufferWithMinimumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

