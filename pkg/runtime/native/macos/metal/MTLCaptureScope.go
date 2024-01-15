//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
protocol Metal.MTLCaptureScope
*/

type MTLCaptureScope struct {

}

func (r *MTLCaptureScope) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCaptureScope) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureScope) CommandQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureScope) BeginScope() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCaptureScope) EndScope() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCaptureScope) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

