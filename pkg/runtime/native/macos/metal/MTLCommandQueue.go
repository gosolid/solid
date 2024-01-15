//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLCommandQueue
*/

type MTLCommandQueue struct {

}

func (r *MTLCommandQueue) InsertDebugCaptureBoundary() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandQueue) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandQueue) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandQueue) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandQueue) CommandBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandQueue) CommandBufferWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandQueue) CommandBufferWithUnretainedReferences() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

