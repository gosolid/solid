//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSAttributeDescription : CoreData.NSPropertyDescription
*/

type NSAttributeDescription struct {
  *NSPropertyDescription

}

func (r *NSAttributeDescription) SetAttributeValueClassName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) SetPreservesValueInHistoryOnDeletion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) AttributeValueClassName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) DefaultValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) VersionHash() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) SetAllowsExternalBinaryDataStorage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) PreservesValueInHistoryOnDeletion() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) AllowsExternalBinaryDataStorage() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) AllowsCloudEncryption() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) SetAllowsCloudEncryption() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) AttributeType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) SetAttributeType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) SetDefaultValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) ValueTransformerName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAttributeDescription) SetValueTransformerName() error {
  return fmt.Errorf("unimplemented")
}

