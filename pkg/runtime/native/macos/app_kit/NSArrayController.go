//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSArrayController : AppKit.NSObjectController
*/

type NSArrayController struct {
  *NSObjectController

}

func (r *NSArrayController) SelectionIndexes() (*foundation.NSIndexSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetSelectionIndexes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetAutomaticallyRearrangesObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SortDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) CanInsert() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) CanSelectNext() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetSelectionIndex() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) RemoveSelectedObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) PreservesSelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) AddObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) InsertObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) RemoveObjectsAtArrangedObjectIndexes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetSortDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetClearsFilterPredicateOnInsertion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) AvoidsEmptySelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SelectsInsertedObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SelectedObjects() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) ArrangeObjects() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) Remove() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SelectNext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) AddObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) AutomaticallyRearrangesObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) ArrangedObjects() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetSelectsInsertedObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) AlwaysUsesMultipleValuesMarker() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) AddSelectionIndexes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetSelectedObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) InsertObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) RemoveObjectAtArrangedObjectIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetFilterPredicate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetAlwaysUsesMultipleValuesMarker() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SelectionIndex() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetAvoidsEmptySelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) AddSelectedObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) Add() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) RemoveObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) AutomaticRearrangementKeyPaths() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) ClearsFilterPredicateOnInsertion() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) RemoveObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) FilterPredicate() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SetPreservesSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) RearrangeObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) DidChangeArrangementCriteria() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) RemoveSelectionIndexes() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSArrayController) Insert() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) SelectPrevious() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSArrayController) CanSelectPrevious() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

