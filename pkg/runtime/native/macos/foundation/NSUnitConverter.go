//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUnitConverter : objc.NSObject
*/

type NSUnitConverter struct {
  *objc.NSObject

}

func (r *NSUnitConverter) BaseUnitValueFromValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUnitConverter) ValueFromBaseUnitValue() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

