//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSPersistentCloudKitContainer : CoreData.NSPersistentContainer
*/

type NSPersistentCloudKitContainer struct {
  *NSPersistentContainer

}

func (r *NSPersistentCloudKitContainer) CanDeleteRecordForManagedObjectWithID() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainer) CanModifyManagedObjectsInStore() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainer) InitializeCloudKitSchemaWithOptions() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainer) RecordForManagedObjectID() (*CKRecord, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainer) RecordsForManagedObjectIDs() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainer) RecordIDForManagedObjectID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainer) RecordIDsForManagedObjectIDs() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentCloudKitContainer) CanUpdateRecordForManagedObjectWithID() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

