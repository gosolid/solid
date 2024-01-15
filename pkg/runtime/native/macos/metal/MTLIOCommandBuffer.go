//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLIOCommandBuffer
*/

type MTLIOCommandBuffer struct {

}

func (r *MTLIOCommandBuffer) PushDebugGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) PopDebugGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) Enqueue() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) WaitForEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) SignalEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) TryCancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) CopyStatusToBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) AddBarrier() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) Status() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) Error() (*foundation.NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) LoadTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) LoadBytes() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) LoadBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) WaitUntilCompleted() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) AddCompletedHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandBuffer) Commit() error {
  return fmt.Errorf("unimplemented")
}

