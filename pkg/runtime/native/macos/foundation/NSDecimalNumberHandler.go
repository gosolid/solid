//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSDecimalNumberHandler : objc.NSObject
*/

type NSDecimalNumberHandler struct {
  *objc.NSObject
  *NSDecimalNumberBehaviors
  *NSCoding
}

func (r *NSDecimalNumberHandler) DefaultDecimalNumberHandler() (*NSDecimalNumberHandler, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumberHandler) InitWithRoundingMode() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumberHandler) DecimalNumberHandlerWithRoundingMode() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

