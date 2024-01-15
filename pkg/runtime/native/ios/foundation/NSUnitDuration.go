//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitDuration : Foundation.NSDimension
*/

type NSUnitDuration struct {
  *NSDimension

}

func (r *NSUnitDuration) Hours() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Minutes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Seconds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Milliseconds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Microseconds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Nanoseconds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Picoseconds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

