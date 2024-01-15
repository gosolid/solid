//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSEntityDescription : objc.NSObject
*/

type NSEntityDescription struct {
  *objc.NSObject
  *foundation.NSCoding
  *foundation.NSCopying
  *foundation.NSFastEnumeration
}

func (r *NSEntityDescription) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetSubentities() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) RenamingIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) IsKindOfEntity() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetVersionHashModifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) Indexes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetCoreSpotlightDisplayNameExpression() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) EntityForName() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) PropertiesByName() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) VersionHash() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) UniquenessConstraints() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) InsertNewObjectForEntityForName() (*NSManagedObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) ManagedObjectModel() (*NSManagedObjectModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetManagedObjectClassName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) VersionHashModifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetRenamingIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetCompoundIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) RelationshipsWithDestinationEntity() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) Superentity() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) RelationshipsByName() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) ManagedObjectClassName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SubentitiesByName() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) Properties() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetAbstract() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) AttributesByName() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) CompoundIndexes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) IsAbstract() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) Subentities() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetProperties() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) UserInfo() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) SetUniquenessConstraints() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityDescription) CoreSpotlightDisplayNameExpression() (*foundation.NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

