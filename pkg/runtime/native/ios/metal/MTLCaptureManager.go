//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLCaptureManager : objc.NSObject
*/

type MTLCaptureManager struct {
  *objc.NSObject

}

func (r *MTLCaptureManager) StartCaptureWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) StartCaptureWithDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) StopCapture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) StartCaptureWithScope() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) DefaultCaptureScope() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) SupportsDestination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) SetDefaultCaptureScope() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) IsCapturing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) SharedCaptureManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) NewCaptureScopeWithDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) NewCaptureScopeWithCommandQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureManager) StartCaptureWithCommandQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

