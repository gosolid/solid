//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreData.NSCoreDataCoreSpotlightDelegate : objc.NSObject
*/

type NSCoreDataCoreSpotlightDelegate struct {
  *objc.NSObject

}

func (r *NSCoreDataCoreSpotlightDelegate) SearchableIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCoreDataCoreSpotlightDelegate) DomainIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoreDataCoreSpotlightDelegate) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoreDataCoreSpotlightDelegate) StopSpotlightIndexing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCoreDataCoreSpotlightDelegate) AttributeSetForObject() (*CSSearchableItemAttributeSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoreDataCoreSpotlightDelegate) IsIndexingEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCoreDataCoreSpotlightDelegate) IndexName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoreDataCoreSpotlightDelegate) InitForStoreWithDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCoreDataCoreSpotlightDelegate) StartSpotlightIndexing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCoreDataCoreSpotlightDelegate) DeleteSpotlightIndexWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

