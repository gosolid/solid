//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSUnitConverter : objc.NSObject
*/

type NSUnitConverter struct {
  *objc.NSObject

}

func (r *NSUnitConverter) BaseUnitValueFromValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnitConverter) ValueFromBaseUnitValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

