//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLDepthStencilState
*/

type MTLDepthStencilState struct {

}

func (r *MTLDepthStencilState) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLDepthStencilState) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

