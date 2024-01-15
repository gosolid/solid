//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTreeController : AppKit.NSObjectController
*/

type NSTreeController struct {
  *NSObjectController

}

func (r *NSTreeController) SelectedNodes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) Remove() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) RemoveObjectAtArrangedObjectIndexPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SelectionIndexPaths() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) MoveNode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetPreservesSelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) Insert() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) InsertObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) RemoveSelectionIndexPaths() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) CanInsertChild() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) PreservesSelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SelectsInsertedObjects() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetAlwaysUsesMultipleValuesMarker() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SelectedObjects() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) Add() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetLeafKeyPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SortDescriptors() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) Content() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetAvoidsEmptySelection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) RemoveObjectsAtArrangedObjectIndexPaths() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) MoveNodes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetSortDescriptors() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) AvoidsEmptySelection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SelectionIndexPath() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetSelectionIndexPath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) CountKeyPathForNode() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) ArrangedObjects() (*NSTreeNode, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) AlwaysUsesMultipleValuesMarker() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) InsertChild() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) ChildrenKeyPathForNode() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) LeafKeyPathForNode() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) CanInsert() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) CanAddChild() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetSelectsInsertedObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) AddChild() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetChildrenKeyPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) LeafKeyPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) AddSelectionIndexPaths() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) ChildrenKeyPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) CountKeyPath() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetCountKeyPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) RearrangeObjects() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) InsertObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTreeController) SetSelectionIndexPaths() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

