//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLCounterSampleBufferDescriptor : objc.NSObject
*/

type MTLCounterSampleBufferDescriptor struct {
  *objc.NSObject

}

func (r *MTLCounterSampleBufferDescriptor) SetSampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) CounterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) SetCounterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) SetLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) StorageMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) SetStorageMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBufferDescriptor) SampleCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

