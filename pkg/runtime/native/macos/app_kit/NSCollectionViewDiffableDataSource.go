//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSCollectionViewDiffableDataSource : objc.NSObject
*/

type NSCollectionViewDiffableDataSource struct {
  *objc.NSObject
  *NSCollectionViewDataSource
}

func (r *NSCollectionViewDiffableDataSource) InitWithCollectionView() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewDiffableDataSource) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewDiffableDataSource) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewDiffableDataSource) ApplySnapshot() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewDiffableDataSource) ItemIdentifierForIndexPath() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewDiffableDataSource) IndexPathForItemIdentifier() (*foundation.NSIndexPath, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewDiffableDataSource) Snapshot() (*NSDiffableDataSourceSnapshot, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewDiffableDataSource) SupplementaryViewProvider() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCollectionViewDiffableDataSource) SetSupplementaryViewProvider() error {
  return fmt.Errorf("unimplemented")
}

