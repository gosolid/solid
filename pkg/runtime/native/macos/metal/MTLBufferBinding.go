//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLBufferBinding
*/

type MTLBufferBinding struct {

}

func (r *MTLBufferBinding) BufferDataSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLBufferBinding) BufferDataType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLBufferBinding) BufferStructType() (*MTLStructType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBufferBinding) BufferPointerType() (*MTLPointerType, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLBufferBinding) BufferAlignment() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

