//js:package native/macos/foundation
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
  *NSSecureCoding
}

func (r *NSUnitDuration) Microseconds() (*NSUnitDuration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Nanoseconds() (*NSUnitDuration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Picoseconds() (*NSUnitDuration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Hours() (*NSUnitDuration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Minutes() (*NSUnitDuration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Seconds() (*NSUnitDuration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitDuration) Milliseconds() (*NSUnitDuration, error) {
  return nil, fmt.Errorf("unimplemented")
}

