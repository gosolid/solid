//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSFetchRequest : CoreData.NSPersistentStoreRequest
*/

type NSFetchRequest struct {
  *NSPersistentStoreRequest
  *foundation.NSCoding
  *foundation.NSCopying
}

func (r *NSFetchRequest) SetFetchBatchSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetPredicate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetFetchLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) ResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetPropertiesToGroupBy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) HavingPredicate() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) FetchRequestWithEntityName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) ReturnsObjectsAsFaults() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) PropertiesToGroupBy() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetPropertiesToFetch() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetAffectedStores() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetIncludesSubentities() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetReturnsObjectsAsFaults() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) Execute() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetResultType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetIncludesPendingChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) Predicate() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetSortDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) IncludesSubentities() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) IncludesPendingChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) FetchOffset() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) InitWithEntityName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) Entity() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) EntityName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetShouldRefreshRefetchedObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetFetchOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) FetchBatchSize() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) ShouldRefreshRefetchedObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) PropertiesToFetch() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) AffectedStores() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) RelationshipKeyPathsForPrefetching() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetRelationshipKeyPathsForPrefetching() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) IncludesPropertyValues() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) ReturnsDistinctResults() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetReturnsDistinctResults() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetHavingPredicate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetEntity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SortDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) FetchLimit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequest) SetIncludesPropertyValues() error {
  return fmt.Errorf("unimplemented")
}

