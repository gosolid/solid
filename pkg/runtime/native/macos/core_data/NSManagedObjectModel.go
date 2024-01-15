//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSManagedObjectModel : objc.NSObject
*/

type NSManagedObjectModel struct {
  *objc.NSObject
  *foundation.NSCoding
  *foundation.NSCopying
  *foundation.NSFastEnumeration
}

func (r *NSManagedObjectModel) VersionIdentifiers() (*foundation.NSSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) VersionChecksum() (*objc.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) ChecksumsForVersionedModelAtURL() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) LocalizationDictionary() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) Configurations() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) FetchRequestTemplatesByName() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) EntityVersionHashesByName() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) FetchRequestTemplateForName() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) Entities() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) SetEntities() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) FetchRequestFromTemplateWithName() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) ModelByMergingModels() (*NSManagedObjectModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) SetFetchRequestTemplate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) IsConfiguration() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) EntitiesByName() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) SetLocalizationDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) SetVersionIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) MergedModelFromBundles() (*NSManagedObjectModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModel) EntitiesForConfiguration() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

