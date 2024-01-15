//js:package native/ios/foundation
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

}

func (r *NSDimension) BaseUnit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDimension) Converter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDimension) InitWithSymbol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

