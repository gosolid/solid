//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLHeapDescriptor : objc.NSObject
*/

type MTLHeapDescriptor struct {
  *objc.NSObject

}

func (r *MTLHeapDescriptor) Size() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) CpuCacheMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetStorageMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SparsePageSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetSparsePageSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) ResourceOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetResourceOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) StorageMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetCpuCacheMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) HazardTrackingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLHeapDescriptor) SetHazardTrackingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

