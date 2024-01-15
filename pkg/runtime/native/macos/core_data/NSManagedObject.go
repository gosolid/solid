//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSManagedObject : objc.NSObject
*/

type NSManagedObject struct {
  *objc.NSObject

}

func (r *NSManagedObject) DidAccessValueForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) PrepareForDeletion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) CommittedValuesForKeys() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) IsInserted() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) InitWithEntity() (*NSManagedObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ValueForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ChangedValues() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ValidateForInsert() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) InitWithContext() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) SetValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ChangedValuesForCurrentEvent() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ValidateForUpdate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) DidChangeValueForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) SetPrimitiveValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) IsUpdated() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ObjectID() (*NSManagedObjectID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) IsDeleted() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) HasChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) FaultingState() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) Entity() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) AwakeFromInsert() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) WillSave() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) PrimitiveValueForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ObservationInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) IsFault() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) AwakeFromSnapshotEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) DidTurnIntoFault() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ValidateForDelete() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) SetObservationInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ManagedObjectContext() (*NSManagedObjectContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) HasPersistentChangedValues() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) HasFaultForRelationshipNamed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) DidSave() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) WillTurnIntoFault() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ContextShouldIgnoreUnmodeledPropertyChanges() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) AwakeFromFetch() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ValidateValue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) FetchRequest() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) ObjectIDsForRelationshipNamed() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) WillAccessValueForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObject) WillChangeValueForKey() error {
  return fmt.Errorf("unimplemented")
}

