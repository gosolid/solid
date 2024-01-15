//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CoreData.NSFetchedResultsSectionInfo
*/

type NSFetchedResultsSectionInfo struct {

}

func (r *NSFetchedResultsSectionInfo) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsSectionInfo) IndexTitle() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsSectionInfo) NumberOfObjects() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsSectionInfo) Objects() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

