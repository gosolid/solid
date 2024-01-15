//js:package native/macos/foundation
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
  *NSSecureCoding
}

func (r *NSUnitFrequency) Terahertz() (*NSUnitFrequency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Hertz() (*NSUnitFrequency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Microhertz() (*NSUnitFrequency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Nanohertz() (*NSUnitFrequency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) FramesPerSecond() (*NSUnitFrequency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Gigahertz() (*NSUnitFrequency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Megahertz() (*NSUnitFrequency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Kilohertz() (*NSUnitFrequency, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitFrequency) Millihertz() (*NSUnitFrequency, error) {
  return nil, fmt.Errorf("unimplemented")
}

