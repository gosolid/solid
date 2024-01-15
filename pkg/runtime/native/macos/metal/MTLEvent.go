//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLEvent
*/

type MTLEvent struct {

}

func (r *MTLEvent) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLEvent) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLEvent) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

