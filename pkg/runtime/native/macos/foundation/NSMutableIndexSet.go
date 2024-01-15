//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableIndexSet : Foundation.NSIndexSet
*/

type NSMutableIndexSet struct {
  *NSIndexSet

}

func (r *NSMutableIndexSet) AddIndexesInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) RemoveIndexesInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) ShiftIndexesStartingAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) AddIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) RemoveIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) RemoveAllIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) AddIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) RemoveIndex() error {
  return fmt.Errorf("unimplemented")
}

