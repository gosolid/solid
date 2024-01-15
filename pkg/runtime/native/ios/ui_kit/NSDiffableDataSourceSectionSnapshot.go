//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.NSDiffableDataSourceSectionSnapshot : objc.NSObject
*/

type NSDiffableDataSourceSectionSnapshot struct {
  *objc.NSObject

}

func (r *NSDiffableDataSourceSectionSnapshot) ReplaceChildrenOfParentItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) ParentOfChildItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) SnapshotOfParentItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) VisualDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) VisibleItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) ContainsItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) LevelOfItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) DeleteItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) ExpandItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) InsertSnapshot() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) IsExpanded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) IsVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) InsertItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) DeleteAllItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) Items() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) ExpandedItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) AppendItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) CollapseItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) IndexOfItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSectionSnapshot) RootItems() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

