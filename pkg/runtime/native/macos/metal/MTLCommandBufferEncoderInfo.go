//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLCommandBufferEncoderInfo
*/

type MTLCommandBufferEncoderInfo struct {

}

func (r *MTLCommandBufferEncoderInfo) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBufferEncoderInfo) DebugSignposts() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCommandBufferEncoderInfo) ErrorState() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

