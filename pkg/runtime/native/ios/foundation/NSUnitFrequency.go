//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitFrequency : Foundation.NSDimension
*/

type NSUnitFrequency struct {
  *NSDimension

}

func (r *NSUnitFrequency) Terahertz() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Gigahertz() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) FramesPerSecond() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Nanohertz() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Megahertz() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Kilohertz() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Hertz() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Millihertz() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Microhertz() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

