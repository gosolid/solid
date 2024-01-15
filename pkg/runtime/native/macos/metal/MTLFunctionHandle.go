//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLFunctionHandle
*/

type MTLFunctionHandle struct {

}

func (r *MTLFunctionHandle) FunctionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionHandle) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionHandle) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

