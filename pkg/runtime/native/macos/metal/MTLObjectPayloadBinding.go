//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLObjectPayloadBinding
*/

type MTLObjectPayloadBinding struct {

}

func (r *MTLObjectPayloadBinding) ObjectPayloadAlignment() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLObjectPayloadBinding) ObjectPayloadDataSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

