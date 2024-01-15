//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CoreData.NSPropertyDescription : objc.NSObject
*/

type NSPropertyDescription struct {
  *objc.NSObject
  *foundation.NSCoding
  *foundation.NSCopying
}

func (r *NSPropertyDescription) SetValidationPredicates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) Entity() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) ValidationWarnings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) SetRenamingIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) SetIndexedBySpotlight() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) RenamingIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) IsOptional() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) SetOptional() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) VersionHash() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) SetVersionHashModifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) IsTransient() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) ValidationPredicates() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) IsIndexed() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) SetIndexed() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) IsIndexedBySpotlight() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) SetTransient() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) UserInfo() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) VersionHashModifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) IsStoredInExternalRecord() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPropertyDescription) SetStoredInExternalRecord() error {
  return fmt.Errorf("unimplemented")
}

