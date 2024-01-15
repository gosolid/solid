//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLFunction
*/

type MTLFunction struct {

}

func (r *MTLFunction) StageInputAttributes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) FunctionConstantsDictionary() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLFunction) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) FunctionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) PatchType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) PatchControlPointCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) NewArgumentEncoderWithBufferIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) VertexAttributes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLFunction) Options() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

