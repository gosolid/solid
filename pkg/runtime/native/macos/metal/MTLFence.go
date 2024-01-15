//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLFence
*/

type MTLFence struct {

}

func (r *MTLFence) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFence) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFence) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

