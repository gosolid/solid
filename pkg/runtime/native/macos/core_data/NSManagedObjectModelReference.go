//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSManagedObjectModelReference : objc.NSObject
*/

type NSManagedObjectModelReference struct {
  *objc.NSObject

}

func (r *NSManagedObjectModelReference) VersionChecksum() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModelReference) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModelReference) InitWithModel() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModelReference) InitWithFileURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModelReference) InitWithEntityVersionHashes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModelReference) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSManagedObjectModelReference) ResolvedModel() (*NSManagedObjectModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

