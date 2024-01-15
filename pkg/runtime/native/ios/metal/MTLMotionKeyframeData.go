//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLMotionKeyframeData : objc.NSObject
*/

type MTLMotionKeyframeData struct {
  *objc.NSObject

}

func (r *MTLMotionKeyframeData) Data() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMotionKeyframeData) Buffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMotionKeyframeData) SetBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMotionKeyframeData) Offset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLMotionKeyframeData) SetOffset() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

