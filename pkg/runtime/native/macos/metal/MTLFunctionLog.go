//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLFunctionLog
*/

type MTLFunctionLog struct {

}

func (r *MTLFunctionLog) DebugLocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionLog) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionLog) EncoderLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunctionLog) Function() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

