//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDimension : Foundation.NSUnit
*/

type NSDimension struct {
  *NSUnit
  *NSSecureCoding
}

func (r *NSDimension) InitWithSymbol() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDimension) BaseUnit() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDimension) Converter() (*NSUnitConverter, error) {
  return nil, fmt.Errorf("unimplemented")
}

