//js:package native/ios/foundation
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

}

func (r *NSUnitDispersion) PartsPerMillion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

