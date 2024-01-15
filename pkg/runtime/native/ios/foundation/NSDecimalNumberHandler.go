//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSDecimalNumberHandler : objc.NSObject
*/

type NSDecimalNumberHandler struct {
  *objc.NSObject

}

func (r *NSDecimalNumberHandler) InitWithRoundingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumberHandler) DecimalNumberHandlerWithRoundingMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumberHandler) DefaultDecimalNumberHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

