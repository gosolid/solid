//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSEntityMapping : objc.NSObject
*/

type NSEntityMapping struct {
  *objc.NSObject

}

func (r *NSEntityMapping) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) DestinationEntityName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetDestinationEntityName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetUserInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) AttributeMappings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetAttributeMappings() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) MappingType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetMappingType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetSourceEntityName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SourceEntityVersionHash() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) DestinationEntityVersionHash() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) RelationshipMappings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) UserInfo() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) EntityMigrationPolicyClassName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SourceEntityName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetSourceEntityVersionHash() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetDestinationEntityVersionHash() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SourceExpression() (*foundation.NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetSourceExpression() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetEntityMigrationPolicyClassName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSEntityMapping) SetRelationshipMappings() error {
  return fmt.Errorf("unimplemented")
}

