//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewDiffableDataSourceSectionSnapshotHandlers : objc.NSObject
*/

type UICollectionViewDiffableDataSourceSectionSnapshotHandlers struct {
  *objc.NSObject

}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) WillCollapseItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) SnapshotForExpandingParentItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) SetSnapshotForExpandingParentItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) ShouldCollapseItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) SetShouldExpandItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) WillExpandItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) SetWillExpandItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) SetShouldCollapseItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) SetWillCollapseItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSourceSectionSnapshotHandlers) ShouldExpandItemHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

