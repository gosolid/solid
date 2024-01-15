//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSFetchedResultsController : objc.NSObject
*/

type NSFetchedResultsController struct {
  *objc.NSObject

}

func (r *NSFetchedResultsController) ManagedObjectContext() (*NSManagedObjectContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) SectionIndexTitles() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) PerformFetch() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) IndexPathForObject() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) SectionForSectionIndexTitle() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) FetchRequest() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) CacheName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) FetchedObjects() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) InitWithFetchRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) DeleteCacheWithName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) SectionIndexTitleForSectionName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) SectionNameKeyPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) ObjectAtIndexPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsController) Sections() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

