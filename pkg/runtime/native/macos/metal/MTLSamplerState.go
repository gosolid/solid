//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLSamplerState
*/

type MTLSamplerState struct {

}

func (r *MTLSamplerState) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerState) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSamplerState) GpuResourceID() (MTLResourceID, error) {
  return MTLResourceID{}, fmt.Errorf("unimplemented")
}

