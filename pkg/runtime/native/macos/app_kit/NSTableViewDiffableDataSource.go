//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSTableViewDiffableDataSource : objc.NSObject
*/

type NSTableViewDiffableDataSource struct {
  *objc.NSObject
  *NSTableViewDataSource
}

func (r *NSTableViewDiffableDataSource) Snapshot() (*NSDiffableDataSourceSnapshot, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) SetRowViewProvider() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) InitWithTableView() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) SectionIdentifierForRow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) SectionHeaderViewProvider() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) ItemIdentifierForRow() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) RowForSectionIdentifier() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) SetSectionHeaderViewProvider() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) SetDefaultRowAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) ApplySnapshot() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) RowForItemIdentifier() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) RowViewProvider() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTableViewDiffableDataSource) DefaultRowAnimation() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

