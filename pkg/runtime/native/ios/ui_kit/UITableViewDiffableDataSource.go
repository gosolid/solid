//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITableViewDiffableDataSource : objc.NSObject
*/

type UITableViewDiffableDataSource struct {
  *objc.NSObject

}

func (r *UITableViewDiffableDataSource) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) Snapshot() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) SectionIdentifierForIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) IndexForSectionIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) ItemIdentifierForIndexPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) DefaultRowAnimation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) InitWithTableView() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) ApplySnapshot() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) ApplySnapshotUsingReloadData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) IndexPathForItemIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) SetDefaultRowAnimation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITableViewDiffableDataSource) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

