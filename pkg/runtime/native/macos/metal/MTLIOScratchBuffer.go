//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLIOScratchBuffer
*/

type MTLIOScratchBuffer struct {

}

func (r *MTLIOScratchBuffer) Buffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

