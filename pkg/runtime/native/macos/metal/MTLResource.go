//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLResource
*/

type MTLResource struct {

}

func (r *MTLResource) SetPurgeableState() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLResource) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResource) StorageMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLResource) HeapOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLResource) MakeAliasable() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLResource) IsAliasable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLResource) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLResource) AllocatedSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLResource) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLResource) CpuCacheMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLResource) HazardTrackingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLResource) ResourceOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLResource) Heap() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

