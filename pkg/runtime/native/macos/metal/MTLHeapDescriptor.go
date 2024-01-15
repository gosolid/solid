//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLHeapDescriptor : objc.NSObject
*/

type MTLHeapDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLHeapDescriptor) SetSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetSparsePageSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) ResourceOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) Size() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetCpuCacheMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetStorageMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) CpuCacheMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetHazardTrackingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) StorageMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SparsePageSize() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) HazardTrackingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetResourceOptions() error {
  return fmt.Errorf("unimplemented")
}

