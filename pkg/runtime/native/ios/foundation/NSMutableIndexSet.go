//js:package native/ios/foundation
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

func (r *NSMutableIndexSet) AddIndexesInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) RemoveIndexesInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) ShiftIndexesStartingAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) AddIndexes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) RemoveIndexes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) RemoveAllIndexes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) AddIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableIndexSet) RemoveIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

