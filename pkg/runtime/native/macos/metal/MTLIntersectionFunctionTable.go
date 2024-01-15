//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLIntersectionFunctionTable
*/

type MTLIntersectionFunctionTable struct {

}

func (r *MTLIntersectionFunctionTable) SetFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTable) SetFunctions() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTable) SetOpaqueTriangleIntersectionFunctionWithSignature() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTable) SetOpaqueCurveIntersectionFunctionWithSignature() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTable) SetVisibleFunctionTable() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTable) SetVisibleFunctionTables() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTable) SetBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTable) SetBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIntersectionFunctionTable) GpuResourceID() (MTLResourceID, error) {
  return MTLResourceID{}, fmt.Errorf("unimplemented")
}

