//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSManagedObjectContext : objc.NSObject
*/

type NSManagedObjectContext struct {
  *objc.NSObject
  *foundation.NSCoding
  *foundation.NSLocking
}

func (r *NSManagedObjectContext) SetRetainsRegisteredObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) AutomaticallyMergesChangesFromParent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetQueryGenerationFromToken() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ConcurrencyType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) RefreshAllObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) Reset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) InsertObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) TryLock() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ShouldHandleInaccessibleFault() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) PersistentStoreCoordinator() (*NSPersistentStoreCoordinator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetPersistentStoreCoordinator() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ShouldDeleteInaccessibleFaults() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) StalenessInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) MergePolicy() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) Undo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ExistingObjectWithID() (*NSManagedObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) CountForFetchRequest() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) Name() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetStalenessInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetMergePolicy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) InitWithConcurrencyType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ParentContext() (*NSManagedObjectContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) RegisteredObjects() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetAutomaticallyMergesChangesFromParent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ObserveValueForKeyPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) Rollback() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) RetainsRegisteredObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ObjectWithID() (*NSManagedObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) Unlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) PropagatesDeletesAtEndOfEvent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ExecuteFetchRequest() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) Lock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) MergeChangesFromContextDidSaveNotification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetTransactionAuthor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) Save() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) Redo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetPropagatesDeletesAtEndOfEvent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) PerformBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) RefreshObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) AssignObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetShouldDeleteInaccessibleFaults() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ExecuteRequest() (*NSPersistentStoreResult, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ProcessPendingChanges() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) UserInfo() (*foundation.NSMutableDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) DeleteObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ObjectRegisteredForID() (*NSManagedObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) DetectConflictsForObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) ObtainPermanentIDsForObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetParentContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) HasChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) InsertedObjects() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) DeletedObjects() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) QueryGenerationToken() (*NSQueryGenerationToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) TransactionAuthor() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) MergeChangesFromRemoteContextSave() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) UndoManager() (*NSUndoManager, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) SetUndoManager() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) UpdatedObjects() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectContext) PerformBlockAndWait() error {
  return fmt.Errorf("unimplemented")
}

