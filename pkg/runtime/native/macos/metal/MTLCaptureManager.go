//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLCaptureManager : objc.NSObject
*/

type MTLCaptureManager struct {
  *objc.NSObject

}

func (r *MTLCaptureManager) StartCaptureWithDevice() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) StopCapture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) IsCapturing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) SharedCaptureManager() (*MTLCaptureManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) NewCaptureScopeWithDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) NewCaptureScopeWithCommandQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) SupportsDestination() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) StartCaptureWithCommandQueue() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) StartCaptureWithScope() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) StartCaptureWithDescriptor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) DefaultCaptureScope() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) SetDefaultCaptureScope() error {
  return fmt.Errorf("unimplemented")
}

