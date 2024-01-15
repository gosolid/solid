//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreData.NSFetchedResultsControllerDelegate
*/

type NSFetchedResultsControllerDelegate struct {

}

func (r *NSFetchedResultsControllerDelegate) Controller() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsControllerDelegate) ControllerWillChangeContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchedResultsControllerDelegate) ControllerDidChangeContent() error {
  return fmt.Errorf("unimplemented")
}

