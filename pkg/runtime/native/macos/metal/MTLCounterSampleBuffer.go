//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLCounterSampleBuffer
*/

type MTLCounterSampleBuffer struct {

}

func (r *MTLCounterSampleBuffer) ResolveCounterRange() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBuffer) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBuffer) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSampleBuffer) SampleCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

