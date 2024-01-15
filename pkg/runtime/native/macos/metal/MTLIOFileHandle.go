//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLIOFileHandle
*/

type MTLIOFileHandle struct {

}

func (r *MTLIOFileHandle) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOFileHandle) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

