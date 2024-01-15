//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitPower : Foundation.NSDimension
*/

type NSUnitPower struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitPower) Gigawatts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Watts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Picowatts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Femtowatts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Horsepower() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Terawatts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Megawatts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Kilowatts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Milliwatts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Microwatts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitPower) Nanowatts() (*NSUnitPower, error) {
  return nil, fmt.Errorf("unimplemented")
}

