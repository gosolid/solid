//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Metal.MTLCaptureDescriptor : objc.NSObject
*/

type MTLCaptureDescriptor struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *MTLCaptureDescriptor) Destination() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) SetDestination() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) OutputURL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) SetOutputURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) CaptureObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) SetCaptureObject() error {
  return fmt.Errorf("unimplemented")
}

