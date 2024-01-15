//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLCounterSampleBufferDescriptor : objc.NSObject
*/

type MTLCounterSampleBufferDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLCounterSampleBufferDescriptor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) StorageMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) SetStorageMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) SampleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) SetSampleCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) CounterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) SetCounterSet() error {
  return fmt.Errorf("unimplemented")
}

