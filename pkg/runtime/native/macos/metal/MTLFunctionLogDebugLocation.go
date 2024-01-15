//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLFunctionLogDebugLocation
*/

type MTLFunctionLogDebugLocation struct {

}

func (r *MTLFunctionLogDebugLocation) FunctionName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionLogDebugLocation) URL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionLogDebugLocation) Line() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionLogDebugLocation) Column() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

