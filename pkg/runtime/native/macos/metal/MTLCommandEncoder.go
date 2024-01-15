//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLCommandEncoder
*/

type MTLCommandEncoder struct {

}

func (r *MTLCommandEncoder) PushDebugGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandEncoder) PopDebugGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandEncoder) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandEncoder) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandEncoder) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandEncoder) EndEncoding() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandEncoder) InsertDebugSignpost() error {
  return fmt.Errorf("unimplemented")
}

