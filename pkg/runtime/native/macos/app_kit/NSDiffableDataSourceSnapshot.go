//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSDiffableDataSourceSnapshot : objc.NSObject
*/

type NSDiffableDataSourceSnapshot struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSDiffableDataSourceSnapshot) ItemIdentifiersInSectionWithIdentifier() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) IndexOfItemIdentifier() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) IndexOfSectionIdentifier() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) InsertItemsWithIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) DeleteAllItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) NumberOfItemsInSection() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) AppendSectionsWithIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) NumberOfSections() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) SectionIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) MoveItemWithIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) InsertSectionsWithIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) DeleteSectionsWithIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) NumberOfItems() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ItemIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ReloadItemsWithIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) AppendItemsWithIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) DeleteItemsWithIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) MoveSectionWithIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) ReloadSectionsWithIdentifiers() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDiffableDataSourceSnapshot) SectionIdentifierForSectionContainingItemIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

