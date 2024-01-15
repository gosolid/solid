//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLIOCommandQueue
*/

type MTLIOCommandQueue struct {

}

func (r *MTLIOCommandQueue) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueue) EnqueueBarrier() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueue) CommandBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueue) CommandBufferWithUnretainedReferences() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueue) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

