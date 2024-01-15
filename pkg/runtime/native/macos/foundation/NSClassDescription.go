//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSClassDescription : objc.NSObject
*/

type NSClassDescription struct {
  *objc.NSObject

}

func (r *NSClassDescription) ClassDescriptionForClass() (*NSClassDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSClassDescription) InverseForRelationshipKey() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSClassDescription) AttributeKeys() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSClassDescription) ToOneRelationshipKeys() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSClassDescription) ToManyRelationshipKeys() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSClassDescription) RegisterClassDescription() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSClassDescription) InvalidateClassDescriptionCache() error {
  return fmt.Errorf("unimplemented")
}

