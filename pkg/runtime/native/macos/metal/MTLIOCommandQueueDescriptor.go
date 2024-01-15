//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface Metal.MTLIOCommandQueueDescriptor : objc.NSObject
*/

type MTLIOCommandQueueDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLIOCommandQueueDescriptor) SetMaxCommandBufferCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) Priority() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) ScratchBufferAllocator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) SetMaxCommandsInFlight() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) SetScratchBufferAllocator() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) MaxCommandBufferCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) SetPriority() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) SetType() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) MaxCommandsInFlight() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

