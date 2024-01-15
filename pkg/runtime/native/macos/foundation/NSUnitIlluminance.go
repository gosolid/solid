//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitIlluminance : Foundation.NSDimension
*/

type NSUnitIlluminance struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitIlluminance) Lux() (*NSUnitIlluminance, error) {
  return nil, fmt.Errorf("unimplemented")
}

