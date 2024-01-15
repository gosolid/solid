//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSRelationshipDescription : CoreData.NSPropertyDescription
*/

type NSRelationshipDescription struct {
  *NSPropertyDescription

}

func (r *NSRelationshipDescription) SetInverseRelationship() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) MinCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) SetMinCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) DeleteRule() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) SetOrdered() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) InverseRelationship() (*NSRelationshipDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) MaxCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) SetDeleteRule() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) SetDestinationEntity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) SetMaxCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) IsToMany() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) VersionHash() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) IsOrdered() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRelationshipDescription) DestinationEntity() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

