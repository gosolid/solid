//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLBuffer
*/

type MTLBuffer struct {

}

func (r *MTLBuffer) NewTextureWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBuffer) AddDebugMarker() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBuffer) RemoveAllDebugMarkers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBuffer) RemoteStorageBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBuffer) Contents() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBuffer) DidModifyRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBuffer) NewRemoteBufferViewForDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBuffer) Length() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLBuffer) GpuAddress() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

