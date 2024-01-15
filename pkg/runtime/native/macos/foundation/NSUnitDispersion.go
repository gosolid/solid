//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnitDispersion : Foundation.NSDimension
*/

type NSUnitDispersion struct {
  *NSDimension
  *NSSecureCoding
}

func (r *NSUnitDispersion) PartsPerMillion() (*NSUnitDispersion, error) {
  return nil, fmt.Errorf("unimplemented")
}

