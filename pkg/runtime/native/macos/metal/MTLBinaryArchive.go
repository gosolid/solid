//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLBinaryArchive
*/

type MTLBinaryArchive struct {

}

func (r *MTLBinaryArchive) AddRenderPipelineFunctionsWithDescriptor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLBinaryArchive) AddTileRenderPipelineFunctionsWithDescriptor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLBinaryArchive) SerializeToURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLBinaryArchive) AddFunctionWithDescriptor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLBinaryArchive) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBinaryArchive) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLBinaryArchive) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBinaryArchive) AddComputePipelineFunctionsWithDescriptor() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

