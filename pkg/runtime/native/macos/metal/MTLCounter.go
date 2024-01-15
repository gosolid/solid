//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLCounter
*/

type MTLCounter struct {

}

func (r *MTLCounter) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

