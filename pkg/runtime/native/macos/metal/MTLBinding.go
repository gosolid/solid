//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLBinding
*/

type MTLBinding struct {

}

func (r *MTLBinding) Index() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLBinding) IsUsed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLBinding) IsArgument() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLBinding) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBinding) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLBinding) Access() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

