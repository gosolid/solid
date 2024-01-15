//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UICollectionViewDiffableDataSource : objc.NSObject
*/

type UICollectionViewDiffableDataSource struct {
  *objc.NSObject

}

func (r *UICollectionViewDiffableDataSource) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) ApplySnapshot() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) ItemIdentifierForIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) ReorderingHandlers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) InitWithCollectionView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) IndexForSectionIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) IndexPathForItemIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) SupplementaryViewProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) ApplySnapshotUsingReloadData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) SectionIdentifierForIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) SetSupplementaryViewProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) SectionSnapshotHandlers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) Snapshot() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) SnapshotForSection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) SetReorderingHandlers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UICollectionViewDiffableDataSource) SetSectionSnapshotHandlers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

