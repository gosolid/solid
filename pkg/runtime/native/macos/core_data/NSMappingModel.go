//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSMappingModel : objc.NSObject
*/

type NSMappingModel struct {
  *objc.NSObject

}

func (r *NSMappingModel) MappingModelFromBundles() (*NSMappingModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMappingModel) InferredMappingModelForSourceModel() (*NSMappingModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMappingModel) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMappingModel) EntityMappings() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMappingModel) SetEntityMappings() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMappingModel) EntityMappingsByName() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

