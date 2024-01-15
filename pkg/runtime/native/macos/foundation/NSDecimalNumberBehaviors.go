//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSDecimalNumberBehaviors
*/

type NSDecimalNumberBehaviors struct {

}

func (r *NSDecimalNumberBehaviors) Scale() (int16, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumberBehaviors) ExceptionDuringOperation() (*NSDecimalNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDecimalNumberBehaviors) RoundingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

