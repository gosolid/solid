//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLIOCommandQueueDescriptor : objc.NSObject
*/

type MTLIOCommandQueueDescriptor struct {
  *objc.NSObject

}

func (r *MTLIOCommandQueueDescriptor) MaxCommandBufferCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) SetPriority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) SetType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) MaxCommandsInFlight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) SetMaxCommandsInFlight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) SetMaxCommandBufferCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) Priority() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) ScratchBufferAllocator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLIOCommandQueueDescriptor) SetScratchBufferAllocator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

