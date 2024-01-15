//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLHeap
*/

type MTLHeap struct {

}

func (r *MTLHeap) SetPurgeableState() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLHeap) CpuCacheMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) Size() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) NewAccelerationStructureWithSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) NewAccelerationStructureWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) NewBufferWithLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) NewTextureWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) StorageMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) HazardTrackingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) ResourceOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) MaxAvailableSizeWithAlignment() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) UsedSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeap) CurrentAllocatedSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

