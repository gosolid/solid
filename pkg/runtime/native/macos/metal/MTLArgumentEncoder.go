//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLArgumentEncoder
*/

type MTLArgumentEncoder struct {

}

func (r *MTLArgumentEncoder) SetVisibleFunctionTable() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetComputePipelineStates() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetIndirectCommandBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetAccelerationStructure() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) NewArgumentEncoderForBufferAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetTextures() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetIndirectCommandBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetArgumentBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetSamplerStates() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetRenderPipelineState() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetIntersectionFunctionTable() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetSamplerState() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetComputePipelineState() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetIntersectionFunctionTables() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) ConstantDataAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetVisibleFunctionTables() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) EncodedLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) Alignment() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetRenderPipelineStates() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetBuffers() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLArgumentEncoder) SetBuffer() error {
  return fmt.Errorf("unimplemented")
}

