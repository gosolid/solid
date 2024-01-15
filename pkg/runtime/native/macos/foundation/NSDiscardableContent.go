//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSDiscardableContent
*/

type NSDiscardableContent struct {

}

func (r *NSDiscardableContent) BeginContentAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSDiscardableContent) EndContentAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiscardableContent) DiscardContentIfPossible() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiscardableContent) IsContentDiscarded() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

