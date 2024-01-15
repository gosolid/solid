//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLCaptureDescriptor : objc.NSObject
*/

type MTLCaptureDescriptor struct {
  *objc.NSObject

}

func (r *MTLCaptureDescriptor) CaptureObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) SetCaptureObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) Destination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) SetDestination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) OutputURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCaptureDescriptor) SetOutputURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

