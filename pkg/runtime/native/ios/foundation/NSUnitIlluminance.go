//js:package native/ios/foundation
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

}

func (r *NSUnitIlluminance) Lux() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

