//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLVisibleFunctionTable
*/

type MTLVisibleFunctionTable struct {

}

func (r *MTLVisibleFunctionTable) SetFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLVisibleFunctionTable) SetFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLVisibleFunctionTable) GpuResourceID() (MTLResourceID, error) {
  return MTLResourceID{}, fmt.Errorf("unimplemented")
}

