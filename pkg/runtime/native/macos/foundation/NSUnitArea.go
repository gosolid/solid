//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitArea : Foundation.NSDimension
*/

type NSUnitArea struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitArea) SquareMeters() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareCentimeters() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareMillimeters() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareYards() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareMegameters() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareKilometers() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) Ares() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) Hectares() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareMicrometers() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareNanometers() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareMiles() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) Acres() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareInches() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitArea) SquareFeet() (*NSUnitArea, error) {
  return nil, fmt.Errorf("unimplemented")
}

