//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Metal.MTLMotionKeyframeData : objc.NSObject
*/

type MTLMotionKeyframeData struct {
  *objc.NSObject

}

func (r *MTLMotionKeyframeData) SetBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMotionKeyframeData) Offset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLMotionKeyframeData) SetOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLMotionKeyframeData) Data() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMotionKeyframeData) Buffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

