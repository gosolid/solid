//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol Metal.MTLCounterSet
*/

type MTLCounterSet struct {

}

func (r *MTLCounterSet) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLCounterSet) Counters() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

