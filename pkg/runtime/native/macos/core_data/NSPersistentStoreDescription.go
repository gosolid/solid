//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface CoreData.NSPersistentStoreDescription : objc.NSObject
*/

type NSPersistentStoreDescription struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSPersistentStoreDescription) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetTimeout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetShouldMigrateStoreAutomatically() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetOption() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) IsReadOnly() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetReadOnly() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) Timeout() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) ShouldAddStoreAsynchronously() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetShouldAddStoreAsynchronously() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetConfiguration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) URL() (*core_foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SetShouldInferMappingModelAutomatically() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) Configuration() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) Options() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) SqlitePragmas() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) ShouldMigrateStoreAutomatically() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) ShouldInferMappingModelAutomatically() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) PersistentStoreDescriptionWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersistentStoreDescription) Type() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

