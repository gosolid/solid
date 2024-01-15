//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSIndexSet : objc.NSObject
*/

type NSIndexSet struct {
  *objc.NSObject
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
}

func (r *NSIndexSet) IndexLessThanIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexLessThanOrEqualToIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) GetIndexes() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) EnumerateIndexesUsingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexesPassingTest() (*NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexesInRange() (*NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) EnumerateRangesInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexesWithOptions() (*NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) EnumerateRangesUsingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) LastIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexSet() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) InitWithIndexSet() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) EnumerateIndexesInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexWithOptions() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexSetWithIndexesInRange() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) EnumerateRangesWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexInRange() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) ContainsIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) ContainsIndexes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) InitWithIndexesInRange() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexGreaterThanIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexGreaterThanOrEqualToIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) CountOfIndexesInRange() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) InitWithIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IsEqualToIndexSet() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) EnumerateIndexesWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) FirstIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexSetWithIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) ContainsIndexesInRange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IntersectsIndexesInRange() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSIndexSet) IndexPassingTest() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

