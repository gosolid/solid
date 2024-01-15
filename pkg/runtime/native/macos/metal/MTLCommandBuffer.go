//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLCommandBuffer
*/

type MTLCommandBuffer struct {

}

func (r *MTLCommandBuffer) ResourceStateCommandEncoderWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) KernelStartTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) GPUStartTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) GPUEndTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) EncodeWaitForEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) Commit() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) AddCompletedHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) BlitCommandEncoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) RenderCommandEncoderWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) ParallelRenderCommandEncoderWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) Enqueue() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) BlitCommandEncoderWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) ResourceStateCommandEncoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) CommandQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) Logs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) Error() (*foundation.NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) WaitUntilCompleted() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) ComputeCommandEncoderWithDispatchType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) EncodeSignalEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) PopDebugGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) Status() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) ComputeCommandEncoderWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) ComputeCommandEncoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) ErrorOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) PresentDrawable() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) PushDebugGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) WaitUntilScheduled() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) RetainedReferences() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) KernelEndTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) AddScheduledHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) AccelerationStructureCommandEncoderWithDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBuffer) AccelerationStructureCommandEncoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

