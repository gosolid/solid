//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSPersistentHistoryTransaction : objc.NSObject
*/

type NSPersistentHistoryTransaction struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSPersistentHistoryTransaction) FetchRequest() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) Changes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) TransactionNumber() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) EntityDescription() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) BundleID() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) Token() (*NSPersistentHistoryToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) ContextName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) Author() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) ObjectIDNotification() (*foundation.NSNotification, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) Timestamp() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) StoreID() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) ProcessID() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentHistoryTransaction) EntityDescriptionWithContext() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

