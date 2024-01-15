//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSMetadataQuery : objc.NSObject
*/

type NSMetadataQuery struct {
  *objc.NSObject

}

func (r *NSMetadataQuery) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SetPredicate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) IsStarted() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) StartQuery() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SearchItems() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) ValueLists() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SetSortDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SetValueListAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) OperationQueue() (*NSOperationQueue, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) DisableUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) ResultAtIndex() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) EnumerateResultsUsingBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SearchScopes() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SetOperationQueue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) EnableUpdates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SortDescriptors() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) ValueListAttributes() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) NotificationBatchingInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) Results() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) GroupedResults() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) StopQuery() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) GroupingAttributes() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) EnumerateResultsWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SetGroupingAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SetSearchScopes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SetSearchItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) IsGathering() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) IsStopped() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) IndexOfResult() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) ValueOfAttribute() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) Predicate() (*NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) SetNotificationBatchingInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMetadataQuery) ResultCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

